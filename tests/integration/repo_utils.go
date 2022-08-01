package integration

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tiagocesar/geolocation/internal/models"
)

// This file defines utility functions to interact with the db that will be used only for integration testing

type testRepository struct {
	db *sql.DB
}

func NewRepositoryForIntegrationTesting(user, pass, host, port, schema string) (*testRepository, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, schema)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &testRepository{
		db: db,
	}, nil
}

// GetTestRows will get all testing data for comparison on tests
// to identify testing data we use "ZZZ" as country code and "Integration Testing" as country
func (tr *testRepository) GetTestRows(ctx context.Context) ([]models.Geolocation, error) {
	var result []models.Geolocation

	q := `SELECT ip_address, country_code, country, city, latitude, longitude
            FROM public.location_info
           WHERE country_code = 'ZZZ'
             AND country = 'Integration Testing'`

	rows, err := tr.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		location := models.Geolocation{}
		err := rows.Scan(&location.IpAddress, &location.CountryCode, &location.Country, &location.City,
			&location.Latitude, &location.Longitude)
		if err != nil {
			return nil, err
		}

		result = append(result, location)
	}

	return result, nil
}

// CleanDB will remove all testing data from the database.
// to identify testing data we use "ZZZ" as country code and "Integration Testing" as country
func (tr *testRepository) CleanDB(ctx context.Context) error {
	_, err := tr.db.ExecContext(ctx,
		`DELETE FROM public.location_info
                 WHERE country_code = 'ZZZ'
                   AND country = 'Integration Testing'`)

	return err
}
