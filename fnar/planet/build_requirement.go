package planet

type BuildRequirement struct {
	MaterialName     string  `json:"MaterialName"`
	MaterialID       string  `json:"MaterialId"`
	MaterialTicker   string  `json:"MaterialTicker"`
	MaterialCategory string  `json:"MaterialCategory"`
	MaterialAmount   int     `json:"MaterialAmount"`
	MaterialWeight   float64 `json:"MaterialWeight"`
	MaterialVolume   float64 `json:"MaterialVolume"`
}
