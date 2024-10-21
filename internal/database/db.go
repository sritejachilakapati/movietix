package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context) *pgx.Conn {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	return conn
}
