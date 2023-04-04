package workforce

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/koo04/fnar-go/fnar"
	"github.com/pkg/errors"
)

type Need struct {
	Category           string  `json:"Category"`
	Essential          bool    `json:"Essential"`
	MaterialId         string  `json:"MaterialId"`
	MaterialName       string  `json:"MaterialName"`
	MaterialTicket     string  `json:"MaterialTicker"`
	Satisfaction       float32 `json:"Satisfaction"`
	UnitsPerInterval   float32 `json:"UnitsPerInterval"`
	UnitsPerOneHundred float32 `json:"UnitsPerOneHundred"`
}

type Workforce struct {
	Needs           []*Need   `json:"WorkforceNeeds"`
	TypeName        string    `json:"WorkforceTypeName"`
	Population      int32     `json:"Population"`
	Reserve         int32     `json:"Reserve"`
	Capacity        int32     `json:"Capacity"`
	Required        int32     `json:"Required"`
	Satisfaction    float32   `json:"Satisfaction"`
	PlanetId        string    `json:"PlanetId"`
	PlanetNaturalId string    `json:"PlanetNaturalId"`
	PlanetName      string    `json:"PlanetName"`
	TimeStamp       time.Time `json:"TimeStamp"`
}

const endpoint = "/workforce"

func GetAll(ctx context.Context, username string, auth *fnar.Authentication) ([]*Workforce, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fnar.BaseUrl+endpoint+"/"+username, nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating request for getting all workforces")
	}
	req.Header.Add("Authorization", auth.AuthToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "getting all Workforces")
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

	// setup workforces
	workforcesMap := []map[string]interface{}{}
	if err := json.Unmarshal(body, &workforcesMap); err != nil {
		return nil, errors.Wrap(err, "decoding json response")
	}

	// compress the information
	workforces := []*Workforce{}
	for _, workforceMap := range workforcesMap {
		for _, wfs := range workforceMap["Workforces"].([]interface{}) {
			wfsBytes, err := json.Marshal(wfs)
			if err != nil {
				return nil, errors.Wrap(err, "marshaling workforces")
			}

			wf := &Workforce{}
			if err := wf.parse(workforceMap); err != nil {
				return nil, errors.Wrap(err, "parsing workforce")
			}

			if err := json.Unmarshal(wfsBytes, &wf); err != nil {
				return nil, errors.Wrap(err, "unmarshaling to Workforce")
			}

			workforces = append(workforces, wf)
		}
	}

	return workforces, nil
}

func (w *Workforce) parse(workforce map[string]interface{}) error {
	// correct timestamp to time.Time
	var timestamp time.Time
	var err error
	ts, ok := workforce["Timestamp"].(string)
	if ok {
		timestamp, err = time.Parse("2006-01-02T15:04:05.999999", ts)
		if err != nil {
			return err
		}
	}
	workforce["Timestamp"] = timestamp

	// convert to Workforce
	workforceM, err := json.Marshal(workforce)
	if err != nil {
		return errors.Wrap(err, "marshaling site map")
	}
	if err := json.Unmarshal(workforceM, w); err != nil {
		return errors.Wrap(err, "unmarshaling to Workforce")
	}

	return nil
}
