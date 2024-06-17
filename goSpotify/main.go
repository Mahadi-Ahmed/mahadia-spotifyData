package main

import (
	"encoding/json"
	"fmt"
	"goSpotify/models"
	"io"
	"os"
)

func main() {
	endsong_0, err := os.Open("../rawSpotifyData/endsong_0.json")

	if err != nil {
		fmt.Println(err)
	}
	defer endsong_0.Close()

	// Read the JSON file content
	byteValue, err := io.ReadAll(endsong_0)
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
