package db

import (
	"performance_tracker_v2_be/config"
	mainDB "performance_tracker_v2_be/db/main-db"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Make interfaces for databases in the future, definitely don't need to do it now

type Registry struct {
	MainDatabase *pgxpool.Pool
}

func InitializeDatabases(cfg *config.Config) (*Registry, error) {
	pgPool, err := mainDB.Connect(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	println("Initialized connections to all databases")

	return &Registry{
		MainDatabase: pgPool,
	}, nil
}

func (r *Registry) Close() {
	if r.MainDatabase != nil {
		r.MainDatabase.Close()
	}
}
