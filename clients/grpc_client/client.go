package grpc_client

import (
	"context"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"

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
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", c.host, c.port))
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

	data, err := client.GetLocationData(ctx, &pb.LocationRequest{Ip: ipAddress.String()})
	if err != nil {
		return nil, err
	}

	return data, nil
}
