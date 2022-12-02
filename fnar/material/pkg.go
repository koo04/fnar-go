package material

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/koo04/fnar-go/fnar"
	"github.com/pkg/errors"
)

type Material struct {
	ID                string   `json:"MaterialId"`
	Name              string   `json:"Name"`
	Ticker            string   `json:"Ticker"`
	Weight            float64  `json:"Weight"`
	Volume            float64  `json:"Volume"`
	UserNameSubmitted string   `json:"UserNameSubmitted"`
	Timestamp         string   `json:"Timestamp"`
	Category          Category `json:"Category"`
}

const endpoint = "/material"

func GetAll(ctx context.Context) ([]*Material, error) {
	resp, err := http.Get(fnar.BaseUrl + endpoint + "/allmaterials")
	if err != nil {
		return nil, errors.Wrap(err, "getting all materials")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.Wrap(errors.New(string(b)), "not 200")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading the response body")
	}

	sBody := string(body)
	_ = sBody

	materialsMap := []map[string]interface{}{}
	if err := json.Unmarshal(body, &materialsMap); err != nil {
		return nil, errors.Wrap(err, "decoding json response")
	}

	materials := []*Material{}
	for _, materialMap := range materialsMap {
		material := &Material{}
		if err := material.parse(materialMap); err != nil {
			return nil, errors.Wrap(err, "parse material")
		}
		materials = append(materials, material)
	}

	return materials, nil
}

func (m *Material) parse(msi map[string]interface{}) error {
	// correct timestamp to time.Time
	var timestamp time.Time
	var err error
	ts, ok := msi["Timestamp"]
	if ok {
		timestamp, err = time.Parse("2006-01-02T15:04:05.999999", ts.(string))
		if err != nil {
			return err
		}
	}
	msi["Timestamp"] = timestamp

	// parse the category
	category := &Category{
		Name: msi["CategoryName"].(string),
		Id:   msi["CategoryId"].(string),
	}
	msi["Category"] = category

	// convert to Material
	materialM, err := json.Marshal(msi)
	if err != nil {
		return errors.Wrap(err, "marshaling material map")
	}
	if err := json.Unmarshal(materialM, m); err != nil {
		return errors.Wrap(err, "unmarshaling to Material")
	}

	return nil
}
