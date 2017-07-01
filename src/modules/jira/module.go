package jira

import (
	"net/http"
	"time"

	"golang.org/x/oauth2"

	Core "../../core"
)

// const currentlyPlayingBase = "https://open.spotify.com/track/"

// Module Class
type Module struct {
	ticker *time.Ticker

	settings   *Config
	outputs    *Outputs
	data       *Data
	jiraClient *http.Client
	jiraOAuth  oauth2.Config
	jiraToken  string
	j          *Core.JARVIS
}

// Initialize the Logging Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS) {
	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance

	// Make sure flag is toggled off
	m.loadConfig()

	m.j.Log.RegisterChannel("JIRA", "blue", m.settings.Prefix)

	if !m.settings.Enabled {
		return
	}

	m.setupOutputs()
	m.setupData()

	// Create new authenticator with permissions
	m.authenticate()

	// Start the basic polling for information
	m.setupPolling()
	// m.setupEndpoints()
	// m.setupCommands()
}

// // IsEnabled for Usage
// func (m *Module) IsEnabled() bool {
// 	return m.settings.Enabled
// }

// Shutdown Module
func (m *Module) Shutdown() {
	if m != nil {
		if m.ticker != nil {
			m.ticker.Stop()
		}
	}
}
