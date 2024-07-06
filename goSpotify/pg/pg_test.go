package pg

import (
	"context"
	"log"
	"testing"
	"time"

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
  assert.NoError(t, err , "Failed to create all tables")
}

func (suite *PostgresTestSuite) TestCreateAllTable() {
  t := suite.T()
  err := suite.pg.CreateAllTables(suite.ctx)
  assert.NoError(t, err , "Failed to create all tables")
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
