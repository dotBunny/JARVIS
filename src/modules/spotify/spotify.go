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
	m.setupData()

	// Create new authenticator with permissions
	m.authenticate()

	// Start the basic polling for information
	m.setupPolling()
	m.setupEndpoints()
	m.setupCommands()
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
