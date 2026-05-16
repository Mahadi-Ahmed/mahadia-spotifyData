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
		Skipped:                       boolPtr(false),
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

	TestDataValidTrack2 = models.SpotifyData{
		Timestamp:                     parseTime("2019-09-20T14:04:37Z"),
		UserName:                      "mahadi4",
		Platform:                      "Android OS 8.1.0 API 27 (OnePlus, ONEPLUS A6003)",
		MsPlayed:                      207365,
		ConnCountry:                   "SE",
		IpAddrDecrypted:               "94.234.39.50",
		UserAgentDecrypted:            "unknown",
		ReasonStart:                   "clickrow",
		ReasonEnd:                     "endplay",
		Shuffle:                       false,
		Skipped:                       boolPtr(false),
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

	TestDataValidTrackOffline1 = models.SpotifyData{
		Timestamp:                     parseTime("2017-02-19T22:15:52Z"),
		UserName:                      "mahadi4",
		Platform:                      "Android OS 7.0 API 24 (OnePlus, ONEPLUS A3003)",
		MsPlayed:                      691,
		ConnCountry:                   "SE",
		IpAddrDecrypted:               "85.230.36.225",
		UserAgentDecrypted:            "unknown",
		ReasonStart:                   "fwdbtn",
		ReasonEnd:                     "fwdbtn",
		Shuffle:                       true,
		Skipped:                       boolPtr(false),
		Offline:                       true,
		OfflineTimestamp:              1487235592056,
		IncognitoMode:                 false,
		SpotifyTrackUri:               stringPtr("spotify:track:4kbj5MwxO1bq9wjT5g9HaA"),
		MasterMetadataTrackName:       stringPtr("Shut Up and Dance"),
		MasterMetadataAlbumArtistName: stringPtr("WALK THE MOON"),
		MasterMetadataAlbumAlbumName:  stringPtr("TALKING IS HARD"),
		EpisodeName:                   nil,
		EpisodeShowName:               nil,
		SpotifyEpisodeUri:             nil,
	}

	TestDataValidTrackOffline2 = models.SpotifyData{
		Timestamp:                     parseTime("2017-02-19T22:15:52Z"),
		UserName:                      "mahadi4",
		Platform:                      "Android OS 7.0 API 24 (OnePlus, ONEPLUS A3003)",
		MsPlayed:                      431,
		ConnCountry:                   "SE",
		IpAddrDecrypted:               "85.230.36.225",
		UserAgentDecrypted:            "unknown",
		ReasonStart:                   "fwdbtn",
		ReasonEnd:                     "fwdbtn",
		Shuffle:                       true,
		Skipped:                       nil,
		Offline:                       true,
		OfflineTimestamp:              1487235589259,
		IncognitoMode:                 false,
		SpotifyTrackUri:               stringPtr("spotify:track:6yro7PoSqXaUJup5oILqg6"),
		MasterMetadataTrackName:       stringPtr("Alesund"),
		MasterMetadataAlbumArtistName: stringPtr("Sun Kil Moon"),
		MasterMetadataAlbumAlbumName:  stringPtr("On Tour: A Documentary - The Soundtrack"),
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
		Skipped:                       nil,
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

	TestDataUnknown = models.SpotifyData{
		Timestamp:                     parseTime("2017-08-22T18:42:18Z"),
		UserName:                      "mahadi4",
		Platform:                      "Android OS 7.1.1 API 25 (OnePlus, ONEPLUS A3003)",
		MsPlayed:                      179383,
		ConnCountry:                   "SE",
		IpAddrDecrypted:               "77.218.242.182",
		UserAgentDecrypted:            "unknown",
		ReasonStart:                   "fwdbtn",
		ReasonEnd:                     "trackdone",
		Shuffle:                       true,
		Skipped:                       boolPtr(false),
		Offline:                       false,
		OfflineTimestamp:              1503427160116,
		IncognitoMode:                 false,
		MasterMetadataTrackName:       nil,
		MasterMetadataAlbumArtistName: nil,
		MasterMetadataAlbumAlbumName:  nil,
		SpotifyTrackUri:               nil,
		EpisodeName:                   nil,
		EpisodeShowName:               nil,
		SpotifyEpisodeUri:             nil,
	}
)

func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func parseTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}
