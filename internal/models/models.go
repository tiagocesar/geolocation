package models

import (
	"errors"
	"net"
	"strings"
)

var (
	ErrValidationInvalidIP          = errors.New("invalid IP address")
	ErrValidationInvalidCountryCode = errors.New("invalid country code")
	ErrValidationInvalidCountry     = errors.New("invalid country")
	ErrValidationInvalidCity        = errors.New("invalid city")
)

type Geolocation struct {
	IpAddress    string  `csv:"ip_address" json:"ip_address"`
	CountryCode  string  `csv:"country_code" json:"country_code"`
	Country      string  `csv:"country" json:"country"`
	City         string  `csv:"city" json:"city"`
	Latitude     float64 `csv:"latitude" json:"latitude"`
	Longitude    float64 `csv:"longitude" json:"longitude"`
	MysteryValue string  `csv:"mystery_value" json:"mystery_value,omitempty"`
}

func (g Geolocation) Validate() error {
	// Checking if the IP is valid
	ip := net.ParseIP(g.IpAddress)
	if ip == nil {
		return ErrValidationInvalidIP
	}

	if strings.TrimSpace(g.CountryCode) == "" {
		return ErrValidationInvalidCountryCode
	}

	if strings.TrimSpace(g.Country) == "" {
		return ErrValidationInvalidCountry
	}

	if strings.TrimSpace(g.City) == "" {
		return ErrValidationInvalidCity
	}

	return nil
}
