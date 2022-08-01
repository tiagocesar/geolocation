package grpc

import (
	"context"
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
		name       string
		repository *mockRepository
	}{
		{
			name: "success",
			repository: &mockRepository{
				GetLocationInfoByIPFn: func(s context.Context, ipAddress string) (*models.Geolocation, error) {
					return &models.Geolocation{}, nil
				},
			},
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

			require.NoError(t, err)
			require.Equal(t, result, result)
		})
	}
}
