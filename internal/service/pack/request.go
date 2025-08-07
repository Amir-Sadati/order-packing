package pack

type CalculatePackRequest struct {
	OrderItemQuantity int `form:"orderItemQuantity"`
}

type AddPackSizeRequest struct {
	Size int `json:"size" binding:"required"`
}

type RemovePackSizeRequest struct {
	Size int `json:"size" binding:"required"`
}
