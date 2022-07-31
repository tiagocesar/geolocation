package grpc

import (
	"context"

	pb "github.com/tiagocesar/geolocation/handler/grpc/schema"
)

type grpcHandler struct {
	pb.UnimplementedGeolocationServer
}

func (h *grpcHandler) GetLocationData(ctx context.Context, in *pb.LocationRequest) (*pb.LocationResponse, error) {

	return nil, nil
}
