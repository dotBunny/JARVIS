package modules

import (
	"time"

	Core "../core"
	"github.com/zmb3/spotify"
)

// SpotifyConfig elements
type SpotifyConfig struct {
	Enabled             bool
	PollingFrequency    string
	ClientID            string
	ClientSecret        string
	RedirectURI         string
	TruncateTrackLength int
	TruncateTrackRunes  string
}

// import (
// 	"bufio"
// 	"bytes"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"path"
// 	"path/filepath"
// 	"strconv"
// 	"time"

// 	Core "../core"
// 	"github.com/skratchdot/open-golang/open"
// 	"github.com/zmb3/spotify"
// )

// var (
// 	ch = make(chan *spotify.Client)
// )

// const currentlyPlayingBase = "https://open.spotify.com/track/"

// SpotifyModule Class
type SpotifyModule struct {
	LastInfoData        []byte
	LastImageData       []byte
	DurationMS          int
	PlayedMS            int
	CurrentlyPlaying    bool
	CurrentlyPlayingURL string
	Ticker              *time.Ticker

	auth      spotify.Authenticator
	songPath  string
	imagePath string
	urlPath   string

	stateHash     string
	authenticated bool
	settings      *SpotifyConfig
	j             *Core.JARVIS
}

// Initialize the Logging Module
func (m *SpotifyModule) Initialize(jarvisInstance *Core.JARVIS) {

	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance

	// Make sure flag is toggled off
	m.authenticated = false

	// Create default general settings
	m.settings = new(SpotifyConfig)
}

// Connect to Spotify
func (m *SpotifyModule) Connect() {

	m.stateHash = Core.RandomString(5)

}

// // GetCurrentlyPlayingMessage
// func (m *SpotifyModule) GetCurrentlyPlayingMessage() string {
// 	return string(m.LastInfoData) + " Mau5 " + m.CurrentlyPlayingURL
// }

// 	// Build Paths
// 	m.songPath = filepath.Join(m.config.General.OutputPath, "Spotify_LatestSong.txt")
// 	m.urlPath = filepath.Join(m.config.General.OutputPath, "Spotify_LatestURL.txt")
// 	m.imagePath = filepath.Join(m.config.General.OutputPath, "Spotify_LatestImage.jpg")

// 	// Make sure file slots are there, but dont do anything else, we want them overwritten
// 	if config.Spotify.Output {
// 		Core.Touch(m.songPath)
// 		Core.Touch(m.urlPath)
// 		Core.Touch(m.imagePath)
// 	}

// 	// Write default image to server
// 	defaultImage, err := ioutil.ReadFile(path.Join(m.config.AppDir, "resources", "overlay", "content", "img", "jarvis-spotify.jpg"))
// 	if err == nil {
// 		m.LastImageData = defaultImage
// 		Core.SaveFile(m.LastImageData, m.imagePath)
// 	}

// 	// Create new authenticator with permissions
// 	m.auth = spotify.NewAuthenticator("http://localhost:"+strconv.Itoa(m.config.General.ServerPort)+m.config.Spotify.Callback,
// 		spotify.ScopeUserReadCurrentlyPlaying,
// 		spotify.ScopeUserReadRecentlyPlayed,
// 		spotify.ScopeUserModifyPlaybackState)

// 	// Start Login AUTH Procedures
// 	m.auth.SetAuthInfo(m.config.Spotify.ClientID, m.config.Spotify.ClientSecret)

// 	// TODO: Add something to retain login info?

// 	// Add Endpoint for Callbac
// 	Core.RegisterEndpoint(m.config.Spotify.Callback, m.authenticateCallback)

// 	url := m.auth.AuthURL(m.state)
// 	Core.Log("SPOTIFY", "IMPORTANT", "Please log in to Spotify (URL copied to your clipboard as well): "+url)
// 	Core.CopyToClipboard(url)

// 	// Pop open browser window
// 	if m.config.Spotify.AutoLogin {
// 		open.Run(url)
// 	}

// 	// wait for Auth to complete
// 	client := <-ch

// 	// Add Endpoints
// 	Core.RegisterEndpoint("/spotify/track", m.endpointTrack)
// 	Core.RegisterEndpoint("/spotify/image", m.endpointImage)

// 	// use the client to make calls that require authorization
// 	user, err := client.CurrentUser()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	Core.Log("SPOTIFY", "LOG", "You are logged in as: "+user.ID)

// 	// Assign Client
// 	m.client = client

// 	// Create Ticker
// 	spotifyPollingFrequency, spotifyPollingError := time.ParseDuration(m.config.Spotify.PollingFrequency)
// 	if spotifyPollingError != nil {
// 		spotifyPollingFrequency, _ = time.ParseDuration("5s")
// 	}
// 	m.Ticker = time.NewTicker(spotifyPollingFrequency)

// 	console.AddHandler("/spotify.next", "Skips to the next track in the user's Spotify queue.", m.consoleNextTrack)
// 	console.AddAlias("/next", "/spotify.next")
// 	console.AddAlias("/n", "/spotify.next")
// 	console.AddAlias("/skip", "/spotify.next")
// 	console.AddHandler("/spotify.pause", "Pause/Play the current track in Spotify.", m.consolePausePlay)
// 	console.AddAlias("/p", "/spotify.pause")
// 	console.AddHandler("/spotify.stats", "Display some stats from Spotify.", m.consoleStats)
// 	console.AddHandler("/spotify.update", "Force polling Spotify for updates.", m.consoleUpdate)
// }

// // Loop awaiting ticker
// func (m *SpotifyModule) Loop() {
// 	for {
// 		select {
// 		case <-m.Ticker.C:
// 			m.Poll()
// 		}
// 	}
// }

