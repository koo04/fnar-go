package country

type Country struct {
	RegistryCountryID   string `json:"CountryRegistryCountryId"`
	Code                string `json:"CountryCode"`
	Name                string `json:"CountryName"`
	CurrencyNumericCode int    `json:"CurrencyNumericCode"`
	CurrencyCode        string `json:"CurrencyCode"`
	CurrencyName        string `json:"CurrencyName"`
	CurrencyDecimals    int    `json:"CurrencyDecimals"`
}
