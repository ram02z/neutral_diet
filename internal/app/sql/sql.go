package sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	DBConfig *DBConfig `env:",prefix=DB_"`
}

type DBConfig struct {
	User string `env:"USER,default=postgres"`
	Pass string `env:"PASS,default=postgres"`
	Host string `env:"HOST,default=localhost"`
	Name string `env:"NAME,default=postgres"`
	Port int    `env:"PORT,default=5432"`
}

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
	conn, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return conn, err
}
