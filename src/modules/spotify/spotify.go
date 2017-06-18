package spotify

import (
	"time"

	Core "../../core"

	"github.com/zmb3/spotify"
)

// const currentlyPlayingBase = "https://open.spotify.com/track/"

// SpotifyModule Class
type SpotifyModule struct {
	ticker        *time.Ticker
	spotifyOAuth  spotify.Authenticator
	spotifyClient *spotify.Client
	stateHash     string
	authenticated bool
	settings      *SpotifyConfig
	outputs       *SpotifyOutputs
	data          *SpotifyData
	j             *Core.JARVIS
}

// Connect to Spotify
func (m *SpotifyModule) Connect() {
	// Bail if not enabled
	if !m.IsEnabled() {
		return
	}

	m.stateHash = Core.RandomString(5)
}

// Initialize the Logging Module
func (m *SpotifyModule) Initialize(jarvisInstance *Core.JARVIS) {
	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance

	// Make sure flag is toggled off
	m.authenticated = false

	m.loadConfig()

	// If we're not enabled go no further
	if !m.IsEnabled() {
		return
	}

	m.setupOutputs()

	// Create new authenticator with permissions
	m.authenticate()

	// Start the basic polling for information
	m.setupPolling()
	m.setupEndpoints()
}

// IsEnabled for Usage
func (m *SpotifyModule) IsEnabled() bool {
	return m.settings.Enabled
}

// Shutdown Module
func (m *SpotifyModule) Shutdown() {
	if m != nil {
		if m.ticker != nil {
			m.ticker.Stop()
		}
	}
}

// // GetCurrentlyPlayingMessage
// func (m *SpotifyModule) GetCurrentlyPlayingMessage() string {
// 	return string(m.LastInfoData) + " Mau5 " + m.CurrentlyPlayingURL
// }

// 	console.AddHandler("/spotify.next", "Skips to the next track in the user's Spotify queue.", m.consoleNextTrack)
// 	console.AddAlias("/next", "/spotify.next")
// 	console.AddAlias("/n", "/spotify.next")
// 	console.AddAlias("/skip", "/spotify.next")
// 	console.AddHandler("/spotify.pause", "Pause/Play the current track in Spotify.", m.consolePausePlay)
// 	console.AddAlias("/p", "/spotify.pause")
// 	console.AddHandler("/spotify.stats", "Display some stats from Spotify.", m.consoleStats)
// 	console.AddHandler("/spotify.update", "Force polling Spotify for updates.", m.consoleUpdate)
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
