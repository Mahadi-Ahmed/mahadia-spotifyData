package main

import (
	"encoding/json"
	"fmt"
	"goSpotify/models"
	"io"
	"os"
)

func main() {
	fmt.Println("hej")
	jsonFile, err := os.Open("../rawSpotifyData/smallSample.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// Read the JSON file content
	byteValue, err := io.ReadAll(jsonFile)
	var spotifyData []models.SpotifyData
	err = json.Unmarshal(byteValue, &spotifyData)
	if err != nil {
		fmt.Println("kaos", err)
	}
	fmt.Println(spotifyData)
	for _, data := range spotifyData {
		fmt.Printf("Timestamp: %s\n", data.Timestamp)
		fmt.Printf("Username: %s\n", data.Username)
		fmt.Printf("Platform: %s\n", data.Platform)
		fmt.Printf("MsPlayed: %d\n", data.MsPlayed)
		fmt.Printf("ConnCountry: %s\n", data.ConnCountry)
		fmt.Printf("IpAddrDecrypted: %s\n", data.IpAddrDecrypted)
		fmt.Printf("UserAgentDecrypted: %s\n", data.UserAgentDecrypted)
		fmt.Printf("MasterMetadataTrackName: %s\n", data.MasterMetadataTrackName)
		fmt.Printf("MasterMetadataAlbumArtistName: %s\n", data.MasterMetadataAlbumArtistName)
		fmt.Printf("MasterMetadataAlbumAlbumName: %s\n", data.MasterMetadataAlbumAlbumName)
		fmt.Printf("SpotifyTrackUri: %s\n", data.SpotifyTrackUri)
		fmt.Printf("EpisodeName: %v\n", data.EpisodeName)
		fmt.Printf("EpisodeShowName: %v\n", data.EpisodeShowName)
		fmt.Printf("SpotifyEpisodeUri: %v\n", data.SpotifyEpisodeUri)
		fmt.Printf("ReasonStart: %s\n", data.ReasonStart)
		fmt.Printf("ReasonEnd: %s\n", data.ReasonEnd)
		fmt.Printf("Shuffle: %v\n", data.Shuffle)
		fmt.Printf("Skipped: %v\n", data.Skipped)
		fmt.Printf("Offline: %v\n", data.Offline)
		fmt.Printf("OfflineTimestamp: %d\n", data.OfflineTimestamp)
		fmt.Printf("IncognitoMode: %v\n", data.IncognitoMode)
		fmt.Println()
	}
}
