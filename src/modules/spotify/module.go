package spotify

import (
	"time"

	Core "../../core"

	"github.com/zmb3/spotify"
)

// const currentlyPlayingBase = "https://open.spotify.com/track/"

// Module Class
type Module struct {
	ticker        *time.Ticker
	spotifyOAuth  spotify.Authenticator
	spotifyClient *spotify.Client
	stateHash     string
	settings      *Config
	outputs       *Outputs
	data          *Data
	j             *Core.JARVIS
}

// Initialize the Logging Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS) {
	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance

	// Make sure flag is toggled off
	m.loadConfig()

	m.j.Log.RegisterChannel("Spotify", "green", m.settings.Prefix)

	// If we're not enabled go no further
	if !m.IsEnabled() {
		return
	}

	m.setupOutputs()
	m.setupData()

	// Create new authenticator with permissions
	m.authenticate()

	// Start the basic polling for information

	m.setupEndpoints()
	m.setupCommands()

	m.Start()
}

// IsEnabled for Usage
func (m *Module) IsEnabled() bool {
	return m.settings.Enabled
}

// Shutdown Module
func (m *Module) Shutdown() {
	m.Stop()
}

// Start Spotify Polling / IRC
func (m *Module) Start() {
	m.setupPolling()
}

// Stop Spotify Polling / IRC
func (m *Module) Stop() {
	if m != nil {
		if m.ticker != nil {
			m.ticker.Stop()
		}
	}
}
