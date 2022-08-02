package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"

	pb "github.com/tiagocesar/geolocation/handler/grpc/schema"
	"github.com/tiagocesar/geolocation/internal/models"
)

type mockGrpcClient struct {
	GetLocationDataFn func(ctx context.Context, ip string) (*pb.LocationResponse, error)
}

func (m *mockGrpcClient) GetLocationData(ctx context.Context, ip string) (*pb.LocationResponse, error) {
	return m.GetLocationDataFn(ctx, ip)
}

func TestHandler_getGeolocationData(t *testing.T) {
	tests := []struct {
		name             string
		grpcClientMock   locationFinder
		ipAddress        string
		expectedRespCode int
		expectedRespBody func() string
	}{
		{
			name: "success",
			grpcClientMock: &mockGrpcClient{
				GetLocationDataFn: func(ctx context.Context, ip string) (*pb.LocationResponse, error) {
					return &pb.LocationResponse{
						Ip:          "192.168.0.1",
						CountryCode: "ZZZ",
						Country:     "Unit Tests",
					}, nil
				},
			},
			ipAddress:        "192.168.0.1",
			expectedRespCode: http.StatusOK,
			expectedRespBody: func() string {
				ev := models.Geolocation{
					IpAddress:   "192.168.0.1",
					CountryCode: "ZZZ",
					Country:     "Unit Tests",
				}

				s, _ := json.Marshal(ev)
				return string(s)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			rr := httptest.NewRecorder()

			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("ip", test.ipAddress)

			req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/locations/{ip}", nil)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
			require.NoError(t, err)

			h := httpServer{
				grpcClient: test.grpcClientMock,
			}

			h.getGeolocationData(rr, req)

			require.Equal(t, test.expectedRespCode, rr.Code)
			resp := test.expectedRespBody()
			require.Equal(t, resp, rr.Body.String())
		})
	}
}
