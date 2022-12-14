package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/tiagocesar/geolocation/handler/grpc"
	"github.com/tiagocesar/geolocation/internal/processor"
	"github.com/tiagocesar/geolocation/internal/repo"
)

const (
	totalRoutines = 10

	EnvDumpFile = "DUMP_FILE"

	EnvDbUser   = "DB_USER"
	EnvDbPass   = "DB_PASS"
	EnvDbHost   = "DB_HOST"
	EnvDbPort   = "DB_PORT"
	EnvDbSchema = "DB_SCHEMA"

	EnvGrpcServerPort = "GRPC_SERVER_PORT"
)

func main() {
	var wg sync.WaitGroup

	// Getting environment vars
	envVars, err := getEnvVars()
	if err != nil {
		log.Fatal(err)
	}

	// Configuring access to the repository and opening the SQL connection
	repository, err := repo.NewRepository(envVars[EnvDbUser], envVars[EnvDbPass], envVars[EnvDbHost],
		envVars[EnvDbPort], envVars[EnvDbSchema])
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	// Importing the dump file to the data store
	go func() {
		defer wg.Done()

		processor.NewFileProcessor(repository).ExecuteFileImport(context.Background(), envVars[EnvDumpFile], totalRoutines)
	}()

	// Starting the GRPC server with signals to gracefully stop it
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	listener, grpcServer, err := grpc.NewGrpcServer(envVars[EnvGrpcServerPort], repository)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	go func() {
		s := <-sigCh
		log.Printf("Got signal %v, stopping server\n", s)
		grpcServer.GracefulStop()
		wg.Done()
	}()

	log.Println("Starting GRPC server")
	err = grpcServer.Serve(*listener)
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()

	log.Println("Shutdown successful")
}

// getEnvVars gets all environment variables necessary for this service to run.
func getEnvVars() (map[string]string, error) {
	result := make(map[string]string)
	var ok bool

	// File name for the dump file
	if result[EnvDumpFile], ok = os.LookupEnv(EnvDumpFile); !ok {
		return nil, errors.New(fmt.Sprintf("environment variable %s not set", EnvDumpFile))
	}

	// DB vars
	if result[EnvDbUser], ok = os.LookupEnv(EnvDbUser); !ok {
		return nil, errors.New(fmt.Sprintf("environment variable %s not set", EnvDbUser))
	}

	if result[EnvDbPass], ok = os.LookupEnv(EnvDbPass); !ok {
		return nil, errors.New(fmt.Sprintf("environment variable %s not set", EnvDbPass))
	}

	if result[EnvDbHost], ok = os.LookupEnv(EnvDbHost); !ok {
		return nil, errors.New(fmt.Sprintf("environment variable %s not set", EnvDbHost))
	}

	if result[EnvDbPort], ok = os.LookupEnv(EnvDbPort); !ok {
		return nil, errors.New(fmt.Sprintf("environment variable %s not set", EnvDbPort))
	}

	if result[EnvDbSchema], ok = os.LookupEnv(EnvDbSchema); !ok {
		return nil, errors.New(fmt.Sprintf("environment variable %s not set", EnvDbSchema))
	}

	// GRPC server
	if result[EnvGrpcServerPort], ok = os.LookupEnv(EnvGrpcServerPort); !ok {
		return nil, errors.New(fmt.Sprintf("environment variable %s not set", EnvGrpcServerPort))
	}

	return result, nil
}
