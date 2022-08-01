//go:build !integration

package grpc

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	pb "github.com/tiagocesar/geolocation/handler/grpc/schema"
	"github.com/tiagocesar/geolocation/internal/models"
)

type mockRepository struct {
	GetLocationInfoByIPFn func(ctx context.Context, ipAddress string) (*models.Geolocation, error)
}

func (m *mockRepository) GetLocationInfoByIP(ctx context.Context, ipAddress string) (*models.Geolocation, error) {
	return m.GetLocationInfoByIPFn(ctx, ipAddress)
}

func Test_GetLocationData(t *testing.T) {
	tests := []struct {
		name          string
		repository    *mockRepository
		expectedModel *models.Geolocation
		expectedError error
	}{
		{
			name: "success",
			repository: &mockRepository{
				GetLocationInfoByIPFn: func(s context.Context, ipAddress string) (*models.Geolocation, error) {
					return mockGeolocation(), nil
				},
			},
			expectedModel: mockGeolocation(),
		},
		{
			name: "no row found - sql.ErrNoRows should not return an error from the GRPC server",
			repository: &mockRepository{
				GetLocationInfoByIPFn: func(ctx context.Context, ipAddress string) (*models.Geolocation, error) {
					return &models.Geolocation{}, sql.ErrNoRows
				},
			},
			expectedError: nil,
		},
		{
			name: "an unexpected error happened",
			repository: &mockRepository{
				GetLocationInfoByIPFn: func(ctx context.Context, ipAddress string) (*models.Geolocation, error) {
					return nil, errors.New("random error")
				},
			},
			expectedError: errors.New("random error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			handler := &grpcHandler{
				repository: test.repository,
			}

			in := &pb.LocationRequest{Ip: ""}
			result, err := handler.GetLocationData(context.Background(), in)

			require.Equal(t, test.expectedError, err)
			if result != nil {
				location := test.expectedModel
				require.Equal(t, location.IpAddress, result.Ip)
				require.Equal(t, location.CountryCode, result.CountryCode)
				require.Equal(t, location.Country, result.Country)
				require.Equal(t, location.City, result.City)
				require.Equal(t, location.Latitude, result.Latitude)
				require.Equal(t, location.Longitude, result.Longitude)
			}
		})
	}
}

func mockGeolocation() *models.Geolocation {
	return &models.Geolocation{
		IpAddress:    "192.168.0.1",
		CountryCode:  "BR",
		Country:      "Brazil",
		City:         "Brasilia",
		Latitude:     0,
		Longitude:    0,
		MysteryValue: "Home sweet home",
	}
}
