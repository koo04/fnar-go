package inventory

type Type int32

const (
	Unknown Type = iota
	FTLFuelStore
	STLFuelStore
	WarehouseStore
	ShipStore
	Store
)

func (t Type) String() string {
	return [...]string{"UNKNOWN", "FTL_FUEL_STORE", "STL_FUEL_STORE", "WAREHOUSE_STORE", "SHIP_STORE", "STORE"}[t-1]
}

func (t Type) EnumIndex() int {
	return int(t)
}

func ToType(s string) Type {
	switch s {
	case "FTL_FUEL_STORE":
		return FTLFuelStore
	case "STL_FUEL_STORE":
		return STLFuelStore
	case "WAREHOUSE_STORE":
		return WarehouseStore
	case "SHIP_STORE":
		return ShipStore
	case "STORE":
		return Store
	}
	return Unknown
}
