package twitch

import (
	"encoding/json"
	"strings"
)

// Config elements
type Config struct {
	Channel                     string
	ChannelID                   string
	Chat                        bool
	ChatChannelID               string
	ChatUsername                string
	ChatOAuth                   string
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
	m.settings.Channel = "#reapazor"
	m.settings.ChannelID = "21139969"
	m.settings.Chat = true
	m.settings.ChatChannelID = "325241876758396928"
	m.settings.ChatUsername = "mod_jarvis"
	m.settings.ChatOAuth = "You need to make one"
	m.settings.ClientID = "You need to set your ClientID"
	m.settings.ClientSecret = "You need to set your ClientSecret"
	m.settings.Enabled = true
	m.settings.LastFollowersCount = 10
	m.settings.LastSubscribersCount = 10
	m.settings.PollingFrequency = 7
	m.settings.Prefix = ":twitch: "
	m.settings.PadChannelFollowersOutput = 3
	m.settings.PadChannelViewersOutput = 3
	m.settings.PadChannelSubscribersOutput = 3

	// Check Raw Data
	if m.j.Config.IsInitialized() {
		m.j.Config.LoadConfig("twitch.json", "Twitch")
		if !m.j.Config.IsValidKey("Twitch") {
			m.j.Log.Message("Twitch", "Unable to find \"Twitch\" config section. Using defaults.")
		} else {

			errorCheck := json.Unmarshal(m.j.Config.GetConfigData("Twitch"), &m.settings)
			if errorCheck != nil {
				m.j.Log.Message("Config", "Unable to properly parse Twitch Config, somethings may be wonky.")
			}
		}
	}

	// Sanitize
	m.settings.ChatUsername = strings.ToLower(m.settings.ChatUsername)
}
