package tasks

import (
	"encoding/json"
)

// Config elements
type Config struct {
	UseJIRAByDefault bool
}

// Initialize the Logging Module
func (m *Module) loadConfig() {

	// Create default general settings
	m.settings = new(Config)

	// Check Raw Data
	if m.j.Config.IsInitialized() {
		if !m.j.Config.IsValidKey("Tasks") {
			m.j.Log.Message("Config", "Unable to find \"Tasks\" config section. Using defaults.")
		} else {

			errorCheck := json.Unmarshal(*m.j.Config.GetConfigData("Tasks"), &m.settings)
			if errorCheck != nil {
				m.j.Log.Message("Config", "Unable to properly parse Tasks Config, somethings may be wonky.")
			}
		}
	}
}
