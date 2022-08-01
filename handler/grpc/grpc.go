package grpc

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

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

func NewGrpcServer(port string, repository geolocationQuerier) (*net.Listener, *grpc.Server, error) {
	handler := &grpcHandler{
		repository: repository,
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return nil, nil, fmt.Errorf("grpc server - failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGeolocationServer(grpcServer, handler)
	reflection.Register(grpcServer)

	return &lis, grpcServer, nil
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
