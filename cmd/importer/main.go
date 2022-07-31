package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gocarina/gocsv"

	"github.com/tiagocesar/geolocation/internal/models"
)

/* TODO
-[ ] Configure the file importer
-[ ] Configure and start the GRPC server
-[ ] Orchestrate via a compose file
*/

const totalRoutines = 5

type program struct {
	wg           sync.WaitGroup
	data         chan models.Geolocation
	totalLines   uint64
	invalidLines uint64
}

func main() {
	filename, dbUser, dbPass, err := getEnvVars()
	if err != nil {
		panic(err)
	}

	// FIXME connect to the db
	_, _ = dbUser, dbPass

	p := &program{
		data: make(chan models.Geolocation),
	}

	p.wg.Add(1) // We need to wait for the file processor + geolocation persistence routines

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
	}(filename)

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

// saveGeoData validates and persists geolocation data, returning invalidLines to a channel
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
func getEnvVars() (string, string, string, error) {
	var filename, dbUser, dbPass string
	var ok bool

	// Getting the file name with the dump file
	if filename, ok = os.LookupEnv("DUMP_FILE"); !ok {
		return "", "", "", errors.New("environment variable DUMP_FILE not set")
	}

	// DB vars
	if dbUser, ok = os.LookupEnv("DB_USER"); !ok {
		return "", "", "", errors.New("environment variable DB_USER not set")
	}

	if dbPass, ok = os.LookupEnv("DB_PASS"); !ok {
		return "", "", "", errors.New("environment variable DB_PASS not set")
	}

	return filename, dbUser, dbPass, nil
}
