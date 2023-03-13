package company

import (
	"github.com/koo04/fnar-go/fnar/country"
	"github.com/koo04/fnar-go/fnar/planet"
)

type Company struct {
	Balances                       interface{}     `json:"Balances"`
	Planets                        []planet.Planet `json:"Planets"`
	Offices                        []interface{}   `json:"Offices"`
	UserDataID                     string          `json:"UserDataId"`
	UserID                         string          `json:"UserId"`
	UserName                       string          `json:"UserName"`
	SubscriptionLevel              string          `json:"SubscriptionLevel"`
	Tier                           interface{}     `json:"Tier"`
	Team                           bool            `json:"Team"`
	Pioneer                        bool            `json:"Pioneer"`
	Moderator                      bool            `json:"Moderator"`
	CreatedEpochMs                 int64           `json:"CreatedEpochMs"`
	ID                             string          `json:"CompanyId"`
	Name                           string          `json:"CompanyName"`
	Code                           string          `json:"CompanyCode"`
	Country                        country.Country `json:"Country"`
	CorporationID                  string          `json:"CorporationId"`
	CorporationName                string          `json:"CorporationName"`
	CorporationCode                string          `json:"CorporationCode"`
	OverallRating                  string          `json:"OverallRating"`
	ActivityRating                 string          `json:"ActivityRating"`
	ReliabilityRating              string          `json:"ReliabilityRating"`
	StabilityRating                string          `json:"StabilityRating"`
	HeadquartersNaturalID          interface{}     `json:"HeadquartersNaturalId"`
	HeadquartersLevel              int             `json:"HeadquartersLevel"`
	HeadquartersBasePermits        int             `json:"HeadquartersBasePermits"`
	HeadquartersUsedBasePermits    int             `json:"HeadquartersUsedBasePermits"`
	AdditionalBasePermits          int             `json:"AdditionalBasePermits"`
	AdditionalProductionQueueSlots int             `json:"AdditionalProductionQueueSlots"`
	RelocationLocked               bool            `json:"RelocationLocked"`
	NextRelocationTimeEpochMs      int             `json:"NextRelocationTimeEpochMs"`
	UserNameSubmitted              string          `json:"UserNameSubmitted"`
	Timestamp                      string          `json:"Timestamp"`
}
