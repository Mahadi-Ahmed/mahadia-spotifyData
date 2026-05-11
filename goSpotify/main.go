package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mahadia/mahadia-spotifyData/goSpotify/models"
	"github.com/mahadia/mahadia-spotifyData/goSpotify/pg"
)

func main() {
	startTime := time.Now()
	logger, logFile := newLogger()
	defer logFile.Close()

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
			logger.Error("Error loading .env file")
			os.Exit(1)
		}

	} else {
		logger.Info("Loading env files", "env", "test")
		err := godotenv.Load(".env.test")
		if err != nil {
			logger.Error("Error loading .env file")
			os.Exit(1)
		}
	}

	var (
		dbHost     = os.Getenv("DB_HOST")
		dbUser     = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PW")
		dbName     = os.Getenv("DB_NAME")
	)

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbName)
	pgConn, err := pg.NewPG(context.Background(), connStr, logger)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	defer pgConn.Close()

	// NOTE: Drop tables
	err = pgConn.DropAllTables(context.Background())
	if err != nil {
		logger.Error("Failed to drop tables", "error", err)
		os.Exit(1)
	}

	// NOTE: Create tables
	err = pgConn.CreateAllTables(context.Background())
	if err != nil {
		logger.Error("Failed to create tables", "error", err)
		os.Exit(1)
	}

	data, err := processSpotifyData(largeFlag, logger)
	if err != nil {
		logger.Error("failed to process spotify data", "error", err)
		os.Exit(1)
	}

	for _, v := range data {
		if v.Offline {
			ms := int64(v.OfflineTimestamp)
			v.Timestamp = time.UnixMilli(ms).UTC()
		}
		err := pgConn.InsertIntoDb(context.Background(), v)
		fmt.Println()
		if err != nil {
			logger.Error("failed insert into db", "error", err)
		}
	}
	elapsedTime := time.Since(startTime).Seconds()
	logger.Info("Seeding DB finished", "duration", elapsedTime, "items", len(data))
}

func processSpotifyData(largeFlag bool, logger *slog.Logger) ([]models.SpotifyData, error) {
	var file string
	if largeFlag {
		file = "../rawSpotifyData/endsong_0.json"
	} else {
		file = "../rawSpotifyData/smallSample.json"
	}

	jsonFile, err := os.Open(file)
	if err != nil {
		logger.Error("error parsing data", "error", err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		logger.Error("error reading data", "error", err)
		return nil, err
	}
	var spotifyMiniData []models.SpotifyData

	err = json.Unmarshal(byteValue, &spotifyMiniData)
	if err != nil {
		logger.Error("failed to unmarshal spotify data", "error", err)
		return nil, err
	}

	return spotifyMiniData, nil
}

func newLogger() (*slog.Logger, *os.File) {
	logDir := "logs"
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		log.Fatalf("failed to create log directory")
	}
	logFilePath := fmt.Sprintf("%s/logs.txt", logDir)
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Failed to open file %v: ", err)
	}

	writer := io.MultiWriter(os.Stdout, logFile)
	handler := slog.NewJSONHandler(writer, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	})

	return slog.New(handler), logFile
}
