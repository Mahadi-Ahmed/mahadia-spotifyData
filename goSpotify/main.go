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
  conn, err := pgx.Connect(context.Background(), connStr)
  if err != nil {
    fmt.Println("unable to connect to database", err)
    os.Exit(1)
  }
  defer conn.Close(context.Background())

  var employees string
  err = conn.QueryRow(context.Background(), "select name from employees where salary=65000").Scan(&employees)
  if err != nil {
    fmt.Println("Query failed")
    fmt.Println(err)
  }
  fmt.Println(employees)
}
