package main

import (
	"encoding/json"
	"fmt"
	"github.com/mahadia/mahadia-spotifyData/countDataSheets/models"
	"io"
	"os"
  "strings"
)

func main() {
	// countSheets()
	// fileName := fmt.Sprintf("%s%d%s", "../rawSpotifyData/MyDataGo/endsong_", i, ".json")
	// jsonFile, err := os.Open("../rawSpotifyData/endsong_0.json")

	jsonFile, err := os.Open("../rawSpotifyData/smallSample.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	var spotifyMiniData []models.SpotifyData
	err = json.Unmarshal(byteValue, &spotifyMiniData)
	if err != nil {
		fmt.Println("kaos with unmarshal spotify data", err)
	}
	for _, v := range spotifyMiniData {
		processData(v)
	}
}

func processData(playback models.SpotifyData) error {
  b := strings.TrimPrefix(playback.SpotifyTrackUri, "spotify:track:")
	fmt.Println(playback.SpotifyTrackUri)
	fmt.Println(b)
	return nil
}

func countSheets() (int, error) {
	fileIndex := 9

	totalCount := 0
	for i := 0; i <= fileIndex; i++ {
		fileName := fmt.Sprintf("%s%d%s", "../rawSpotifyData/MyDataGo/endsong_", i, ".json")
		jsonFile, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err)
			return 0, nil
		}
		defer jsonFile.Close()

		// Read the JSON file content
		byteValue, err := io.ReadAll(jsonFile)
		var spotifyData []models.SpotifyData
		err = json.Unmarshal(byteValue, &spotifyData)

		if err != nil {
			fmt.Println("kaos with unmarshal spotify data", err)
			return 0, err
		}
		totalCount += getObjectCount(spotifyData)
	}
	return totalCount, nil
}

func getObjectCount(dataSheet []models.SpotifyData) (count int) {
	count = 0
	for i := range dataSheet {
		count = i
	}
	return count + 1
}
