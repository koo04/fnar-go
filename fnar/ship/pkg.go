package ship

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/koo04/fnar-go/fnar"
	"github.com/pkg/errors"
)

type Ship struct {
	RepairMaterials []struct {
		ShipRepairMaterialID string `json:"ShipRepairMaterialId"`
		MaterialName         string `json:"MaterialName"`
		MaterialID           string `json:"MaterialId"`
		MaterialTicker       string `json:"MaterialTicker"`
		Amount               int    `json:"Amount"`
	} `json:"RepairMaterials"`
	AddressLines []struct {
		LineID    string `json:"LineId"`
		LineType  string `json:"LineType"`
		NaturalID string `json:"NaturalId"`
		Name      string `json:"Name"`
	} `json:"AddressLines"`
	ShipID                   string    `json:"ShipId"`
	StoreID                  string    `json:"StoreId"`
	StlFuelStoreID           string    `json:"StlFuelStoreId"`
	FtlFuelStoreID           string    `json:"FtlFuelStoreId"`
	Registration             string    `json:"Registration"`
	Name                     *string   `json:"Name"`
	CommissioningTimeEpochMs int64     `json:"CommissioningTimeEpochMs"`
	BlueprintNaturalID       string    `json:"BlueprintNaturalId"`
	FlightID                 string    `json:"FlightId"`
	Acceleration             float64   `json:"Acceleration"`
	Thrust                   int       `json:"Thrust"`
	Mass                     float64   `json:"Mass"`
	OperatingEmptyMass       float64   `json:"OperatingEmptyMass"`
	ReactorPower             int       `json:"ReactorPower"`
	EmitterPower             int       `json:"EmitterPower"`
	Volume                   int       `json:"Volume"`
	Condition                float64   `json:"Condition"`
	LastRepairEpochMs        int       `json:"LastRepairEpochMs"`
	Location                 string    `json:"Location"`
	StlFuelFlowRate          float64   `json:"StlFuelFlowRate"`
	UserNameSubmitted        string    `json:"UserNameSubmitted"`
	Timestamp                time.Time `json:"Timestamp"`
}

type Fuel struct {
	FTLAmount float64 `json:"ftl_amount"`
	STLAmount float64 `json:"stl_amount"`
}

const endpoint = "/ship"

func GetAll(ctx context.Context, username string, auth *fnar.Authentication) ([]*Ship, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fnar.BaseUrl+endpoint+"/ships/"+username, nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating request for getting all ships")
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

	// setup ship
	shipsMap := []map[string]interface{}{}
	if err := json.Unmarshal(body, &shipsMap); err != nil {
		return nil, errors.Wrap(err, "decoding json response")
	}

	ships := []*Ship{}
	for _, shipMap := range shipsMap {
		ship := &Ship{}
		if err := ship.parse(shipMap); err != nil {
			return nil, errors.Wrap(err, "parse ship")
		}
		ships = append(ships, ship)
	}

	return ships, nil
}

func GetAllShipsFuel(ctx context.Context, username string, auth *fnar.Authentication) (map[string]*Fuel, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fnar.BaseUrl+endpoint+"/ships/fuel/"+username, nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating request for getting all ships")
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

	// setup ship
	shipsFuelMap := []map[string]interface{}{}
	if err := json.Unmarshal(body, &shipsFuelMap); err != nil {
		return nil, errors.Wrap(err, "decoding json response")
	}

	shipsFuels := map[string]*Fuel{}
	for _, shipFuelMap := range shipsFuelMap {
		fuel := &Fuel{}
		amount := shipFuelMap["StorageItems"].([]interface{})[0].(map[string]interface{})["MaterialAmount"].(float64)
		switch shipFuelMap["Type"].(string) {
		case "FTL_FUEL_STORE":
			fuel.FTLAmount = amount
		case "STL_FUEL_STORE":
			fuel.STLAmount = amount
		}
		shipsFuels[shipFuelMap["Name"].(string)] = fuel
	}

	return shipsFuels, nil
}

func (s *Ship) parse(ship map[string]interface{}) error {
	// correct timestamp to time.Time
	var timestamp time.Time
	var err error
	ts, ok := ship["Timestamp"]
	if ok {
		timestamp, err = time.Parse("2006-01-02T15:04:05.999999", ts.(string))
		if err != nil {
			return err
		}
	}
	ship["Timestamp"] = timestamp

	// convert to Ship
	shipM, err := json.Marshal(ship)
	if err != nil {
		return errors.Wrap(err, "marshaling ship map")
	}
	if err := json.Unmarshal(shipM, s); err != nil {
		return errors.Wrap(err, "unmarshaling to Ship")
	}

	return nil
}
