package grpc_client

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	pb "github.com/tiagocesar/geolocation/handler/grpc/schema"
)

type grpcClientMock struct {
	GetLocationDataFn func(ctx context.Context, in *pb.LocationRequest) (*pb.LocationResponse, error)
}

func (m *grpcClientMock) GetLocationData(ctx context.Context, in *pb.LocationRequest,
	_ ...grpc.CallOption) (*pb.LocationResponse, error) {

	return m.GetLocationDataFn(ctx, in)
}

func Test_GetLocationData(t *testing.T) {
	tests := []struct {
		name        string
		ipAddress   string
		grpcClient  pb.GeolocationClient
		expectedErr error
	}{
		{
			name:      "success",
			ipAddress: "192.168.0.1",
			grpcClient: &grpcClientMock{
				GetLocationDataFn: func(ctx context.Context, in *pb.LocationRequest) (*pb.LocationResponse, error) {
					return &pb.LocationResponse{}, nil
				},
			},
		},
		{
			name:        "invalid IP address should return error",
			ipAddress:   "a",
			expectedErr: ErrInvalidIP,
		},
		{
			name:      "internal server error should return error",
			ipAddress: "192.168.0.1",
			grpcClient: &grpcClientMock{
				GetLocationDataFn: func(ctx context.Context, in *pb.LocationRequest) (*pb.LocationResponse, error) {
					return nil, errors.New("internal server error")
				},
			},
			expectedErr: errors.New("internal server error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			client := Client{grpcClient: test.grpcClient}

			_, err := client.GetLocationData(context.Background(), test.ipAddress)

			require.Equal(t, test.expectedErr, err)
		})
	}
}
