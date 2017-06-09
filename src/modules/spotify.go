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

const redirectURI = "http://localhost:8080/callback"

var (
	auth  = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadCurrentlyPlaying, spotify.ScopeUserReadRecentlyPlayed)
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

// SpotifyLogin authenticates with spotify
func SpotifyLogin(config *Core.SpotifyConfig) *spotify.Client {

	auth.SetAuthInfo(config.ClientID, config.ClientSecret)
	// first start an HTTP server
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("Got request for:", r.URL.String())
		})

	go http.ListenAndServe(":8080", nil)

	url := auth.AuthURL(state)
	Core.Log2("Please log in to Spotify by visiting the following page in your browser:\n\n", url)

	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	Core.Log2("You are logged in as:", user.ID)

	// Pathing Check
	os.MkdirAll(filepath.Dir(config.CurrentInfoPath), 0755)

	// Info Path
	if _, err := os.Stat(config.CurrentInfoPath); os.IsNotExist(err) {
		ioutil.WriteFile(config.CurrentInfoPath, nil, 0755)
	}

	// Image Path
	if _, err := os.Stat(config.CurrentImagePath); os.IsNotExist(err) {
		ioutil.WriteFile(config.CurrentImagePath, nil, 0755)
	}

	return client
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
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

// SpotifyPoll For Current Playing
func SpotifyPoll(client *spotify.Client, config *Core.SpotifyConfig) {

	state, err := client.PlayerCurrentlyPlaying()

	if err != nil {
		Core.Log("Unable to retrieve currently playing song")
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

		if !bytes.Equal(buffer.Bytes(), config.LastInfoData) {

			Core.Log(buffer.String())

			Core.SaveFile(buffer.Bytes(), config.CurrentInfoPath)
			config.LastInfoData = buffer.Bytes()

			// Clear buffer
			buffer.Reset()

			// Image Data - Check we have 1
			if len(state.Item.Album.Images) > 0 {

				writer := bufio.NewWriter(&buffer)
				state.Item.Album.Images[0].Download(writer)

				if !bytes.Equal(buffer.Bytes(), config.LastImageData) {

					Core.SaveFile(buffer.Bytes(), config.CurrentImagePath)
					config.LastImageData = buffer.Bytes()
				}
			}
		}
	}
}
