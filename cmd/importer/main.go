package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/gocarina/gocsv"

	"github.com/tiagocesar/geolocation/internal/models"
)

/* TODO
-[ ] Configure the file importer
-[ ] Configure and start the GRPC server
-[ ] Orchestrate via a compose file
*/

type program struct {
	wg sync.WaitGroup
}

func main() {
	filename, dbUser, dbPass, err := getEnvVars()
	if err != nil {
		panic(err)
	}

	p := &program{}

	// Processing the file
	go func(filename string) {
		err := p.processFile(filename)
		if err != nil {
			panic(err)
		}
	}(filename)

	// FIXME connect to the db
	_, _ = dbUser, dbPass
}

func (p *program) processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) { _ = file.Close() }(file)

	totalRoutines := 10
	// FIXME Fire up the goroutines that will process the file contents
	_ = totalRoutines

	header := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if header == "" {
			// First line is the header. We unmarshal just to make sure it's in the expected format
			_, err := csvLineToStruct(header, scanner.Text())
			if err != nil {
				return err
			}
			header = scanner.Text()
			continue
		}

		// Once the header is known we can continue to the proper lines in the CSV
		g, err := csvLineToStruct(header, scanner.Text())
		if err != nil {
			return err
		}

		fmt.Println(g)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func csvLineToStruct(csvHeader, line string) (*models.Geolocation, error) {
	var g models.Geolocation
	if err := gocsv.UnmarshalString(fmt.Sprintf("%s\n%s", csvHeader, line), &g); err != nil {
		return nil, err
	}

	return &g, nil
}

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
