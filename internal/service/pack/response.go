package pack

// CalculatePackResponse represents the response for pack calculation
type CalculatePackResponse struct {
	Packs map[int]int `json:"packs"`
}

// GetPackSizesResponse represents the response for getting pack sizes
type GetPackSizesResponse struct {
	Sizes []int `json:"sizes"`
}
