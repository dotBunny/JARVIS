package stats

import (
	"encoding/json"
)

// Config elements
type Config struct {
	PadCoffeeOutput int
	PadSavesOutput  int
}

// Initialize the Logging Module
func (m *Module) loadConfig() {
	// Create default general settings
	m.settings = new(Config)

	m.settings.PadCoffeeOutput = 2
	m.settings.PadSavesOutput = 2

	// Check Raw Data
	if m.j.Config.IsInitialized() {
		if !m.j.Config.IsValidKey("Stats") {
			m.j.Log.Message("Stats", "Unable to find \"Stats\" config section. Using defaults.")
		} else {

			errorCheck := json.Unmarshal(*m.j.Config.GetConfigData("Stats"), &m.settings)
			if errorCheck != nil {
				m.j.Log.Message("Config", "Unable to properly parse Stats Config, somethings may be wonky.")
			}
		}
	}
}
