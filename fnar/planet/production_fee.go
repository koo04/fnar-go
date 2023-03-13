package planet

type ProductionFee struct {
	Category       string  `json:"Category"`
	WorkforceLevel string  `json:"WorkforceLevel"`
	FeeAmount      float64 `json:"FeeAmount"`
	FeeCurrency    string  `json:"FeeCurrency"`
}
