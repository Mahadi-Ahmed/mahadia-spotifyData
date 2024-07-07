package pg

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/mahadia/mahadia-spotifyData/goSpotify/pg/testHelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresTestSuite struct {
	suite.Suite
	ctx       context.Context
	container *postgres.PostgresContainer
	pg        *Postgres
}

func (suite *PostgresTestSuite) TestDropAllTable() {
	t := suite.T()
	err := suite.pg.DropAllTables(suite.ctx)
	assert.NoError(t, err, "Failed to create all tables")
}

func (suite *PostgresTestSuite) TestCreateAllTable() {
	t := suite.T()
	err := suite.pg.CreateAllTables(suite.ctx)
	assert.NoError(t, err, "Failed to create all tables")
}

func (suite *PostgresTestSuite) TestInsertTrackIntoDbTable() {
	t := suite.T()
	suite.pg.CreateAllTables(suite.ctx)
	err := suite.pg.InsertIntoDb(suite.ctx, pg_testhelper.TestDataValidTrack)
	assert.NoError(t, err, "Failed funciton InsertIntoDb")

	var countPlayback int
	var countTrack int
	var countPodcast int
	var countUser int
	var countMedia int

	errCountPlayback := suite.pg.Db.QueryRow(suite.ctx, "select count(*) from playback").Scan(&countPlayback)
	errCountTrack := suite.pg.Db.QueryRow(suite.ctx, "select count(*) from track").Scan(&countTrack)
	errCountPodcast := suite.pg.Db.QueryRow(suite.ctx, "select count(*) from podcast").Scan(&countPodcast)
	errCountUser := suite.pg.Db.QueryRow(suite.ctx, "select count(*) from user").Scan(&countUser)
	errCountMedia := suite.pg.Db.QueryRow(suite.ctx, "select count(*) from media").Scan(&countMedia)
	assert.NoError(t, errCountPlayback, "Failed to query playback")
	assert.NoError(t, errCountTrack, "Failed to query track")
	assert.NoError(t, errCountPodcast, "Failed to query podcast")
	assert.NoError(t, errCountUser, "Failed to query user")
	assert.NoError(t, errCountMedia, "Failed to query media")

	assert.Equal(t, 1, countPlayback, "playback was not inserted correctly")
	assert.Equal(t, 1, countTrack, "track was not inserted correctly")
	assert.Equal(t, 0, countPodcast, "podcast was not inserted correctly")
	assert.Equal(t, 1, countUser, "user was not inserted correctly")
	assert.Equal(t, 1, countMedia, "media was not inserted correctly")
}

func (suite *PostgresTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := postgres.RunContainer(suite.ctx,
		testcontainers.WithImage("postgres:16.3-alpine"),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatal(err)
	}

	suite.container = pgContainer

	connStr, err := pgContainer.ConnectionString(suite.ctx, "sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	pgDb, err := NewPG(suite.ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}
	// suite.pg = *pgDb
	suite.pg = pgDb
}

func (suite *PostgresTestSuite) TearDownSuite() {
	if err := suite.container.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func TestPostgresSpotifyTestSuite(t *testing.T) {
	suite.Run(t, new(PostgresTestSuite))
}
