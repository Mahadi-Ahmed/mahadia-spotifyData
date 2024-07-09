package pg_testhelper

import (
	"time"

	"github.com/mahadia/mahadia-spotifyData/goSpotify/models"
)

var (
	TestDataValidTrack = models.SpotifyData{
		Timestamp:                     parseTime("2018-07-14T22:36:12Z"),
		UserName:                      "mahadi4",
		Platform:                      "Android OS 8.1.0 API 27 (OnePlus, ONEPLUS A6003)",
		MsPlayed:                      12491,
		ConnCountry:                   "SE",
		IpAddrDecrypted:               "85.228.229.56",
		UserAgentDecrypted:            "unknown",
		ReasonStart:                   "clickrow",
		ReasonEnd:                     "endplay",
		Shuffle:                       false,
		Skipped:                       false,
		Offline:                       false,
		OfflineTimestamp:              1531607758299,
		IncognitoMode:                 false,
		SpotifyTrackUri:               stringPtr("spotify:track:3clX2NMmjaAHmBjeTSa9vV"),
		MasterMetadataTrackName:       stringPtr("Actin Crazy"),
		MasterMetadataAlbumArtistName: stringPtr("Action Bronson"),
		MasterMetadataAlbumAlbumName:  stringPtr("Mr. Wonderful"),
		EpisodeName:                   nil,
		EpisodeShowName:               nil,
		SpotifyEpisodeUri:             nil,
	}

	TestDataValidPodcast = models.SpotifyData{
		Timestamp:                     parseTime("2020-09-03T09:25:47Z"),
		UserName:                      "mahadi4",
		Platform:                      "Android OS 10 API 29 (OnePlus, ONEPLUS A6003)",
		MsPlayed:                      732683,
		ConnCountry:                   "SE",
		IpAddrDecrypted:               "84.216.129.54",
		UserAgentDecrypted:            "unknown",
		ReasonStart:                   "clickrow",
		ReasonEnd:                     "logout",
		Shuffle:                       false,
		Skipped:                       false,
		Offline:                       false,
		OfflineTimestamp:              1599072060296,
		IncognitoMode:                 false,
		SpotifyEpisodeUri:             stringPtr("spotify:episode:0clCLLUxZTRE2wNlEEM9xL"),
		EpisodeName:                   stringPtr("Don't Be a YouTuber in Japan (ft. Abroad in Japan) | Trash Taste #5"),
		EpisodeShowName:               stringPtr("Trash Taste Podcast"),
		MasterMetadataTrackName:       nil,
		MasterMetadataAlbumArtistName: nil,
		MasterMetadataAlbumAlbumName:  nil,
		SpotifyTrackUri:               nil,
	}
)

func stringPtr(s string) *string {
	return &s
}

func parseTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}
