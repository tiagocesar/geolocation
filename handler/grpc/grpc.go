package grpc

import (
	"context"

	pb "github.com/tiagocesar/geolocation/handler/grpc/schema"
	"github.com/tiagocesar/geolocation/internal/models"
)

type geolocationQuerier interface {
	GetLocationInfoByIP(ctx context.Context, ipAddress string) (*models.Geolocation, error)
}

type grpcHandler struct {
	pb.UnimplementedGeolocationServer
	repository geolocationQuerier
}

func (h *grpcHandler) GetLocationData(ctx context.Context, in *pb.LocationRequest) (*pb.LocationResponse, error) {
	location, err := h.repository.GetLocationInfoByIP(ctx, in.GetIp())
	if err != nil {
		return nil, err
	}

	return &pb.LocationResponse{
		Ip:          location.IpAddress,
		CountryCode: location.CountryCode,
		Country:     location.Country,
		City:        location.City,
		Latitude:    location.Latitude,
		Longitude:   location.Longitude,
	}, nil
}
