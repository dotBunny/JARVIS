package jira

import (
	"time"

	Core "../../core"
	Stats "../stats"
	"github.com/andygrunwald/go-jira"
)

// const currentlyPlayingBase = "https://open.spotify.com/track/"

// Module Class
type Module struct {
	ticker *time.Ticker

	settings   *Config
	outputs    *Outputs
	data       *Data
	jiraClient *jira.Client
	j          *Core.JARVIS
	stats      *Stats.Module
}

// Initialize the Logging Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS, statsModule *Stats.Module) {
	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance
	m.stats = statsModule

	// Make sure flag is toggled off
	m.loadConfig()
	m.j.Log.RegisterChannel("JIRA", "blue", m.settings.Prefix)

	if !m.settings.Enabled {
		return
	}

	m.stats.UseJIRAForWork = true
	m.setupOutputs()
	m.setupData()

	// Create new http Client
	var jiraError error
	m.jiraClient, jiraError = jira.NewClient(nil, m.settings.Instance)
	if jiraError != nil {
		m.j.Log.Error("JIRA", jiraError.Error())
		return
	}
	m.jiraClient.Authentication.SetBasicAuth(m.settings.BasicAuthUsername, m.settings.BasicAuthPassword)
	sessionCookie, errorCookie := m.jiraClient.Authentication.AcquireSessionCookie(m.settings.BasicAuthUsername, m.settings.BasicAuthPassword)
	if errorCookie != nil || sessionCookie == false {
		m.j.Log.Error("JIRA", errorCookie.Error())
		return
	}

	// Start the basic polling for information
	m.setupPolling()
	m.setupEndpoints()
	// m.setupCommands()
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
