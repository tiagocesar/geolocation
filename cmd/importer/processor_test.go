package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tiagocesar/geolocation/internal/models"
)

type mockRepository struct {
	AddLocationInfoInvokedCount int
	AddLocationInfoFn           func(ctx context.Context, locationInfo models.Geolocation) error
}

func (m *mockRepository) AddLocationInfo(ctx context.Context, locationInfo models.Geolocation) error {
	m.AddLocationInfoInvokedCount++

	return m.AddLocationInfoFn(ctx, locationInfo)
}

func Test_persistGeoData(t *testing.T) {
	t.Run("Successful persistence - asserts AddLocationInfo was called exactly once", func(t *testing.T) {
		t.Parallel()

		repository := mockRepository{
			AddLocationInfoFn: func(ctx context.Context, locationInfo models.Geolocation) error {
				return nil
			},
		}
		fp := NewFileProcessor(&repository)

		go fp.persistGeoData()

		fp.data <- mockGeolocation()

		assert.True(t, repository.AddLocationInfoInvokedCount == 1)
	})

	t.Run("invalid data - atomic invalid count should increment", func(t *testing.T) {
		t.Parallel()

		repository := &mockRepository{}
		fp := NewFileProcessor(repository)

		assert.True(t, fp.invalidLines == 0)

		go fp.persistGeoData()

		fp.data <- models.Geolocation{}

		assert.True(t, repository.AddLocationInfoInvokedCount == 0)
		assert.True(t, fp.invalidLines == 1)
	})
}

func mockGeolocation() models.Geolocation {
	return models.Geolocation{
		IpAddress:    "192.168.0.1",
		CountryCode:  "BR",
		Country:      "Brazil",
		City:         "Brasilia",
		Latitude:     0,
		Longitude:    0,
		MysteryValue: "Home sweet home",
	}
}
