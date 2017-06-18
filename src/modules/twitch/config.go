package twitch

import (
	"encoding/json"
)

// Config elements
type Config struct {
	Channel            string
	ChannelID          string
	ChatSync           bool
	ChatSyncChannelID  string
	ClientID           string
	ClientSecret       string
	Enabled            bool
	LastFollowersCount int
	PollingFrequency   int
	Username           string
	Prefix             string
}

// Initialize the Logging Module
func (m *Module) loadConfig() {

	// Create default general settings
	m.settings = new(Config)

	// TWitch Default Config
	m.settings.Channel = "#reapazor"
	m.settings.ChannelID = "21139969"
	m.settings.ChatSync = true
	m.settings.ChatSyncChannelID = "325241876758396928"
	m.settings.ClientID = "You need to set your ClientID"
	m.settings.ClientSecret = "You need to set your ClientSecret"
	m.settings.Enabled = true
	m.settings.LastFollowersCount = 10
	m.settings.PollingFrequency = 7
	m.settings.Username = "mod_jarvis"
	m.settings.Prefix = ":twitch: "

	// Check Raw Data
	if m.j.Config.IsInitialized() {
		if !m.j.Config.IsValidKey("Spotify") {
			m.j.Log.Message("Spotify", "Unable to find \"Spotify\" config section. Using defaults.")
		} else {

			errorCheck := json.Unmarshal(*m.j.Config.GetConfigData("Spotify"), &m.settings)
			if errorCheck != nil {
				m.j.Log.Message("Config", "Unable to properly parse Twitch Config, somethings may be wonky.")
			}
		}
	}
}
