package planet

type Resource struct {
	MaterialID   string  `json:"MaterialId"`
	ResourceType string  `json:"ResourceType"`
	Factor       float64 `json:"Factor"`
}
