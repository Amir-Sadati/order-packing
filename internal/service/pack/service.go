package pack

import (
	"context"
	"errors"
	"strconv"

	"github.com/Amir-Sadati/order-packing/internal/constants"
	"github.com/redis/go-redis/v9"
)

var (
	ErrInvalidOrderItemQuantity = errors.New("invalid order-item-quantity")
	ErrNotFoundPackSize         = errors.New("pack size not found")
)

type Service struct {
	rdb *redis.Client
}

func NewService(redisClinet *redis.Client) *Service {
	return &Service{
		rdb: redisClinet,
	}
}

func (s *Service) CalculatePack(context context.Context, req CalculatePackRequest) (CalculatePackResponse, error) {
	if req.OrderItemQuantity < 1 {
		return CalculatePackResponse{}, ErrInvalidOrderItemQuantity
	}

	packVals, err := s.rdb.ZRevRange(
		context,
		string(constants.RedisKeyPackSizes),
		0, -1,
	).Result()
	if err != nil {
		return CalculatePackResponse{}, err
	}

	packSizez := make([]int, len(packVals))
	for i, v := range packVals {
		packSize, err := strconv.Atoi(v)
		if err != nil {
			return CalculatePackResponse{}, err
		}
		packSizez[i] = packSize
	}

	resp := calculatePacks(req.OrderItemQuantity, packSizez)
	return CalculatePackResponse{
		Packs: resp.Packs,
	}, nil
}

// GetPackSizes returns all pack sizes in descending order (largest to smallest)
func (s *Service) GetPackSizes(ctx context.Context) (GetPackSizesResponse, error) {
	vals, err := s.rdb.ZRevRange(ctx, string(constants.RedisKeyPackSizes), 0, -1).Result()
	if err != nil {
		return GetPackSizesResponse{}, err
	}

	packSizes := make([]int, 0, len(vals))
	for _, v := range vals {
		n, err := strconv.Atoi(v)
		if err != nil {
			return GetPackSizesResponse{}, err
		}
		packSizes = append(packSizes, n)
	}

	return GetPackSizesResponse{Sizes: packSizes}, nil
}

// AddPackSize adds a new pack size to the Redis sorted set
func (s *Service) AddPackSize(ctx context.Context, req AddPackSizeRequest) error {
	return s.rdb.ZAdd(ctx, string(constants.RedisKeyPackSizes), redis.Z{
		Score:  float64(req.Size),
		Member: req.Size,
	}).Err()
}

// RemovePackSize removes a pack size from the Redis sorted set
func (s *Service) RemovePackSize(ctx context.Context, req RemovePackSizeRequest) error {
	removedCount, err := s.rdb.ZRem(ctx, string(constants.RedisKeyPackSizes), req.Size).Result()
	if err != nil {
		return err
	}
	if removedCount == 0 {
		return ErrNotFoundPackSize
	}
	return nil
}
