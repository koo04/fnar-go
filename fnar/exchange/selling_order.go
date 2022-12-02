package exchange

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type SellingOrder struct {
	OrderID     string  `json:"OrderId"`
	CompanyID   string  `json:"CompanyId"`
	CompanyName string  `json:"CompanyName"`
	CompanyCode string  `json:"CompanyCode"`
	ItemCount   int     `json:"ItemCount"`
	ItemCost    float32 `json:"ItemCost"`
}

func (so *SellingOrder) Parse(m map[string]interface{}) error {
	mb, err := json.Marshal(m)
	if err != nil {
		return errors.Wrap(err, "marshaling selling order map[string]interface")
	}
	if err := json.Unmarshal(mb, so); err != nil {
		return errors.Wrap(err, "unmarshaling to SellingOrder")
	}

	return nil
}
