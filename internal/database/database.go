package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func New() (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return db, nil
}
