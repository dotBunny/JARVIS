package jira

import (
	"encoding/json"
)

// Config elements
type Config struct {
	BasicAuthUsername string
	BasicAuthPassword string
	Enabled           bool
	PollingFrequency  int
	Prefix            string
	Query             string
}

// Initialize the Logging Module
func (m *Module) loadConfig() {
	// Create default general settings
	m.settings = new(Config)

	m.settings.Enabled = true
	m.settings.PollingFrequency = 7
	m.settings.Prefix = ":jira: "

	// Check Raw Data
	if m.j.Config.IsInitialized() {
		if !m.j.Config.IsValidKey("JIRA") {
			m.j.Log.Message("JIRA", "Unable to find \"JIRA\" config section. Using defaults.")
		} else {

			errorCheck := json.Unmarshal(*m.j.Config.GetConfigData("JIRA"), &m.settings)
			if errorCheck != nil {
				m.j.Log.Message("Config", "Unable to properly parse JIRA Config, somethings may be wonky.")
			}
		}
	}
}
