package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type postgres struct {
	db *pgxpool.Pool
}

var (
  pgInstance *postgres
  pgOnce sync.Once
)

func NewPG(ctx context.Context, connStr string) (*postgres, error) {
	var initErr error
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connStr)
		if err != nil {
			initErr = fmt.Errorf("unable to create connection pool: %w", err)
			return
		}
		pgInstance = &postgres{db}
	})

	if initErr != nil {
		return nil, initErr
	}

	return pgInstance, nil
}

func (pg *postgres) Close() {
	pg.db.Close()
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	var (
		dbHost     = os.Getenv("DB_HOST")
		dbUser     = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PW")
		dbName     = os.Getenv("DB_NAME")
	)

	connStr := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName + "?sslmode=disable"
  pool, err := NewPG(context.Background(), connStr)
  defer pool.Close()

	var employees string
	err = pool.db.QueryRow(context.Background(), "select name from employees where salary=65000").Scan(&employees)
	if err != nil {
		fmt.Println("Query failed")
		fmt.Println(err)
	}
	fmt.Println(employees)
}
