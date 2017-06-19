package twitch

import (
	"encoding/json"
)

// Config elements
type Config struct {
	Channel                     string
	ChannelID                   string
	ChatSync                    bool
	ChatSyncChannelID           string
	ClientID                    string
	ClientSecret                string
	Enabled                     bool
	LastFollowersCount          int
	LastSubscribersCount        int
	PollingFrequency            int
	Username                    string
	Prefix                      string
	PadChannelFollowersOutput   int
	PadChannelViewersOutput     int
	PadChannelSubscribersOutput int
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
	m.settings.LastSubscribersCount = 10
	m.settings.PollingFrequency = 7
	m.settings.Username = "mod_jarvis"
	m.settings.Prefix = ":twitch: "
	m.settings.PadChannelFollowersOutput = 3
	m.settings.PadChannelViewersOutput = 3
	m.settings.PadChannelSubscribersOutput = 3

	// Check Raw Data
	if m.j.Config.IsInitialized() {
		if !m.j.Config.IsValidKey("Twitch") {
			m.j.Log.Message("Twitch", "Unable to find \"Twitch\" config section. Using defaults.")
		} else {

			errorCheck := json.Unmarshal(*m.j.Config.GetConfigData("Twitch"), &m.settings)
			if errorCheck != nil {
				m.j.Log.Message("Config", "Unable to properly parse Twitch Config, somethings may be wonky.")
			}
		}
	}
}
