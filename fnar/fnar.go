package fnar

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type Authentication struct {
	AuthToken       string    `json:"AuthToken"`
	Expiry          time.Time `json:"Expiry"`
	IsAdministrator bool      `json:"IsAdministrator"`
}

const BaseUrl = "https://rest.fnar.net"

func Login(username string, password string) (*Authentication, error) {
	values := map[string]string{"UserName": username, "Password": password}
	valuesJson, err := json.Marshal(values)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling values")
	}
	resp, err := http.Post(BaseUrl+"/auth/login", "application/json", bytes.NewBuffer(valuesJson))
	if err != nil {
		return nil, errors.Wrap(err, "logging in")
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

	authenticationMap := map[string]interface{}{}
	if err := json.Unmarshal(body, &authenticationMap); err != nil {
		return nil, errors.Wrap(err, "marshaling body")
	}

	auth := &Authentication{}
	if err := auth.parse(authenticationMap); err != nil {
		return nil, errors.Wrap(err, "parsing authentication")
	}

	return auth, nil
}

func (a *Authentication) parse(auth map[string]interface{}) error {
	// correct timestamp to time.Time
	var timestamp time.Time
	var err error
	ts, ok := auth["Expiry"]
	if ok {
		timestamp, err = time.Parse("2006-01-02T15:04:05.999999", ts.(string))
		if err != nil {
			return err
		}
	}
	auth["Expiry"] = timestamp

	// convert to Authentication
	authenticationM, err := json.Marshal(auth)
	if err != nil {
		return errors.Wrap(err, "marshaling authentication map")
	}
	if err := json.Unmarshal(authenticationM, a); err != nil {
		return errors.Wrap(err, "unmarshaling to Authentication")
	}

	return nil
}
