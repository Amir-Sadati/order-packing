package pack

type CalculatePackResponse struct {
	Packs map[int]int `json:"packs"`
}

type GetPackSizesResponse struct {
	Sizes []int `json:"sizes"`
}
