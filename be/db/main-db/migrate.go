package main_db

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

func RunMigrations(pool *pgxpool.Pool) error {
	// 1. Get a standard *sql.DB from the pgx pool
	db := stdlib.OpenDBFromPool(pool)

	// 2. Create a driver instance for golang-migrate
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create migrate driver: %w", err)
	}

	// 3. Initialize the migrate instance
	// Point this to your filesystem path where .sql files are stored
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/main-db/migrations", // Source URL (file system)
		"postgres",                     // Database name
		driver,
	)
	if err != nil {
		return fmt.Errorf("could not create migrate instance: %w", err)
	}

	// 4. Run "Up" to apply all pending migrations
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("Database is up to date (no new migrations)")
			return nil
		}
		return fmt.Errorf("could not run up migrations: %w", err)
	}

	log.Println("Database migrations applied successfully")
	return nil
}
