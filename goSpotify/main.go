package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mahadia/mahadia-spotifyData/goSpotify/pg"
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

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbName)
	pgConn, err := pg.NewPG(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pgConn.Close()

	var employees string
	err = pgConn.Db.QueryRow(context.Background(), "SELECT name FROM employees WHERE salary=65000").Scan(&employees)
	if err != nil {
		fmt.Println("Query failed")
		fmt.Println(err)
	} else {
		fmt.Println(employees)
	}
}
