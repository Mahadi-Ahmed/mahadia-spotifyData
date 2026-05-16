package pg

import (
	"context"
	"fmt"
	"log"
	"log/slog"
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
	suite.SetupTest()
	t := suite.T()
	err := suite.pg.InsertIntoDb(suite.ctx, pg_testhelper.TestDataValidTrack)
	assert.NoError(t, err, "Failed funciton InsertIntoDb")

	var countPlayback int
	var countTrack int
	var countPodcast int
	var countUser int

	errCountPlayback := suite.pg.db.QueryRow(suite.ctx, "select count(*) from playback").Scan(&countPlayback)
	errCountTrack := suite.pg.db.QueryRow(suite.ctx, "select count(*) from track").Scan(&countTrack)
	errCountPodcast := suite.pg.db.QueryRow(suite.ctx, "select count(*) from podcast").Scan(&countPodcast)
	errCountUser := suite.pg.db.QueryRow(suite.ctx, "select count(*) from users").Scan(&countUser)
	assert.NoError(t, errCountPlayback, "Failed to query playback")
	assert.NoError(t, errCountTrack, "Failed to query track")
	assert.NoError(t, errCountPodcast, "Failed to query podcast")
	assert.NoError(t, errCountUser, "Failed to query user")

	assert.Equal(t, 1, countPlayback, "playback was not inserted correctly")
	assert.Equal(t, 1, countTrack, "track was not inserted correctly")
	assert.Equal(t, 0, countPodcast, "podcast was not inserted correctly")
	assert.Equal(t, 1, countUser, "user was not inserted correctly")
}

func (suite *PostgresTestSuite) TestInsertAllCollision() {
	suite.SetupTest()
	fmt.Println()
	t := suite.T()
	firstInsert := suite.pg.InsertIntoDb(suite.ctx, pg_testhelper.TestDataValidTrack)
	secondInsert := suite.pg.InsertIntoDb(suite.ctx, pg_testhelper.TestDataValidTrack)
	thirdInsert := suite.pg.InsertIntoDb(suite.ctx, pg_testhelper.TestDataValidTrack2)

	assert.NoError(t, firstInsert, "Should succeed")
	assert.Error(t, secondInsert, "Swallow Duplicate insert")
	assert.NoError(t, thirdInsert, "Should succeed, new entry")
}

func (suite *PostgresTestSuite) TestInsertPodcastIntoDbTable() {
	suite.SetupTest()
	t := suite.T()
	err := suite.pg.InsertIntoDb(suite.ctx, pg_testhelper.TestDataValidPodcast)
	assert.NoError(t, err, "Failed funciton InsertIntoDb")

	var countPlayback int
	var countTrack int
	var countPodcast int
	var countUser int

	errCountPlayback := suite.pg.db.QueryRow(suite.ctx, "select count(*) from playback").Scan(&countPlayback)
	errCountTrack := suite.pg.db.QueryRow(suite.ctx, "select count(*) from track").Scan(&countTrack)
	errCountPodcast := suite.pg.db.QueryRow(suite.ctx, "select count(*) from podcast").Scan(&countPodcast)
	errCountUser := suite.pg.db.QueryRow(suite.ctx, "select count(*) from users").Scan(&countUser)
	assert.NoError(t, errCountPlayback, "Failed to query playback")
	assert.NoError(t, errCountTrack, "Failed to query track")
	assert.NoError(t, errCountPodcast, "Failed to query podcast")
	assert.NoError(t, errCountUser, "Failed to query user")

	assert.Equal(t, 1, countPlayback, "playback was not inserted correctly")
	assert.Equal(t, 0, countTrack, "track was not inserted correctly")
	assert.Equal(t, 1, countPodcast, "podcast was not inserted correctly")
	assert.Equal(t, 1, countUser, "user was not inserted correctly")
}

/*
	 NOTE: Some objects have exact same timestamp due to offline sync issues. Need to check
		The offline timestamp for the actual timestamp
*/
func (suite *PostgresTestSuite) TestInsertOffline() {
	suite.SetupTest()
	t := suite.T()
	firstInsert := suite.pg.InsertIntoDb(suite.ctx, pg_testhelper.TestDataValidTrackOffline1)
	assert.NoError(t, firstInsert, "Should succeed")
	var d time.Time
	suite.pg.db.QueryRow(suite.ctx, "select ts from playback").Scan(&d)
	assert.Equal(t, d, time.Time(time.Date(2017, time.February, 16, 8, 59, 52, 56000000, time.UTC)))
}

/*
	 NOTE: Some objects have exact same timestamp due to offline sync issues. Need to check
		The offline timestamp for the actual timestamp
		Check that we can handle multiple objects with same timestamp but different offline timestamp
*/
func (suite *PostgresTestSuite) TestInsertOfflineCollision() {
	suite.SetupTest()
	t := suite.T()
	firstInsert := suite.pg.InsertIntoDb(suite.ctx, pg_testhelper.TestDataValidTrackOffline1)
	secondInsert := suite.pg.InsertIntoDb(suite.ctx, pg_testhelper.TestDataValidTrackOffline2)
	assert.NoError(t, firstInsert, "Should succeed")
	assert.NoError(t, secondInsert, "Should succeed")
}

// NOTE: Test for insertion of "unknown" tracks, where  track uri & episode uri is missing
func (suite *PostgresTestSuite) TestInsertUnknown() {
	suite.SetupTest()
	t := suite.T()
	unknownInsert := suite.pg.InsertIntoDb(suite.ctx, pg_testhelper.TestDataUnknown)
	assert.NoError(t, unknownInsert, "Failed to insert unknown data")
	var countPlayback, countTrack int
	suite.pg.db.QueryRow(suite.ctx, "select count(*) from playback").Scan(&countPlayback)
	suite.pg.db.QueryRow(suite.ctx, "select count(*) from track").Scan(&countTrack)

	assert.Equal(t, 1, countPlayback, "Playback has NOT 1 row")
	assert.Equal(t, 0, countTrack, "Track has NOT 0 row")

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

	pgDb, err := NewPG(suite.ctx, connStr, slog.Default())
	if err != nil {
		log.Fatal(err)
	}
	suite.pg = pgDb
}

func (suite *PostgresTestSuite) TearDownSuite() {
	if err := suite.container.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *PostgresTestSuite) SetupTest() {
	err := suite.pg.DropAllTables(suite.ctx)
	// NOTE: Use require to fail fast and exit test suite
	suite.Require().NoError(err, "Failed to drop all tables")

	err = suite.pg.CreateAllTables(suite.ctx)
	suite.Require().NoError(err, "Failed to create all tables")
}

func TestPostgresSpotifyTestSuite(t *testing.T) {
	suite.Run(t, new(PostgresTestSuite))
}
