package youtube

import (
	"encoding/json"
)

// Config elements
type Config struct {
	Chat                        bool
	ChatChannelID               string
	ClientID                    string
	ClientSecret                string
	Enabled                     bool
	LastFollowersCount          int
	LastSubscribersCount        int
	PollingFrequency            int
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
	m.settings.Chat = true
	m.settings.ChatChannelID = "325241876758396928"
	m.settings.Enabled = true
	m.settings.LastFollowersCount = 10
	m.settings.LastSubscribersCount = 10
	m.settings.PollingFrequency = 7
	m.settings.Prefix = "<:youtube:328598129710596096> "
	m.settings.PadChannelFollowersOutput = 3
	m.settings.PadChannelViewersOutput = 3
	m.settings.PadChannelSubscribersOutput = 3

	// Check Raw Data
	if m.j.Config.IsInitialized() {
		if !m.j.Config.IsValidKey("YouTube") {
			m.j.Log.Message("YouTube", "Unable to find \"YouTube\" config section. Using defaults.")
		} else {

			errorCheck := json.Unmarshal(*m.j.Config.GetConfigData("YouTube"), &m.settings)
			if errorCheck != nil {
				m.j.Log.Message("Config", "Unable to properly parse YouTube Config, somethings may be wonky.")
			}
		}
	}

	// Sanitize
	//m.settings.ChatUsername = strings.ToLower(m.settings.ChatUsername)
}
