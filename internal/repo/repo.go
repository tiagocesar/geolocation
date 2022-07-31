package repo

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/tiagocesar/geolocation/internal/models"
)

type repository struct {
	db *sql.DB
}

func NewRepository(user, pass, host, port, schema string) (*repository, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, schema)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer func(db *sql.DB) { _ = db.Close() }(db)

	return &repository{
		db: db,
	}, nil
}

func (r *repository) AddLocationInfo(ctx context.Context, locationInfo models.Geolocation) error {
	return nil
}

func (r *repository) GetLocationInfoByIP(ctx context.Context, ipAddress string) (*models.Geolocation, error) {
	return nil, nil
}
