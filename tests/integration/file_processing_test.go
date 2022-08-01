//go:build integration

package integration

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tiagocesar/geolocation/internal/processor"
	"github.com/tiagocesar/geolocation/internal/repo"
)

const (
	EnvDbUser   = "DB_USER"
	EnvDbPass   = "DB_PASS"
	EnvDbHost   = "DB_HOST"
	EnvDbPort   = "DB_PORT"
	EnvDbSchema = "DB_SCHEMA"
)

// Test_NewFileProcessor will do integration testing for the file processing package.
// This acts on some assumptions based on the sample file available on this folder
// (data_dump_sample.csv):
//
// Total lines to process: 10
// Invalid lines: 3
func Test_NewFileProcessor(t *testing.T) {
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

	ctx := context.Background()
	fp := processor.NewFileProcessor(repository)

	fp.ExecuteFileImport(ctx, "data_dump_sample.csv", 10)

	// Based on the sample file:
	// Total lines to process: 10
	// Invalid lines: 3
	assert.Equal(t, uint64(10), fp.TotalLines)
	assert.Equal(t, uint64(3), fp.InvalidLines)

	// Configuring a new repository - with testing methods - to check the data that was inserted
	testRepository, err := NewRepositoryForIntegrationTesting(envVars[EnvDbUser], envVars[EnvDbPass], envVars[EnvDbHost],
		envVars[EnvDbPort], envVars[EnvDbSchema])
	if err != nil {
		log.Fatal(err)
	}

	rows, err := testRepository.GetTestRows(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Based on the sample file:
	// Total lines to process: 10
	// Invalid lines: 3
	// Lines that should have been inserted: 7
	assert.True(t, len(rows) == 7)

	// Cleaning up the db
	err = testRepository.CleanDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

// getEnvVars gets all environment variables necessary for this service to run.
func getEnvVars() (map[string]string, error) {
	result := make(map[string]string)
	var ok bool

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
