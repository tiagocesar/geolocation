package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gocarina/gocsv"

	"github.com/tiagocesar/geolocation/internal/models"
)

type geolocationPersister interface {
	AddLocationInfo(ctx context.Context, locationInfo models.Geolocation) error
}

type fileProcessor struct {
	wg           sync.WaitGroup
	data         chan models.Geolocation
	totalLines   uint64
	invalidLines uint64

	repository geolocationPersister
}

func NewFileProcessor(repository geolocationPersister) *fileProcessor {
	return &fileProcessor{
		data:       make(chan models.Geolocation),
		repository: repository,
	}
}

func (p *fileProcessor) ExecuteFileImport(dumpFile string) {
	startTime := time.Now()

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

			p.persistGeoData()
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
	}(dumpFile)

	p.wg.Wait()

	fmt.Printf("Elapsed time: %s\n", time.Since(startTime))
	fmt.Printf("Total lines: %d, invalid lines: %d\n", p.totalLines, p.invalidLines)
}

// processFile opens the file specified in the DUMP_FILE environment var, checks if it's valid against the defined
// csv schema (defined by the header) and sends each line in the CSV for async processing.
//
// The actual contents of each line (after being converted to a models.Geolocation struct) is validated before
// persisting it.
func (p *fileProcessor) processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) { _ = file.Close() }(file)

	header := ""
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
			p.incrementInvalidCount()
			continue
		}

		p.data <- g
	}

	close(p.data)

	if err := scanner.Err(); err != nil {
		return err
	}

	p.wg.Done()

	return nil
}

// persistGeoData validates and persists geolocation data, feeding invalidLines via an atomic operation
func (p *fileProcessor) persistGeoData() {
	for g := range p.data {
		// Checking if the data is valid
		if err := g.Validate(); err != nil {
			// Incrementing one on the list of total errors
			p.incrementInvalidCount()
			continue
		}

		err := p.repository.AddLocationInfo(context.Background(), g)
		if err != nil {
			p.incrementInvalidCount()
		}
	}
}

func (p *fileProcessor) incrementInvalidCount() {
	atomic.AddUint64(&p.invalidLines, 1)
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
