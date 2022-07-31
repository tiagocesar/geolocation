package models

import (
	"errors"
	"net"
	"strings"
)

type Geolocation struct {
	IpAddress    string  `csv:"ip_address"`
	CountryCode  string  `csv:"country_code"`
	Country      string  `csv:"country"`
	City         string  `csv:"city"`
	Latitude     float64 `csv:"latitude"`
	Longitude    float64 `csv:"longitude"`
	MysteryValue string  `csv:"mystery_value"`
}

func (g Geolocation) Validate() error {
	// Checking if the IP is valid
	ip := net.ParseIP(g.IpAddress)
	if ip == nil {
		return errors.New("invalid IP address")
	}

	if strings.TrimSpace(g.CountryCode) == "" {
		return errors.New("invalid country code")
	}

	if strings.TrimSpace(g.Country) == "" {
		return errors.New("invalid country")
	}

	if strings.TrimSpace(g.City) == "" {
		return errors.New("invalid city")
	}

	return nil
}
