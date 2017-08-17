package jira

import (
	"time"

	Core "../../../core"
	"github.com/andygrunwald/go-jira"
)

// Module Class
type Module struct {
	ticker  *time.Ticker
	Polling bool

	settings   *Config
	outputs    *Outputs
	data       *Data
	jiraClient *jira.Client

	j *Core.JARVIS

	workingOn Core.DataModifier
}

type DataSetter func(string, bool)
type DataGetter func() string

// Initialize the Logging Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS, modifierInterface Core.DataModifier) {
	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance

	// TODO: I m not a fan of these call backs - they seem to be a source of issue
	m.workingOn = modifierInterface
	m.Polling = false

	// Make sure flag is toggled off
	m.loadConfig()
	m.j.Log.RegisterChannel("JIRA", "blue", m.settings.Prefix)

	if !m.settings.Enabled {
		return
	}

	m.setupOutputs()
	m.setupData()

	// Create new http Client
	var jiraError error
	m.jiraClient, jiraError = jira.NewClient(nil, m.settings.Instance)
	if jiraError != nil {
		m.j.Log.Error("JIRA", "Creating Client Failed. "+jiraError.Error())
		return
	}

	m.jiraClient.Authentication.SetBasicAuth(m.settings.BasicAuthUsername, m.settings.BasicAuthPassword)
	sessionCookie, errorCookie := m.jiraClient.Authentication.AcquireSessionCookie(m.settings.BasicAuthUsername, m.settings.BasicAuthPassword)
	if errorCookie != nil || sessionCookie == false {
		m.j.Log.Error("JIRA", "Cookie Request Failed. "+errorCookie.Error())
		return
	}

	// Start the basic polling for information
	if m.workingOn.ShouldUpdate() {
		m.setupPolling()
	}
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
	if !m.Polling {
		m.setupPolling()
	}
}

// Stop Spotify Polling / IRC
func (m *Module) Stop() {
	if m != nil {
		if m.ticker != nil {
			m.ticker.Stop()
		}
	}
	m.Polling = false
}
