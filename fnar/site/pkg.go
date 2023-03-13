package site

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
	Id     string `json:"MaterialId"`
	Name   string `json:"MaterialName"`
	Ticker string `json:"MaterialTicker"`
	Amount int32  `json:"MaterialAmount"`
}

type Building struct {
	ReclaimableMaterials []*Material `json:"ReclaimableMaterials"`
	RepairMaterials      []*Material `json:"RepairMaterials"`
	SiteId               string      `json:"SiteBuildingId"`
	Id                   string      `json:"BuildingId"`
	Created              time.Time   `json:"BuildingCreated"`
	Name                 string      `json:"BuildingName"`
	Ticker               string      `json:"BuildingTicker"`
	LastRepair           time.Time   `json:"LastRepair"`
	Condition            float64     `json:"Condition"`
}

type Site struct {
	Id               string `json:"SiteId"`
	PlanetId         string `json:"PlanetId"`
	PlanetIdentifier string `json:"PlanetIdentifier"`
	PlanetName       string `json:"PlanetName"`
	// PlanetFoundedEpocMs int64 `json:"PlanetFoundedEpocMs"` This needs to be looked at
	InvestedPermits   int32       `json:"InvestedPermits"`
	MaximumPermits    int32       `json:"MaximumPermits"`
	UserNameSubmitted string      `json:"UserNameSubmitted"`
	Timestamp         time.Time   `json:"Timestamp"`
	Buildings         []*Building `json:"Buildings"`
}

const endpoint = "/sites"

func GetAll(ctx context.Context, username string, auth *fnar.Authentication) ([]*Site, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fnar.BaseUrl+endpoint+"/"+username, nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating request for getting all sites")
	}
	req.Header.Add("Authorization", auth.AuthToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "getting all sites")
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

	// setup sites
	sitesMap := []map[string]interface{}{}
	if err := json.Unmarshal(body, &sitesMap); err != nil {
		return nil, errors.Wrap(err, "decoding json response")
	}

	sites := []*Site{}
	for _, siteMap := range sitesMap {
		site := &Site{}
		if err := site.parse(siteMap); err != nil {
			return nil, errors.Wrap(err, "parse site")
		}
		sites = append(sites, site)
	}

	return sites, nil
}

func (s *Site) parse(site map[string]interface{}) error {
	// correct timestamp to time.Time
	var timestamp time.Time
	var err error
	ts, ok := site["Timestamp"]
	if ok {
		timestamp, err = time.Parse("2006-01-02T15:04:05.999999", ts.(string))
		if err != nil {
			return err
		}
	}
	site["Timestamp"] = timestamp

	// update the buildings timestamps
	buildings, ok := site["Buildings"].([]interface{})
	if ok {
		for k, v := range buildings {
			bc, ok := v.(map[string]interface{})["BuildingCreated"].(float64)
			if ok {
				v.(map[string]interface{})["BuildingCreated"] = time.Unix(0, int64(bc)*int64(time.Millisecond))
			}

			blr, ok := v.(map[string]interface{})["BuildingLastRepair"].(float64)
			if ok {
				v.(map[string]interface{})["BuildingLastRepair"] = time.Unix(0, int64(blr)*int64(time.Millisecond))
			}

			buildings[k] = v
		}
	}

	// convert to Site
	siteM, err := json.Marshal(site)
	if err != nil {
		return errors.Wrap(err, "marshaling site map")
	}
	if err := json.Unmarshal(siteM, s); err != nil {
		return errors.Wrap(err, "unmarshaling to Site")
	}

	return nil
}
