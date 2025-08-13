package pack

// CalculatePackRequest represents a request to calculate optimal packing
type CalculatePackRequest struct {
	OrderItemQuantity int `form:"orderItemQuantity"`
}

// AddPackSizeRequest represents a request to add a new pack size
type AddPackSizeRequest struct {
	Size int `json:"size" binding:"required"`
}

// RemovePackSizeRequest represents a request to remove a pack size
type RemovePackSizeRequest struct {
	Size int `json:"size" binding:"required"`
}
