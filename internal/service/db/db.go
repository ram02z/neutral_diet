// Package db implements database interactions for neutral-diet service RPCs.
package db

import "github.com/jackc/pgx/v5/pgxpool"

// Store represents a connection pool to the database.
type Store struct {
	dbPool *pgxpool.Pool
}

// NewStore create a [Store] instance
func NewStore(dbPool *pgxpool.Pool) *Store {
	return &Store{
		dbPool: dbPool,
	}
}
