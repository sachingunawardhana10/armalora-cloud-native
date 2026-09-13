package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sachingunawardhana10/armalora-cloud-native/internal/config"
)

func Connect(cfg config.Config) (*pgx.Conn, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
