package service

import (
	"context"

	"firebase.google.com/go/messaging"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ram02z/neutral_diet/internal/service/db"
	"github.com/ram02z/neutral_diet/internal/service/job"
)

func NewDataStore(dbPool *pgxpool.Pool) *db.Store {
	return db.NewStore(dbPool)
}

func NewJobWrapper(p *pgxpool.Pool, m *messaging.Client, c *context.Context) *job.JobWrapper {
	return job.NewJobWrapper(p, m, c)
}
