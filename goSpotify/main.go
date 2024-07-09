package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
  "time"

	"github.com/joho/godotenv"
	"github.com/mahadia/mahadia-spotifyData/goSpotify/models"
	"github.com/mahadia/mahadia-spotifyData/goSpotify/pg"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
  startTime := time.Now()
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

	// NOTE: Drop tables
	err = pgConn.DropAllTables(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	// NOTE: Create tables
	err = pgConn.CreateAllTables(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	data, err := processSpotifyData()
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range data {
		err := pgConn.InsertIntoDb(context.Background(), v)
    fmt.Println()
		if err != nil {
			fmt.Println(err)
		}
	}
  elapsedTime := time.Since(startTime)
  fmt.Printf("Total time to seed DB: %s\n", elapsedTime)
}

func processSpotifyData() ([]models.SpotifyData, error) {
	jsonFile, err := os.Open("../rawSpotifyData/smallSample.json")
	// jsonFile, err := os.Open("../rawSpotifyData/MyDataGo/endsong_1.json")
	// jsonFile, err := os.Open("../rawSpotifyData/endsong_0.json")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	var spotifyMiniData []models.SpotifyData

	err = json.Unmarshal(byteValue, &spotifyMiniData)
  // fmt.Println(spotifyMiniData)
	if err != nil {
		fmt.Println("kaos with unmarshal spotify data", err)
		return nil, err
	}

	return spotifyMiniData, nil
}
