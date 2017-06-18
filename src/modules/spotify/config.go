package spotify

import (
	"encoding/json"
)

// SpotifyConfig elements
type SpotifyConfig struct {
	Enabled             bool
	PollingFrequency    int
	ClientID            string
	ClientSecret        string
	TruncateTrackLength int
	TruncateTrackRunes  string
}

// Initialize the Logging Module
func (m *SpotifyModule) loadConfig() {
	// Create default general settings
	m.settings = new(SpotifyConfig)

	m.settings.Enabled = true
	m.settings.PollingFrequency = 5
	m.settings.ClientID = "7b90d69131194380a3734dfb818f8cb5"
	m.settings.ClientSecret = "530dab948cbd4d778ef58a124826a91c"
	m.settings.TruncateTrackLength = 85
	m.settings.TruncateTrackRunes = "..."

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
