package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

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
	// err = pgConn.DropAllTables(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// NOTE: Create tables
	// err = pgConn.CreateAllTables(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// }

	data, err := processSpotifyData()
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range data {
		err := pg.InsertIntoDb(pgConn, context.Background(), v)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func processSpotifyData() ([]models.SpotifyData, error) {
	jsonFile, err := os.Open("../rawSpotifyData/smallSample.json")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	var spotifyMiniData []models.SpotifyData
	err = json.Unmarshal(byteValue, &spotifyMiniData)
	if err != nil {
		fmt.Println("kaos with unmarshal spotify data", err)
		return nil, err
	}

	return spotifyMiniData, nil
}
