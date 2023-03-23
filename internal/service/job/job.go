package job

import (
	"context"

	"firebase.google.com/go/messaging"
	"github.com/jackc/pgx/v5/pgxpool"
)

type JobWrapper struct {
	p *pgxpool.Pool
	m *messaging.Client
	c *context.Context
}

func NewJobWrapper(p *pgxpool.Pool, m *messaging.Client, c *context.Context) *JobWrapper {
	return &JobWrapper{
		p: p,
		m: m,
		c: c,
	}
}

func (j *JobWrapper) Jobs() []func() {
	return []func(){j.userGoalNotificationsJob}
}
