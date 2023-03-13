package planet

import "time"

type Planet struct {
	Resources               []Resource         `json:"Resources"`
	BuildRequirements       []BuildRequirement `json:"BuildRequirements"`
	ProductionFees          []ProductionFee    `json:"ProductionFees"`
	COGCPrograms            []interface{}      `json:"COGCPrograms"`
	COGCVotes               []interface{}      `json:"COGCVotes"`
	COGCUpkeep              []interface{}      `json:"COGCUpkeep"`
	PlanetID                string             `json:"PlanetId"`
	PlanetNaturalID         string             `json:"PlanetNaturalId"`
	PlanetName              string             `json:"PlanetName"`
	Namer                   interface{}        `json:"Namer"`
	NamingDataEpochMs       int                `json:"NamingDataEpochMs"`
	Nameable                bool               `json:"Nameable"`
	SystemID                string             `json:"SystemId"`
	Gravity                 float64            `json:"Gravity"`
	MagneticField           float64            `json:"MagneticField"`
	Mass                    int64              `json:"Mass"`
	MassEarth               float64            `json:"MassEarth"`
	OrbitSemiMajorAxis      float64            `json:"OrbitSemiMajorAxis"`
	OrbitEccentricity       float64            `json:"OrbitEccentricity"`
	OrbitInclination        float64            `json:"OrbitInclination"`
	OrbitRightAscension     float64            `json:"OrbitRightAscension"`
	OrbitPeriapsis          float64            `json:"OrbitPeriapsis"`
	OrbitIndex              int                `json:"OrbitIndex"`
	Pressure                float64            `json:"Pressure"`
	Radiation               float64            `json:"Radiation"`
	Radius                  float64            `json:"Radius"`
	Sunlight                float64            `json:"Sunlight"`
	Surface                 bool               `json:"Surface"`
	Temperature             float64            `json:"Temperature"`
	Fertility               int                `json:"Fertility"`
	HasLocalMarket          bool               `json:"HasLocalMarket"`
	HasChamberOfCommerce    bool               `json:"HasChamberOfCommerce"`
	HasWarehouse            bool               `json:"HasWarehouse"`
	HasAdministrationCenter bool               `json:"HasAdministrationCenter"`
	HasShipyard             bool               `json:"HasShipyard"`
	FactionCode             interface{}        `json:"FactionCode"`
	FactionName             interface{}        `json:"FactionName"`
	GovernorID              interface{}        `json:"GovernorId"`
	GovernorUserName        interface{}        `json:"GovernorUserName"`
	GovernorCorporationID   interface{}        `json:"GovernorCorporationId"`
	GovernorCorporationName interface{}        `json:"GovernorCorporationName"`
	GovernorCorporationCode interface{}        `json:"GovernorCorporationCode"`
	CurrencyName            interface{}        `json:"CurrencyName"`
	CurrencyCode            interface{}        `json:"CurrencyCode"`
	CollectorID             interface{}        `json:"CollectorId"`
	CollectorName           interface{}        `json:"CollectorName"`
	CollectorCode           interface{}        `json:"CollectorCode"`
	BaseLocalMarketFee      float64            `json:"BaseLocalMarketFee"`
	LocalMarketFeeFactor    float64            `json:"LocalMarketFeeFactor"`
	WarehouseFee            float64            `json:"WarehouseFee"`
	PopulationID            string             `json:"PopulationId"`
	COGCProgramStatus       interface{}        `json:"COGCProgramStatus"`
	PlanetTier              int                `json:"PlanetTier"`
	UserNameSubmitted       string             `json:"UserNameSubmitted"`
	Timestamp               time.Time          `json:"Timestamp"`
}
