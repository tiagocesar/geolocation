package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gocarina/gocsv"

	"github.com/tiagocesar/geolocation/internal/models"
	"github.com/tiagocesar/geolocation/internal/repo"
)

/* TODO
-[x] Configure the file importer
-[ ] Configure and start the GRPC server
-[ ] Orchestrate via a compose file
*/

const (
	totalRoutines = 5

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
	wg           sync.WaitGroup
	data         chan models.Geolocation
	totalLines   uint64
	invalidLines uint64

	repository geolocationPersister
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

	p := &program{
		data:       make(chan models.Geolocation),
		repository: repository,
	}

	p.wg.Add(1)

	// Running the goroutines that will persist valid lines
	for i := 0; i < totalRoutines; i++ {
		p.wg.Add(1)

		go func() {
			defer p.wg.Done()

			defer func() {
				if r := recover(); r != nil {
					fmt.Println(fmt.Errorf("recovered from panic: %e", r))
				}
			}()

			p.saveGeoData()
		}()
	}

	// Processing the file
	go func(filename string) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(fmt.Errorf("recovered from panic: %e", r))
			}
		}()

		err := p.processFile(filename)
		if err != nil {
			panic(err)
		}
	}(envVars[EnvDumpFile])

	p.wg.Wait()

	// FIXME after the GRPC handler is up, program should wait for an exit signal
	fmt.Println(p)
}

// processFile opens the file specified in the DUMP_FILE environment var, checks if it's valid against the defined
// csv schema (defined by the header) and sends each line in the CSV for async processing.
//
// The actual contents of each line (after being converted to a models.Geolocation struct) is validated before
// persisting it.
func (p *program) processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) { _ = file.Close() }(file)

	header := ""
	startTime := time.Now()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if header == "" {
			// First line is the header.
			header = scanner.Text()
			continue
		}

		p.totalLines++

		// Once the header is known we can continue to the proper lines in the CSV
		g, err := csvLineToStruct(header, scanner.Text())
		if err != nil {
			atomic.AddUint64(&p.invalidLines, 1)
			continue
		}

		p.data <- g
	}

	close(p.data)

	if err := scanner.Err(); err != nil {
		return err
	}

	// TODO show time in a nicer way
	fmt.Println(time.Since(startTime))

	p.wg.Done()

	return nil
}

// saveGeoData validates and persists geolocation data, feeding invalidLines via an atomic operation
func (p *program) saveGeoData() {
	for g := range p.data {
		// Checking if the data is valid
		if err := g.Validate(); err != nil {
			// Incrementing one on the list of total errors
			atomic.AddUint64(&p.invalidLines, 1)
			continue
		}

		fmt.Println(g)
	}
}

// csvLineToStruct converts each CSV line to a models.Geolocation struct, given the header of the CSV file
// and the line contents.
func csvLineToStruct(header, line string) (models.Geolocation, error) {
	var g []models.Geolocation
	if err := gocsv.UnmarshalString(fmt.Sprintf("%s\n%s", header, line), &g); err != nil {
		return models.Geolocation{}, err
	}

	// gocsv returns an array even if there's only one value
	return g[0], nil
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
