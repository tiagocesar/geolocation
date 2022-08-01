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
	host string
	port string
}

func NewClient(host, port string) *Client {
	return &Client{
		host: host,
		port: port,
	}
}

func (c *Client) getConnection() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", c.host, c.port), opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *Client) GetLocationData(ctx context.Context, ip string) (*pb.LocationResponse, error) {
	// Checking if the IP is valid
	ipAddress := net.ParseIP(ip)
	if ipAddress == nil {
		return nil, ErrInvalidIP
	}

	conn, err := c.getConnection()
	defer func(conn *grpc.ClientConn) { _ = conn.Close() }(conn)

	client := pb.NewGeolocationClient(conn)

	req := &pb.LocationRequest{Ip: ipAddress.String()}
	data, err := client.GetLocationData(ctx, req)
	if err != nil {
		return nil, err
	}

	return data, nil
}
