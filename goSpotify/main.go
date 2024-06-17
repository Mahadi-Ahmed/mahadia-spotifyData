package main

import (
	"encoding/json"
	"fmt"
	"goSpotify/models"
	"io"
	"os"
)

func main() {
	// jsonFile, err := os.Open("../rawSpotifyData/smallSample.json")
	jsonFile, err := os.Open("../rawSpotifyData/endsong_0.json")

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
  var d0 = getObjectCount(spotifyData)
  fmt.Println(d0)

}

func getObjectCount(dataSheet []models.SpotifyData) (count int) {
  count = 0
  for i:= range dataSheet {
    count = i
  }
  return count + 1
}
