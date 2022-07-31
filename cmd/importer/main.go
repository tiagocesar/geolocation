package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/tiagocesar/geolocation/internal/models"
	"github.com/tiagocesar/geolocation/internal/repo"
)

/* TODO
-[x] Configure the file importer
-[ ] Configure and start the GRPC server
-[ ] Orchestrate via a compose file
*/

const (
	totalRoutines = 10

	EnvDumpFile = "DUMP_FILE"

	EnvDbUser   = "DB_USER"
	EnvDbPass   = "DB_PASS"
	EnvDbHost   = "DB_HOST"
	EnvDbPort   = "DB_PORT"
	EnvDbSchema = "DB_SCHEMA"
)

type geolocationPersister interface {
	AddLocationInfo(ctx context.Context, locationInfo models.Geolocation) error

	GetLocationInfoByIP(ctx context.Context, ipAddress string) (*models.Geolocation, error)
}

type program struct {
	wg sync.WaitGroup
}

func main() {
	// Getting environment vars
	envVars, err := getEnvVars()
	if err != nil {
		panic(err)
	}

	repository, err := repo.NewRepository(envVars[EnvDbUser], envVars[EnvDbPass], envVars[EnvDbHost],
		envVars[EnvDbPort], envVars[EnvDbSchema])
	if err != nil {
		panic(err)
	}

	fp := NewFileProcessor(repository)

	fp.ExecuteFileImport(envVars[EnvDumpFile])

	// FIXME after the GRPC handler is up, program should wait for an exit signal
}

// getEnvVars gets all environment variables necessary for this service to run.
func getEnvVars() (map[string]string, error) {
	result := make(map[string]string)
	var ok bool

	// Getting the file name with the dump file
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

	return result, nil
}
