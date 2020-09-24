//go:generate sh ../../scripts/generate_migrations.sh

package database

import (
	"fmt"
	"github.com/dshemin/gopencov/internal/database/internal/postgresql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgresql driver
)

// New create database instance
// Also will run all necessary migrations
func New(driverName, connURI string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(driverName, connURI)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to databse with %q driver and URI %q: %w", driverName, connURI, err)
	}

	if err := migrateUp(db, driverName); err != nil {
		return nil, fmt.Errorf("cannot migrate DB schema: %w", err)
	}

	return db, nil
}

func migrateUp(db *sqlx.DB, driverName string) error {
	resource := bindata.Resource(postgresql.AssetNames(), func(name string) ([]byte, error) {
		return postgresql.Asset(name)
	})

	drv, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}
	bn, err := bindata.WithInstance(resource)
	if err != nil {
		return err
	}
	migrator, err := migrate.NewWithInstance("go-bindata", bn, driverName, drv)
	if err != nil {
		return err
	}

	err = migrator.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
