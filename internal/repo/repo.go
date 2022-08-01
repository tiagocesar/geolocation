package repo

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/tiagocesar/geolocation/internal/models"
)

const tableLocationInfo = "public.location_info"

type repository struct {
	db *sql.DB
}

func NewRepository(user, pass, host, port, schema string) (*repository, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, schema)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &repository{
		db: db,
	}, nil
}

func (r *repository) AddLocationInfo(ctx context.Context, locationInfo models.Geolocation) error {
	q := `INSERT INTO ` + tableLocationInfo + `(ip_address, country_code, country, city, latitude, longitude, mystery_value)
          VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.db.ExecContext(ctx, q, locationInfo.IpAddress, locationInfo.CountryCode, locationInfo.Country,
		locationInfo.City, locationInfo.Latitude, locationInfo.Longitude, locationInfo.MysteryValue)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetLocationInfoByIP(ctx context.Context, ipAddress string) (*models.Geolocation, error) {
	q := `SELECT ip_address, country_code, country, city, latitude, longitude
		    FROM ` + tableLocationInfo + `
           WHERE ip_address = $1`

	var response models.Geolocation
	// err can be sql.ErrNoRows
	err := r.db.QueryRowContext(ctx, q, ipAddress).Scan(&response.IpAddress, &response.CountryCode, &response.Country,
		&response.City, &response.Latitude, &response.Longitude)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
