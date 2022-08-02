package main

import (
	"log"
	"os"

	"github.com/tiagocesar/geolocation/clients/grpc_client"
	"github.com/tiagocesar/geolocation/handler/http"
)

/* TODO
-[ ] Configure and start the HTTP server
-[ ] Orchestrate via a compose file
*/

const (
	EnvHttpServerPort = "HTTP_SERVER_PORT"

	EnvGrpcServerHost = "GRPC_SERVER_HOST"
	EnvGrpcServerPort = "GRPC_SERVER_PORT"
)

func main() {
	var ok bool

	// HTTP server vars
	var httpHost, httpPort string
	if httpPort, ok = os.LookupEnv(EnvHttpServerPort); !ok {
		log.Fatalf("environment variable %s not set", EnvHttpServerPort)
	}

	// GRPC server vars
	var grpcHost, grpcPort string
	if grpcHost, ok = os.LookupEnv(EnvGrpcServerHost); !ok {
		log.Fatalf("environment variable %s not set", EnvGrpcServerHost)
	}

	if grpcPort, ok = os.LookupEnv(EnvGrpcServerPort); !ok {
		log.Fatalf("environment variable %s not set", EnvGrpcServerPort)
	}

	grpcClient, _ := grpc_client.NewClient(grpcHost, grpcPort)

	httpServer := http.NewHttpServer(grpcClient)

	log.Printf("HTTP server starting on %s:%s", httpHost, httpPort)
	httpServer.ConfigureAndServe(httpPort)

	log.Println("HTTP server exiting")
	os.Exit(0)
}
