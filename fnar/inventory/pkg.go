package inventory

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/koo04/fnar-go/fnar"
	"github.com/pkg/errors"
)

type StorageItem struct {
	Id            string  `json:"MaterialId"`
	Name          string  `json:"MaterialName"`
	Ticker        string  `json:"MaterialTicker"`
	Category      string  `json:"MaterialCategory"`
	Weight        float64 `json:"MaterialWeight"`
	Volume        float64 `json:"MaterialVolume"`
	Amount        int32   `json:"MaterialAmount"`
	Value         float64 `json:"MaterialValue"`
	ValueCurrency string  `json:"MaterialValueCurrency"`
	Type          string  `json:"Type"`
	TotalWeight   float64 `json:"TotalWeight"`
	TotalVolume   float64 `json:"TotalVolume"`
}

type Inventory struct {
	Items             []*StorageItem `json:"StorageItems"`
	Id                string         `json:"StorageId"`
	AddressableId     string         `json:"AddressableId"`
	Name              string         `json:"Name"`
	WeightLoad        float64        `json:"WeightLoad"`
	WeightCapacity    float64        `json:"WeightCapacity"`
	VolumeLoad        float64        `json:"VolumeLoad"`
	VolumeCapacity    float64        `json:"VolumeCapacity"`
	FixedStore        bool           `json:"FixedStore"`
	Type              Type           `json:"Type"`
	UserNameSubmitted string         `json:"UserNameSubmitted"`
	Timestamp         *time.Time     `json:"Timestamp"`
}

const endpoint = "/storage"

func GetAll(ctx context.Context, username string, auth *fnar.Authentication) ([]*Inventory, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fnar.BaseUrl+endpoint+"/"+username, nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating request for getting all inventories")
	}
	req.Header.Add("Authorization", auth.AuthToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "getting all ships")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(b))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading the response body")
	}

	// setup inventory
	inventoriesJson := []map[string]interface{}{}
	if err := json.Unmarshal(body, &inventoriesJson); err != nil {
		return nil, errors.Wrap(err, "decoding json response")
	}

	inventories := []*Inventory{}
	for _, inventoryJson := range inventoriesJson {
		inventory := &Inventory{}
		if err := inventory.parse(inventoryJson); err != nil {
			return nil, errors.Wrap(err, "parsing inventory")
		}

		inventories = append(inventories, inventory)
	}

	return inventories, nil
}

func (i *Inventory) parse(inventory map[string]interface{}) error {
	// correct timestamp to time.Time
	var timestamp time.Time
	var err error
	if ts, ok := inventory["Timestamp"]; ok {
		timestamp, err = time.Parse("2006-01-02T15:04:05.999999", ts.(string))
		if err != nil {
			return err
		}

		inventory["Timestamp"] = timestamp
	}

	// convert type type to actual Type
	if t, ok := inventory["Type"].(string); ok {
		inventory["Type"] = ToType(t)
	}

	// convert to Inventory
	inventoryM, err := json.Marshal(inventory)
	if err != nil {
		return errors.Wrap(err, "marshaling inventory map")
	}
	if err := json.Unmarshal(inventoryM, &i); err != nil {
		return errors.Wrap(err, "unmarshaling to Inventory")
	}

	return nil
}
