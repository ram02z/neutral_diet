// Package sql provides functions for working with PostgreSQL databases.
package sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sethvargo/go-envconfig"
)

// Config contains database configuration information.
type Config struct {
	DBConfig *DBConfig `env:",prefix=DB_"`
}

// DBConfig contains database connection settings.
type DBConfig struct {
	User string `env:"USER,default=postgres"`
	Pass string `env:"PASS,default=postgres"`
	Host string `env:"HOST,default=localhost"`
	Name string `env:"NAME,default=postgres"`
	Port int    `env:"PORT,default=5432"`
}

// NewDatabase creates a new connection pool to a PostgreSQL database using the using the
// configuration specified in the environment variables.
func NewDatabase() (*pgxpool.Pool, error) {
	var cfg Config
	err := envconfig.Process(context.Background(), &cfg)
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DBConfig.User,
		cfg.DBConfig.Pass,
		cfg.DBConfig.Host,
		cfg.DBConfig.Port,
		cfg.DBConfig.Name,
	)
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)

	return pool, err
}
