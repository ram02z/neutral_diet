// Package service implements service requirements.
package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ram02z/neutral_diet/internal/service/db"
)

// NewDataStore is a wrapper around [db.NewStore].
func NewDataStore(dbPool *pgxpool.Pool) *db.Store {
	return db.NewStore(dbPool)
}
