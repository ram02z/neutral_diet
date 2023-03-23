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

type CronJob struct {
	Job   func()
	At    string
}

func NewJobWrapper(p *pgxpool.Pool, m *messaging.Client, c *context.Context) *JobWrapper {
	return &JobWrapper{
		p: p,
		m: m,
		c: c,
	}
}

func (j *JobWrapper) Jobs() []*CronJob {
	return []*CronJob{
		{
			Job:   j.userGoalNotificationsJob,
			At:    "19:00",
		},
		{
			Job:   j.markCompletedGoalsJob,
			At:    "00:00",
		},
	}
}
