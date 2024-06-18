package main

import (
	"encoding/json"
	"fmt"
	"countDataSheets/models"
	"io"
	"os"
)

func main() {
	fileIndex := 9

  totalCount := 0
	for i := 0; i <= fileIndex; i++ {
		fileName := fmt.Sprintf("%s%d%s", "../rawSpotifyData/MyDataGo/endsong_", i, ".json")
		jsonFile, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		// Read the JSON file content
		byteValue, err := io.ReadAll(jsonFile)
		var spotifyData []models.SpotifyData
		err = json.Unmarshal(byteValue, &spotifyData)

		if err != nil {
			fmt.Println("kaos with unmarshal spotify data", err)
		}
		totalCount += getObjectCount(spotifyData)
	}
  fmt.Println(totalCount)
}

func getObjectCount(dataSheet []models.SpotifyData) (count int) {
	count = 0
	for i := range dataSheet {
		count = i
	}
	return count + 1
}
