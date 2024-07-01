package pg

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
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
	if err := pg.CreatePodcastTable(ctx); err != nil {
		return err
	}
	if err := pg.CreatePlaybackTable(ctx); err != nil {
		return err
	}
	if err := pg.CreateMediaTable(ctx); err != nil {
		return err
	}
	return nil
}

func InsertIntoDb(pg *Postgres, ctx context.Context, data models.SpotifyData) error {
	if err := InsertUsersValues(pg, ctx, data); err != nil {
		fmt.Println("error insert user values")
		return err
	}

	if err := InsertTrackValues(pg, ctx, data); err != nil {
		fmt.Println("error insert track values")
		return err
	}

	if err := InsertPodcastValues(pg, ctx, data); err != nil {
		fmt.Println("error insert podcast values")
		return err
	}

	if err := InsertPlaybackValues(pg, ctx, data); err != nil {
		fmt.Println("error insert playback values")
		return err
	}
	if err := InsertMediaValues(pg, ctx, data); err != nil {
		fmt.Println("error insert media values")
		return err
	}

	return nil
}

func InsertTrackValues(pg *Postgres, ctx context.Context, data models.SpotifyData) error {
	if data.SpotifyTrackUri == nil {
		return nil
	}
	query := `insert into track
	   (
	     track_id,
	     track_name,
	     artist_name,
	     album_name,
	     spotify_track_uri
	   ) values ($1, $2, $3, $4, $5) on conflict (track_id) do nothing`

	trackId := trimUri(data.SpotifyTrackUri)
	trackValues := models.Track{
		TrackID:         trackId,
		TrackName:       *data.MasterMetadataTrackName,
		ArtistName:      *data.MasterMetadataAlbumArtistName,
		AlbumName:       *data.MasterMetadataAlbumAlbumName,
		SpotifyTrackUri: *data.SpotifyTrackUri,
	}

	fmt.Printf("insert into track table: %v\n", trackValues.TrackName)

	_, err := pg.Db.Exec(
		ctx, query,
		trackValues.TrackID,
		trackValues.TrackName,
		trackValues.ArtistName,
		trackValues.AlbumName,
		trackValues.SpotifyTrackUri)

	if err != nil {
		return err
	}

	return nil
}

func InsertPodcastValues(pg *Postgres, ctx context.Context, data models.SpotifyData) error {
	if data.SpotifyEpisodeUri == nil {
		return nil
	}
	query := `insert into podcast
	   (
	     podcast_id,
	     episode_name,
	     episode_show_name,
       spotify_episode_uri 
	   ) values ($1, $2, $3, $4) on conflict (podcast_id) do nothing`

	podcastId := trimUri(data.SpotifyEpisodeUri)
	podcastValues := models.Podcast{
		PodcastId:         podcastId,
		EpisodeName:       *data.EpisodeName,
		EpisodeShowName:   *data.EpisodeShowName,
		SpotifyEpisodeUri: *data.SpotifyEpisodeUri,
	}

	fmt.Printf("insert into playback table: %v\n", podcastValues.EpisodeName)

	_, err := pg.Db.Exec(
		ctx, query,
		podcastValues.PodcastId,
		podcastValues.EpisodeName,
		podcastValues.EpisodeShowName,
		podcastValues.SpotifyEpisodeUri,
	)

	if err != nil {
		return err
	}

	return nil
}

func InsertMediaValues(pg *Postgres, ctx context.Context, data models.SpotifyData) error {
	query := `insert into media (
    playback_id,
    media_id,
    media_type
  ) values ($1, $2, $3)`

	mediaType, mediaId := getMediaType(data)
	playbackId, err := generatePlaybackId(data.UserName, data.Timestamp, mediaId)
	if err != nil {
		return err
	}

	mediaValues := models.Media{
		PlaybackId: playbackId,
		MediaId:    mediaId,
		MediaType:  mediaType,
	}

	fmt.Printf("Inserting  into media: %v\n", playbackId)
	fmt.Printf("media_id: %v", mediaId)
	fmt.Printf("media_type %v", mediaType)
	fmt.Println()
	_, errDb := pg.Db.Exec(
		ctx, query,
		mediaValues.PlaybackId,
		mediaValues.MediaId,
		mediaValues.MediaType,
	)

	if errDb != nil {
		return errDb
	}

	return nil
}

func InsertPlaybackValues(pg *Postgres, ctx context.Context, data models.SpotifyData) error {
	query := `INSERT INTO playback (
    id,
		user_name,
    ts,
    platform,
    ms_played,
    conn_country,
    ip_addr_decrypted,
    user_agent_decrypted,
    reason_start,
    reason_end,
    shuffle,
    skipped,
    offline,
    offline_timestamp,
    incognito_mode,
    media_type
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`

	mediaType, mediaId := getMediaType(data)
	playbackId, err := generatePlaybackId(data.UserName, data.Timestamp, mediaId)
	if err != nil {
		return err
	}

	playbackValues := models.Playback{
		Id:                 playbackId,
		UserName:           data.UserName,
		Timestamp:          data.Timestamp,
		Platform:           data.Platform,
		MsPlayed:           data.MsPlayed,
		ConnCountry:        data.ConnCountry,
		IpAddrDecrypted:    data.IpAddrDecrypted,
		UserAgentDecrypted: data.UserAgentDecrypted,
		ReasonStart:        data.ReasonStart,
		ReasonEnd:          data.ReasonEnd,
		Shuffle:            data.Shuffle,
		Skipped:            data.Skipped,
		Offline:            data.Offline,
		OfflineTimestamp:   data.OfflineTimestamp,
		IncognitoMode:      data.IncognitoMode,
		MediaType:          mediaType,
	}

	fmt.Printf("Inserting playback values for user: %v at time: %v\n", playbackValues.UserName, playbackValues.Timestamp)

	_, errDb := pg.Db.Exec(
		ctx, query,
		playbackValues.Id,
		playbackValues.UserName,
		playbackValues.Timestamp,
		playbackValues.Platform,
		playbackValues.MsPlayed,
		playbackValues.ConnCountry,
		playbackValues.IpAddrDecrypted,
		playbackValues.UserAgentDecrypted,
		playbackValues.ReasonStart,
		playbackValues.ReasonEnd,
		playbackValues.Shuffle,
		playbackValues.Skipped,
		playbackValues.Offline,
		playbackValues.OfflineTimestamp,
		playbackValues.IncognitoMode,
		playbackValues.MediaType,
	)

	// NOTE: Handle error gracefully while also notifying if there are any collisions
	if errDb != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				fmt.Printf("Duplicate playback values for user: %v at time: %v already exists\n", playbackValues.UserName, playbackValues.Timestamp)
				return nil
			}
		}
		return errDb
	}

	return nil
}

