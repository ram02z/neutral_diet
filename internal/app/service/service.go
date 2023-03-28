package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ram02z/neutral_diet/internal/service/db"
)

func NewDataStore(dbPool *pgxpool.Pool) *db.Store {
	return db.NewStore(dbPool)
}
