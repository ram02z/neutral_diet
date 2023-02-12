package db

import "github.com/jackc/pgx/v5/pgxpool"

type Store struct {
	dbPool *pgxpool.Pool
}

func NewStore(dbPool *pgxpool.Pool) *Store {
	return &Store{
		dbPool: dbPool,
	}
}
