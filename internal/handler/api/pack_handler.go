package api

import (
	"errors"
	"net/http"

	"github.com/Amir-Sadati/order-packing/internal/handler/api/response"
	"github.com/Amir-Sadati/order-packing/internal/service/pack"
	"github.com/gin-gonic/gin"
)

type PackHandler struct {
	packService *pack.Service
}

func NewPackHandler(packService *pack.Service) *PackHandler {
	return &PackHandler{
		packService: packService,
	}
}

// CalculatePack godoc
//
//	@Summary		Calculate a new pack by order-items
//	@Description	Calculates an optimal pack combination using orderItemQuantity as query param
//	@Tags			packs
//	@Accept			json
//	@Produce		json
//	@Param			orderItemQuantity	query		uint64	true	"Number of items to order"
//	@Success		200	{object}	pack.CalculatePackResponse
//	@Failure		400	{object}	response.ApiResponseNoData
//	@Failure		404	{object}	response.ApiResponseNoData
//	@Failure		500	{object}	response.ApiResponseNoData
//	@Router			/api/v1/packs/calculate [get]
func (h *PackHandler) CalculatePack(c *gin.Context) {
	var req pack.CalculatePackRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		response.WriteFailNoData(c.Writer, http.StatusBadRequest, "Invalid query parameters", "")
		return
	}

	result, err := h.packService.CalculatePack(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, pack.ErrInvalidOrderItemQuantity) {
			response.WriteFailNoData(c.Writer, http.StatusBadRequest, err.Error(), "")
		} else {
			response.WriteFailNoData(c.Writer, http.StatusInternalServerError, "internal_error", "Something went wrong")
		}
		return
	}

	response.WriteSuccess(c.Writer, result, "pack calculated successfully")
}

// GetPackSizes godoc
//
//	@Summary		Get all pack sizes
//	@Description	Returns all available pack sizes from Redis
//	@Tags			packs
//	@Produce		json
//	@Success		200	{object}	pack.GetPackSizesResponse
//	@Failure		500	{object}	response.ApiResponseNoData
//	@Router			/api/v1/packs/sizes [get]
func (h *PackHandler) GetPackSizes(c *gin.Context) {
	result, err := h.packService.GetPackSizes(c.Request.Context())
	if err != nil {
		response.WriteFailNoData(c.Writer, http.StatusInternalServerError, "internal_error", "Something went wrong")
		return
	}
	response.WriteSuccess(c.Writer, result, "pack sizes fetched successfully")
}

// AddPackSize godoc
//
//	@Summary		Add a new pack size
//	@Description	Adds a new pack size to the Redis sorted set
//	@Tags			packs
//	@Accept			json
//	@Produce		json
//	@Param			body	body		pack.AddPackSizeRequest	true	"Pack size to add"
//	@Success		200	{object}	response.ApiResponseNoData
//	@Failure		400	{object}	response.ApiResponseNoData
//	@Failure		500	{object}	response.ApiResponseNoData
//	@Router			/api/v1/packs/sizes [post]
func (h *PackHandler) AddPackSize(c *gin.Context) {
	var req pack.AddPackSizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteFailNoData(c.Writer, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	if err := h.packService.AddPackSize(c.Request.Context(), req); err != nil {
		response.WriteFailNoData(c.Writer, http.StatusInternalServerError, "internal_error", err.Error())
		return
	}

	response.WriteSuccessNoData(c.Writer, "pack size added successfully")
}

// RemovePackSize godoc
//
//	@Summary		Remove a pack size
//	@Description	Removes a pack size from the Redis sorted set
//	@Tags			packs
//	@Accept			json
//	@Produce		json
//	@Param			body	body		pack.RemovePackSizeRequest	true	"Pack size to remove"
//	@Success		200	{object}	response.ApiResponseNoData
//	@Failure		400	{object}	response.ApiResponseNoData
//	@Failure		500	{object}	response.ApiResponseNoData
//	@Router			/api/v1/packs/sizes [delete]
func (h *PackHandler) RemovePackSize(c *gin.Context) {
	var req pack.RemovePackSizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteFailNoData(c.Writer, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	if err := h.packService.RemovePackSize(c.Request.Context(), req); err != nil {
		if errors.Is(err, pack.ErrNotFoundPackSize) {
			response.WriteFailNoData(c.Writer, http.StatusBadRequest, err.Error(), "")
			return
		}
		response.WriteFailNoData(c.Writer, http.StatusInternalServerError, err.Error(), "")
		return
	}

	response.WriteSuccessNoData(c.Writer, "pack size removed successfully")
}
