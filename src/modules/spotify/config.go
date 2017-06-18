package spotify

import (
	"encoding/json"
)

// Config elements
type Config struct {
	Enabled             bool
	PollingFrequency    int
	ClientID            string
	ClientSecret        string
	TruncateTrackLength int
	TruncateTrackRunes  string
	Prefix              string
}

// Initialize the Logging Module
func (m *Module) loadConfig() {
	// Create default general settings
	m.settings = new(Config)

	m.settings.Enabled = true
	m.settings.PollingFrequency = 5
	m.settings.ClientID = "Your secret key needs to be in the config"
	m.settings.ClientSecret = "Your secret key needs to be in the config"
	m.settings.TruncateTrackLength = 85
	m.settings.TruncateTrackRunes = "..."
	m.settings.Prefix = ":spotify: "

	// Check Raw Data
	if m.j.Config.IsInitialized() {
		if !m.j.Config.IsValidKey("Spotify") {
			m.j.Log.Message("Spotify", "Unable to find \"Spotify\" config section. Using defaults.")
		} else {

			errorCheck := json.Unmarshal(*m.j.Config.GetConfigData("Spotify"), &m.settings)
			if errorCheck != nil {
				m.j.Log.Message("Config", "Unable to properly parse Spotify Config, somethings may be wonky.")
			}
		}
	}
}
