package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type AuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	artistId := "7BMccF0hQFBpP6417k1OtQ"
	client := http.Client{}
	spotifyUrl := os.Getenv("SPOTIFY_API_URL")

	token, err := getBearerToken()
  if err != nil {
    fmt.Println(err)
  }

	request, err := http.NewRequest("GET", spotifyUrl+"/artists/"+artistId, nil)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("Authorization", "Bearer " + token.AccessToken)
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	s := string(body)
	fmt.Println(s)

}

func getBearerToken() (*AuthTokenResponse, error) {
	client := http.Client{}
	spotifyClient := os.Getenv("SPOTIFY_CLIENT")
	spotifySecret := os.Getenv("SPOTIFY_SECRET")
	credentials := spotifyClient + ":" + spotifySecret
	credentialsEncoded := base64.StdEncoding.EncodeToString([]byte(credentials))

	requestBody := url.Values{}
	requestBody.Set("grant_type", "client_credentials")

	request, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(requestBody.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Add("Authorization", "Basic "+credentialsEncoded)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	var authTokenResponse AuthTokenResponse
	if err := json.NewDecoder(response.Body).Decode(&authTokenResponse); err != nil {
		return nil, err
	}

	return &authTokenResponse, nil
}
