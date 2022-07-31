package models

type Geolocation struct {
	IpAddress    string  `json:"ip_address"`
	CountryCode  string  `json:"country_code"`
	Country      string  `json:"country"`
	City         string  `json:"city"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	MysteryValue string  `json:"mystery_value"`
}