func InsertUsersValues(pg *Postgres, ctx context.Context, data models.SpotifyData) error {
	query := `insert into users (user_name) 
    values ($1) on conflict (user_name) do nothing`

	userValues := models.Users{
		UserName: &data.UserName,
	}

	fmt.Printf("Insert into users table: %v\n", *userValues.UserName)

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

func (pg *Postgres) CreatePodcastTable(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS podcast (
    podcast_id VARCHAR PRIMARY KEY,
    episode_name VARCHAR,
    episode_show_name VARCHAR,
    spotify_episode_uri VARCHAR
  )`

	_, err := pg.Db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("unable to create table podcast: %w", err)
	}

	fmt.Println("created table podcast")

	return nil
}

func (pg *Postgres) CreateTracksTable(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS track (
    track_id VARCHAR PRIMARY KEY,
    track_name VARCHAR,
    artist_name VARCHAR,
    album_name VARCHAR,
    spotify_track_uri VARCHAR
  )`

	_, err := pg.Db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("unable to create table track: %w", err)
	}

	fmt.Println("created table track")

	return nil
}

func (pg *Postgres) CreateMediaTable(ctx context.Context) error {
	query := `create table if not exists media (
    playback_id varchar,
    media_id varchar,
    media_type varchar,
    foreign key (playback_id) references playback(id)
  )`

	_, err := pg.Db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("unable to create table media: %w", err)
	}

	fmt.Println("created table media")

	return nil
}

func (pg *Postgres) CreatePlaybackTable(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS playback (
    id VARCHAR PRIMARY KEY,
    user_name VARCHAR,
    ts TIMESTAMP,
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
    media_type VARCHAR,
    FOREIGN KEY (user_name) REFERENCES users(user_name),
    CONSTRAINT UNIQUE_USER_TIMESTAMP_COMBO UNIQUE (user_name, ts)
  )`

	_, err := pg.Db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("unable to create table playback: %w", err)
	}

	fmt.Println("created table playback")

	return nil
}

func (pg *Postgres) DropAllTables(ctx context.Context) error {
	if err := pg.DropMediaTable(ctx); err != nil {
		return err
	}
	if err := pg.DropPlaybackTable(ctx); err != nil {
		return err
	}

	if err := pg.DropUsersTable(ctx); err != nil {
		return err
	}

	if err := pg.DropTrackTable(ctx); err != nil {
		return err
	}

	if err := pg.DropPodcastTable(ctx); err != nil {
		return err
	}
	return nil
}

func (pg *Postgres) DropPlaybackTable(ctx context.Context) error {
	query := `drop table if exists playback`
	_, err := pg.Db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("unable to drop table playback: %w", err)
	}

	fmt.Println("dropped table playback")
	return nil
}

func (pg *Postgres) DropMediaTable(ctx context.Context) error {
	query := `drop table if exists media`
	_, err := pg.Db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("unable to drop table media: %w", err)
	}

	fmt.Println("dropped table media")
	return nil
}

func (pg *Postgres) DropUsersTable(ctx context.Context) error {
	query := `drop table if exists users`
	_, err := pg.Db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("unable to drop table users: %w", err)
	}

	fmt.Println("dropped table users")
	return nil
}

func (pg *Postgres) DropTrackTable(ctx context.Context) error {
	query := `drop table if exists track`
	_, err := pg.Db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("unable to drop table track: %w", err)
	}

	fmt.Println("dropped table track")
	return nil
}

func (pg *Postgres) DropPodcastTable(ctx context.Context) error {
	query := `drop table if exists podcast`
	_, err := pg.Db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("unable to drop table podcast: %w", err)
	}

	fmt.Println("dropped table podcast")
	return nil
}

func (pg *Postgres) Close() {
	pg.Db.Close()
}

func generatePlaybackId(userName string, dateTime time.Time, mediaId string) (string, error) {
	ts := getUnixTs(dateTime)
	return userName + ":" + ts + ":" + mediaId, nil
}

func getUnixTs(dateTime time.Time) string {
	unixTime := dateTime.Unix()
	strTime := strconv.FormatInt(unixTime, 10)
	return strTime
}

func trimUri(uri *string) string {
	if uri != nil {
		return strings.TrimPrefix(*uri, "spotify:")
	}
	return ""
}

func getMediaType(data models.SpotifyData) (mediaType, mediaId string) {
	if data.SpotifyTrackUri != nil {
		mediaType = "track"
		mediaId = trimUri(data.SpotifyTrackUri)
	} else {
		mediaType = "podcast"
		mediaId = trimUri(data.SpotifyEpisodeUri)
	}
	return mediaType, mediaId
}
