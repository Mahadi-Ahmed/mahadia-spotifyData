package pg

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Db *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

func NewPG(ctx context.Context, connStr string) (*Postgres, error) {
	var initErr error
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connStr)
		if err != nil {
			initErr = fmt.Errorf("unable to create connection pool: %w", err)
			return
		}
		pgInstance = &Postgres{Db: db}
	})

	if initErr != nil {
		return nil, initErr
	}

	return pgInstance, nil
}

func (pg *Postgres) Close() {
	pg.Db.Close()
}

func (pg *Postgres) CreateAllTables(ctx context.Context) error {
	if err := pg.CreateUsersTable(ctx); err != nil {
		return err
	}
	if err := pg.CreateTracksTable(ctx); err != nil {
		return err
	}
	if err := pg.CreatePlaybackTable(ctx); err != nil {
		return err
	}
	return nil
}

func (pg *Postgres) CreateUsersTable(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS users (
    user_name VARCHAR PRIMARY KEY
  )`

	_, err := pg.Db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("unable to create table: %w", err)
	}

	fmt.Println("created table users")

	return nil
}

func (pg *Postgres) CreateTracksTable(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS track (
    track_id VARCHAR PRIMARY KEY,
    track_name VARCHAR,
    album_artist_name VARCHAR,
    album_album_name VARCHAR,
    spotify_track_uri VARCHAR,
    episode_name VARCHAR,
    episode_show_name VARCHAR,
    spotify_episode_uri VARCHAR
  )`

	_, err := pg.Db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("unable to create table: %w", err)
	}

	fmt.Println("created table track")

	return nil
}

func (pg *Postgres) CreatePlaybackTable(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS playback (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR,
    ts TIMESTAMP,
    track_id VARCHAR,
    platform VARCHAR,
    ms_played BIGINT,
    conn_country VARCHAR,
    ip_addr_decrypted VARCHAR,
    user_agent_decrypted VARCHAR,
    reason_start VARCHAR,
    reason_end VARCHAR,
    shuffle BOOLEAN,
    skipped BOOLEAN,
    offline BOOLEAN,
    offline_timestamp BIGINT,
    incognito_mode BOOLEAN,
    FOREIGN KEY (user_name) REFERENCES users(user_name),
    FOREIGN KEY (track_id) REFERENCES track(track_id)
  )`

	_, err := pg.Db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("unable to create table: %w", err)
	}

	fmt.Println("created table playback")

	return nil
}
