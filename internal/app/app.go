package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Amir-Sadati/order-packing/internal/config"
	"github.com/Amir-Sadati/order-packing/internal/constants"
	"github.com/Amir-Sadati/order-packing/internal/database/redisdb"
	"github.com/Amir-Sadati/order-packing/internal/handler/api"
	"github.com/Amir-Sadati/order-packing/internal/router"
	"github.com/Amir-Sadati/order-packing/internal/service/pack"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type App struct {
	config     *config.Config
	r          *gin.Engine
	httpServer *http.Server
	rdb        *redis.Client
}

func New() *App {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("error in loading config")
	}
	return &App{
		config: cfg,
	}
}

func (a *App) Run() {
	rdb, err := redisdb.NewClient(a.config.Redis)
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}
	a.rdb = rdb

	err = a.seedDefaultPackSizes()
	if err != nil {
		log.Fatalf("failed to seed to redis: %v", err)
	}

	packService := pack.NewService(rdb)
	packHandler := api.NewPackHandler(packService)

	r := router.New(packHandler)
	a.r = r
	go a.ServeHTTP()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()
	log.Println("Shutting down gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := a.httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP shutdown error: %v", err)
	}

}

func (a *App) ServeHTTP() {
	a.httpServer = &http.Server{
		Addr:         net.JoinHostPort(a.config.HTTP.Host, a.config.HTTP.Port),
		Handler:      a.r,
		WriteTimeout: 10 * time.Second,
	}

	err := a.httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}

// seedDefaultPackSizes inserts default pack sizes into Redis if none exist.
func (a *App) seedDefaultPackSizes() error {
	defaultSizes := []int{250, 500, 1000, 2000, 5000}
	count, err := a.rdb.ZCard(context.Background(), string(constants.RedisKeyPackSizes)).Result()
	if err != nil {
		log.Printf("failed to check Redis: %v", err)
		return err
	}
	if count == 0 {
		var zData []redis.Z
		for _, size := range defaultSizes {
			zData = append(zData, redis.Z{
				Score:  float64(size),
				Member: size,
			})
		}
		if err := a.rdb.ZAdd(context.Background(), string(constants.RedisKeyPackSizes), zData...).Err(); err != nil {
			log.Printf("failed to seed pack sizes to Redis: %v", err)
			return err
		}
		log.Println("Default pack sizes seeded into Redis.")
		return nil
	}
	return nil
}
