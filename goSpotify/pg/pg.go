package pg

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mahadia/mahadia-spotifyData/goSpotify/models"
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

func InsertIntoDb(pg *Postgres, ctx context.Context, data models.SpotifyData) error {
	if err := InsertUsersValues(pg, ctx, data); err != nil {
		return err
	}

	if err := InsertTrackValues(pg, ctx, data); err != nil {
		return err
	}

	return nil
}

func InsertPlaybackValues(pg *Postgres, ctx context.Context, data models.SpotifyData) error {

	return nil
}

func InsertTrackValues(pg *Postgres, ctx context.Context, data models.SpotifyData) error {
  fmt.Println()
  fmt.Println("InsertTrackValues")
  fmt.Println()
	query := `insert into track
	   (
	     track_id,
	     track_name,
	     artist_name,
	     album_name,
	     spotify_track_uri,
	     episode_name,
	     episode_show_name,
	     spotify_episode_uri
	   ) values ($1, $2, $3, $4, $5 ,$6, $7, $8)`

	trackId := strings.TrimPrefix(data.SpotifyTrackUri, "spotify:track:")
	trackValues := models.Track{
		TrackID:           &trackId,
		TrackName:         &data.MasterMetadataTrackName,
		ArtistName:        &data.MasterMetadataAlbumArtistName,
		AlbumName:         &data.MasterMetadataAlbumAlbumName,
		SpotifyTrackUri:   &data.SpotifyTrackUri,
		EpisodeName:       data.EpisodeName,
		EpisodeShowName:   data.EpisodeShowName,
		SpotifyEpisodeUri: data.SpotifyEpisodeUri,
	}

	fmt.Println("Inserting these values to Track: ")
	fmt.Println(*trackValues.TrackName + "\n")

	_, err := pg.Db.Exec(
		ctx, query,
		trackValues.TrackID,
		trackValues.TrackName,
		trackValues.ArtistName,
		trackValues.AlbumName,
		trackValues.SpotifyTrackUri,
		trackValues.EpisodeName,
		trackValues.EpisodeShowName,
		trackValues.SpotifyEpisodeUri)

	if err != nil {
		return err
	}

	return nil
}

func InsertUsersValues(pg *Postgres, ctx context.Context, data models.SpotifyData) error {
  fmt.Println()
  fmt.Println("InsertUsersValues")
  fmt.Println()
	query := `insert into users (user_name) 
    values ($1)`

	userValues := models.Users{
		UserName: &data.UserName,
	}

	fmt.Println("gonna insert user: ", *userValues.UserName)

	_, err := pg.Db.Exec(ctx, query, userValues.UserName)
	if err != nil {
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
    artist_name VARCHAR,
    album_name VARCHAR,
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

func (pg *Postgres) DropAllTables(ctx context.Context) error {
	if err := pg.DropPlaybackTable(ctx); err != nil {
		return err
	}

	if err := pg.DropUsersTable(ctx); err != nil {
		return err
	}

	if err := pg.DropTrackTable(ctx); err != nil {
		return err
	}
	return nil
}

func (pg *Postgres) DropPlaybackTable(ctx context.Context) error {
	query := `drop table if exits playback`
	_, err := pg.Db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("unable to create table: %w", err)
	}

	fmt.Println("dropped table playback")
	return nil
}

func (pg *Postgres) DropUsersTable(ctx context.Context) error {
	query := `drop table if exits users`
	_, err := pg.Db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("unable to create table: %w", err)
	}

	fmt.Println("dropped table users")
	return nil
}

func (pg *Postgres) DropTrackTable(ctx context.Context) error {
	query := `drop table if exits track`
	_, err := pg.Db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("unable to create table: %w", err)
	}

	fmt.Println("dropped table track")
	return nil
}

func (pg *Postgres) Close() {
	pg.Db.Close()
}
