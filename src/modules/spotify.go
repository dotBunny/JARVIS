package modules

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"io/ioutil"

	Core "../core"
	"github.com/zmb3/spotify"
)

// SpotifyData contains the live data we compare against
type SpotifyData struct {
	LastInfoData  []byte
	LastImageData []byte
	DurationMS    int
	PlayedMS      int
}

var (
	spotifyData            SpotifyData
	spotifyLatestSongPath  string
	spotifyLatestImagePath string
	auth                   spotify.Authenticator

	ch    = make(chan *spotify.Client)
	state = "JARVIS"
)

// InitializeSpotify Module
func InitializeSpotify(config *Core.Config) *spotify.Client {

	// Create our output paths
	spotifyLatestSongPath = filepath.Join(config.General.OutputPath, "Spotify_LatestSong.txt")
	spotifyLatestImagePath = filepath.Join(config.General.OutputPath, "Spotify_LatestImage.jpg")

	// Check twitchLatestFollowerPath
	if _, err := os.Stat(spotifyLatestSongPath); os.IsNotExist(err) {
		ioutil.WriteFile(spotifyLatestSongPath, nil, 0755)
	}

	// Check twitchLatestFollowerPath
	if _, err := os.Stat(spotifyLatestImagePath); os.IsNotExist(err) {
		ioutil.WriteFile(spotifyLatestImagePath, nil, 0755)
	}

	// Create new authenticator with permissions
	auth = spotify.NewAuthenticator("http://localhost:"+config.General.ServerPort+config.Spotify.Callback,
		spotify.ScopeUserReadCurrentlyPlaying,
		spotify.ScopeUserReadRecentlyPlayed)

	// Start Login AUTH Procedures
	auth.SetAuthInfo(config.Spotify.ClientID, config.Spotify.ClientSecret)

	// TODO: Add something to retain login info?

	// Add Endpoint for Callbac
	Core.AddEndpoint(config.Spotify.Callback, spotifyCompleteAuthentication)

	url := auth.AuthURL(state)
	Core.Log("SPOTIFY", "IMPORTANT", "Please log in to Spotify by visiting the following page in your browser:\n\n"+url+"\n")

	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	Core.Log("SPOTIFY", "LOG", "You are logged in as: "+user.ID)

	return client
}

// PollSpotify For Updates
func PollSpotify(client *spotify.Client, config *Core.Config) {
	spotifyGetCurrentlyPlaying(client, config)
}

func spotifyCompleteAuthentication(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := auth.NewClient(tok)
	fmt.Fprintf(w, "Login Completed!")
	ch <- &client
}

func spotifyGetCurrentlyPlaying(client *spotify.Client, config *Core.Config) {
	state, err := client.PlayerCurrentlyPlaying()

	if err != nil {
		Core.Log("SPOTIFY", "ERROR", "Unable to retrieve currently playing song")
	} else {

		// Handle Basic Track Information
		var buffer bytes.Buffer
		var artistCount = len(state.Item.Artists)
		for i := 0; i < artistCount; i++ {
			buffer.WriteString(state.Item.Artists[i].Name)
			if i < (artistCount - 1) {
				buffer.WriteString(", ")
			}
		}
		buffer.WriteString(" - ")
		buffer.WriteString(state.Item.Name)

		if !bytes.Equal(buffer.Bytes(), spotifyData.LastInfoData) {

			Core.Log("SPOTIFY", "LOG", buffer.String())

			Core.SaveFile(buffer.Bytes(), spotifyLatestSongPath)

			spotifyData.LastInfoData = buffer.Bytes()
			spotifyData.DurationMS = state.Item.Duration
			spotifyData.PlayedMS = state.Progress

			// Clear buffer
			buffer.Reset()

			// Image Data - Check we have 1
			if len(state.Item.Album.Images) > 0 {

				writer := bufio.NewWriter(&buffer)
				state.Item.Album.Images[0].Download(writer)

				if !bytes.Equal(buffer.Bytes(), spotifyData.LastImageData) {

					Core.SaveFile(buffer.Bytes(), spotifyLatestImagePath)
					spotifyData.LastImageData = buffer.Bytes()
				}
			}
		}
	}
}

// SpotifyRender UI component
func SpotifyRender(config *Core.Config) {

}
