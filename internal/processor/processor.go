package processor

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"runtime/debug"
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
	wg            sync.WaitGroup
	data          chan models.Geolocation
	TotalLines    uint64
	AcceptedLines uint64
	InvalidLines  uint64

	repository geolocationPersister
}

func NewFileProcessor(repository geolocationPersister) *fileProcessor {
	return &fileProcessor{
		data:       make(chan models.Geolocation),
		repository: repository,
	}
}

func (fp *fileProcessor) ExecuteFileImport(ctx context.Context, dumpFile string, totalRoutines int) {
	startTime := time.Now()

	fp.wg.Add(1)

	// Running the goroutines that will persist valid lines
	for i := 0; i < totalRoutines; i++ {
		fp.wg.Add(1)

		go func() {
			defer fp.wg.Done()

			defer func() {
				if r := recover(); r != nil {
					stack := debug.Stack()
					log.Println("recovered from panic", r)
					log.Println(string(stack))
				}
			}()

			fp.persistGeoData(ctx)
		}()
	}

	// Processing the file
	go func(filename string) {
		defer func() {
			if r := recover(); r != nil {
				stack := debug.Stack()
				log.Println("recovered from panic", r)
				log.Println(string(stack))
			}
		}()

		err := fp.processFile(filename)
		if err != nil {
			panic(err)
		}
	}(dumpFile)

	fp.wg.Wait()

	log.Printf("File importer is done = Total lines: %d, accepted lines: %d, invalid lines: %d, elapsed time: %s\n",
		fp.TotalLines, fp.AcceptedLines, fp.InvalidLines, time.Since(startTime))
}

// processFile opens the file specified in the DUMP_FILE environment var, checks if it's valid against the defined
// csv schema (defined by the header) and sends each line in the CSV for async processing.
//
// The actual contents of each line (after being converted to a models.Geolocation struct) is validated before
// persisting it.
func (fp *fileProcessor) processFile(filename string) error {
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

		fp.TotalLines++

		// Once the header is known we can continue to the proper lines in the CSV
		g, err := csvLineToStruct(header, scanner.Text())
		if err != nil {
			fp.incrementInvalidCount()
			continue
		}

		fp.data <- g
	}

	close(fp.data)

	if err := scanner.Err(); err != nil {
		return err
	}

	fp.wg.Done()

	return nil
}

// persistGeoData validates and persists geolocation data, feeding InvalidLines via an atomic operation
func (fp *fileProcessor) persistGeoData(ctx context.Context) {
	for g := range fp.data {
		// Checking if the data is valid
		if err := g.Validate(); err != nil {
			// Incrementing one on the list of total errors
			fp.incrementInvalidCount()
			continue
		}

		err := fp.repository.AddLocationInfo(ctx, g)
		if err != nil {
			log.Println(err)
			fp.incrementInvalidCount()
			continue
		}

		fp.incrementAcceptedCount()
	}
}

func (fp *fileProcessor) incrementAcceptedCount() {
	atomic.AddUint64(&fp.AcceptedLines, 1)
}

func (fp *fileProcessor) incrementInvalidCount() {
	atomic.AddUint64(&fp.InvalidLines, 1)
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
