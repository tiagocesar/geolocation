package models

type Geolocation struct {
	IpAddress    string  `csv:"ip_address"`
	CountryCode  string  `csv:"country_code"`
	Country      string  `csv:"country"`
	City         string  `csv:"city"`
	Latitude     float64 `csv:"latitude"`
	Longitude    float64 `csv:"longitude"`
	MysteryValue string  `csv:"mystery_value"`
}
