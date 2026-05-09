package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mahadia/mahadia-spotifyData/goSpotify/models"
	"github.com/mahadia/mahadia-spotifyData/goSpotify/pg"
)

func main() {
	startTime := time.Now()
	var prodFlag bool
	flag.BoolVar(&prodFlag, "prod", false, "prod env")
	flag.BoolVar(&prodFlag, "p", false, "prod env")
	var largeFlag bool
	flag.BoolVar(&largeFlag, "large", false, "Use the large data set")
	flag.BoolVar(&largeFlag, "l", false, "Use the large data set")
	flag.Parse()

	if prodFlag {
		fmt.Println("PROD")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

	} else {
		fmt.Println("TEST")
		err := godotenv.Load(".env.test")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

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

	data, err := processSpotifyData(largeFlag)
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

func processSpotifyData(largeFlag bool) ([]models.SpotifyData, error) {
	var file string
	if largeFlag {
		file = "../rawSpotifyData/endsong_0.json"
	} else {
		file = "../rawSpotifyData/smallSample.json"
	}

	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer jsonFile.Close()

	io.MultiReader()
	byteValue, err := io.ReadAll(jsonFile)
	var spotifyMiniData []models.SpotifyData

	err = json.Unmarshal(byteValue, &spotifyMiniData)
	if err != nil {
		fmt.Println("kaos with unmarshal spotify data", err)
		return nil, err
	}

	return spotifyMiniData, nil
}
