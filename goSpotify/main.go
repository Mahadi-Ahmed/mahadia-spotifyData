package main

import (
	// "fmt"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type postgres struct {
	db *pgx.Conn
}

func NewPG(ctx context.Context, connStr string) (*pgx.Conn, error) {
	// NOTE: pgx.Connect is not concurrency safe, if needed look into pgxpool@github.com/jackc/pgx/v5/pgxpool
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Println("unable to connect to database", err)
		os.Exit(1)
	}
  return conn, nil
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
  db, err := NewPG(context.Background(), connStr)
	defer db.Close(context.Background())

	var employees string
	err = db.QueryRow(context.Background(), "select name from employees where salary=65000").Scan(&employees)
	if err != nil {
		fmt.Println("Query failed")
		fmt.Println(err)
	}
	fmt.Println(employees)
}