// // Poll For Updates
// func (m *SpotifyModule) Poll() {
// 	m.pollCurrentlyPlaying()
// }

// // Shutdown Module
// func (m *SpotifyModule) Shutdown() {
// 	if m != nil {
// 		if m.Ticker != nil {
// 			m.Ticker.Stop()
// 		}
// 	}
// }

// func (m *SpotifyModule) authenticateCallback(w http.ResponseWriter, r *http.Request) {
// 	tok, err := m.auth.Token(m.state, r)
// 	if err != nil {
// 		http.Error(w, "Couldn't get token", http.StatusForbidden)
// 		log.Fatal(err)
// 	}
// 	if st := r.FormValue("state"); st != m.state {
// 		http.NotFound(w, r)
// 		log.Fatalf("State mismatch: %s != %s\n", st, m.state)
// 	}
// 	// use the token to get an authenticated client
// 	client := m.auth.NewClient(tok)

// 	// Output and attempt to close tab
// 	// if m.config.Spotify.AutoLogin {
// 	// 	Core.Log("SPOTIFY", "LOG", "Login Detected. Closing Tab")
// 	// 	fmt.Fprintf(w, "<html><body>Spotify Login Complete<script type=\"text/javascript\">window.close();</script></body></html>")
// 	// } else {
// 	fmt.Fprintf(w, "Login Completed! Please close this tab/window.")
// 	// }
// 	ch <- &client
// }

// func (m *SpotifyModule) consoleNextTrack(args string) {
// 	Core.Log("SPOTIFY", "LOG", "Next Track!")
// 	m.client.Next()
// }
// func (m *SpotifyModule) consolePausePlay(args string) {
// 	if m.CurrentlyPlaying {
// 		Core.Log("SPOTIFY", "LOG", "Paused")
// 		m.client.Pause()
// 		m.CurrentlyPlaying = false
// 	} else {
// 		Core.Log("SPOTIFY", "LOG", "Playing")
// 		m.client.Play()
// 		m.CurrentlyPlaying = true
// 	}
// }

// func (m *SpotifyModule) consoleStats(input string) {

// 	if m.DurationMS == 0 {

// 		Core.Log("SPOTIFY", "LOG", "Currently playing "+string(m.LastInfoData)+" (?)")
// 	} else {
// 		percentComplete := float64(((m.PlayedMS / m.DurationMS) * 100))
// 		Core.Log("SPOTIFY", "LOG", "Currently playing "+string(m.LastInfoData)+" ("+fmt.Sprint(Core.Round(percentComplete, .5, 2))+"%)")
// 	}
// }

// func (m *SpotifyModule) consoleUpdate(input string) {
// 	m.Poll()
// 	Core.Log("SPOTIFY", "LOG", "Force Update")
// }

// func (m *SpotifyModule) endpointImage(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "image/jpeg")
// 	w.Header().Set("Content-Length", strconv.Itoa(len(m.LastImageData)))
// 	if _, err := w.Write(m.LastImageData); err != nil {
// 		Core.Log("SPOTIFY", "ERROR", "Unable to write image")
// 	}
// }

// func (m *SpotifyModule) endpointTrack(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, string(m.LastInfoData))
// }

// func (m *SpotifyModule) pollCurrentlyPlaying() {
// 	state, err := m.client.PlayerCurrentlyPlaying()

// 	if err != nil {
// 		Core.Log("SPOTIFY", "ERROR", "Unable to retrieve currently playing song")
// 	} else {

// 		// Handle Basic Track Information
// 		var buffer bytes.Buffer
// 		var artistCount = len(state.Item.Artists)
// 		for i := 0; i < artistCount; i++ {
// 			buffer.WriteString(state.Item.Artists[i].Name)
// 			if i < (artistCount - 1) {
// 				buffer.WriteString(", ")
// 			}
// 		}
// 		buffer.WriteString(" - ")
// 		buffer.WriteString(state.Item.Name)

// 		if buffer.Len() > m.config.Spotify.TruncateTrackLength {
// 			buffer.Truncate(m.config.Spotify.TruncateTrackLength)
// 			buffer.WriteString(m.config.Spotify.TruncateTrackRunes)
// 		}

// 		if !bytes.Equal(buffer.Bytes(), m.LastInfoData) {

// 			Core.Log("SPOTIFY", "LOG", buffer.String())

// 			if m.config.Spotify.Output {
// 				Core.SaveFile(buffer.Bytes(), m.songPath)
// 				Core.SaveFile([]byte(m.CurrentlyPlayingURL), m.urlPath)
// 			}

// 			m.LastInfoData = buffer.Bytes()
// 			m.DurationMS = state.Item.Duration
// 			m.PlayedMS = state.Progress
// 			m.CurrentlyPlaying = state.Playing

// 			// TODO : Get the URI / name / cache it for sending to channel
// 			m.CurrentlyPlayingURL = state.Item.ExternalURLs["spotify"]

// 			// Clear buffer
// 			buffer.Reset()

// 			// Image Data - Check we have 1
// 			if len(state.Item.Album.Images) > 0 {

// 				writer := bufio.NewWriter(&buffer)
// 				state.Item.Album.Images[0].Download(writer)

// 				if !bytes.Equal(buffer.Bytes(), m.LastImageData) {

// 					Core.SaveFile(buffer.Bytes(), m.imagePath)
// 					m.LastImageData = buffer.Bytes()
// 				}
// 				writer.Flush()

// 			}
// 		}
// 		buffer.Reset()
// 	}
// 	state = nil
// }
