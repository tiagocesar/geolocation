package grpc_client

import (
	"context"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/tiagocesar/geolocation/handler/grpc/schema"
)

var ErrInvalidIP = errors.New("invalid IP address")

type Client struct {
	grpcClient pb.GeolocationClient
}

func NewClient(host, port string) (*Client, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), opts...)
	if err != nil {
		return nil, err
	}

	client := pb.NewGeolocationClient(conn)

	return &Client{
		grpcClient: client,
	}, nil
}

func (c *Client) GetLocationData(ctx context.Context, ip string) (*pb.LocationResponse, error) {
	// Checking if the IP is valid
	ipAddress := net.ParseIP(ip)
	if ipAddress == nil {
		return nil, ErrInvalidIP
	}

	req := &pb.LocationRequest{Ip: ipAddress.String()}
	data, err := c.grpcClient.GetLocationData(ctx, req)
	if err != nil {
		return nil, err
	}

	return data, nil
}
