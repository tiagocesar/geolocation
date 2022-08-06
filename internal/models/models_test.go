//go:build !integration

package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Validate(t *testing.T) {
	tests := []struct {
		name        string
		input       func() *Geolocation
		expectedErr error
	}{
		{
			name: "success",
			input: func() *Geolocation {
				return completeGeolocation()
			},
		},
		{
			name: "missing IP address should return error",
			input: func() *Geolocation {
				g := completeGeolocation()
				g.IpAddress = ""
				return g
			},
			expectedErr: ErrValidationInvalidIP,
		},
		{
			name: "invalid IP address should return error",
			input: func() *Geolocation {
				g := completeGeolocation()
				g.IpAddress = "a.a.a"
				return g
			},
			expectedErr: ErrValidationInvalidIP,
		},
		{
			name: "missing country code should return error",
			input: func() *Geolocation {
				g := completeGeolocation()
				g.CountryCode = ""
				return g
			},
			expectedErr: ErrValidationInvalidCountryCode,
		},
		{
			name: "missing country name should return error",
			input: func() *Geolocation {
				g := completeGeolocation()
				g.Country = ""
				return g
			},
			expectedErr: ErrValidationInvalidCountry,
		},
		{
			name: "missing city should return error",
			input: func() *Geolocation {
				g := completeGeolocation()
				g.City = ""
				return g
			},
			expectedErr: ErrValidationInvalidCity,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			input := test.input()
			err := input.Validate()

			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func completeGeolocation() *Geolocation {
	return &Geolocation{
		IpAddress:   "192.168.0.1",
		CountryCode: "BR",
		Country:     "Brazil",
		City:        "Brasilia",
	}
}
