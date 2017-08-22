package ads

import (
	"encoding/json"
)

// Config elements
type Config struct {
	Definitions []Ad
}
type Ad struct {
	Key      string   `json:"Key"`
	Channels []string `json:"Channels"`
	Content  []string `json:"Content"`
	Interval int      `json:"Interval"`
}

// Initialize the Logging Module
func (m *Module) loadConfig() {

	m.j.Config.LoadConfig("ads.json", "Ads")
	// Create default general settings
	m.settings = new(Config)

	// Check Raw Data
	if m.j.Config.IsInitialized() {
		if !m.j.Config.IsValidKey("Ads") {
			m.j.Log.Message("Config", "Unable to find \"Ads\" config section. Using defaults.")
		} else {

			errorCheck := json.Unmarshal(m.j.Config.GetConfigData("Ads"), &m.settings)
			if errorCheck != nil {
				m.j.Log.Message("Config", "Unable to properly parse Ads Config, somethings may be wonky.")

				// Report Problem
				m.errorCount++
			}
		}
	}
}
