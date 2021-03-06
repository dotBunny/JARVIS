package twitch

// twitch.connect (reconnect - reauth)

// To get your userID
// curl -H 'Accept: application/vnd.twitchtv.v5+json' -H 'Client-ID: <CLIENT ID>' -X GET https://api.twitch.tv/kraken/users?login=<USERNAME>

import (
	"encoding/json"
	"time"

	"golang.org/x/oauth2"

	Core "../../core"
	Command "../command"
	"github.com/thoj/go-ircevent"

	"net/http"
	"strings"
)

// Module Class
type Module struct {
	ticker *time.Ticker

	irc *irc.Connection

	settings         *Config
	outputs          *Outputs
	data             *Data
	twitchClient     *http.Client
	twitchOAuth      oauth2.Config
	twitchToken      string
	twitchStreamName string
	j                *Core.JARVIS
	commandModule    *Command.Module
}

// Initialize the Logging Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS, commandModule *Command.Module) {

	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance
	m.commandModule = commandModule

	// Load Configuration
	m.loadConfig()

	m.j.Log.RegisterChannel("Twitch", "purple", m.settings.Prefix)

	m.setupOutputs()
	m.setupData()

	m.authenticate()

	m.setupEndpoints()

	// Some cached settings
	m.twitchStreamName = strings.TrimLeft(m.settings.Channel, "#")

	m.setupCommands()
	m.j.WebServer.RegisterParser("twitch", m.ParseWebContent)

	m.Start()
}

func (m *Module) getResponse(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		m.j.Log.Error("Twitch", "Unable to create request: "+url+", "+err.Error())
		return nil, nil
	}

	req.Header.Set("User-Agent", "JARVIS")
	req.Header.Set("Client-ID", m.settings.ClientID)
	req.Header.Set("Authorization", "OAuth "+m.twitchToken)
	req.Header.Set("Accept", "application/vnd.twitchtv.v5+json")

	return m.twitchClient.Do(req)
}

// IsEnabled for usage
func (m *Module) IsEnabled() bool {
	return m.settings.Enabled
}

// Shutdown Module
func (m *Module) Shutdown() {
	if m != nil {
		m.Stop()
	}
}

// Start Twitch Polling / IRC
func (m *Module) Start() {
	m.connectIRC()
	m.setupPolling()
}

// Stop Twitch Polling / IRC
func (m *Module) Stop() {
	if m != nil {
		if m.ticker != nil {
			m.ticker.Stop()
		}
		if m.irc != nil {
			m.irc.Quit()
		}
	}
}

func (m *Module) ParseWebContent(content string, mode string, r *http.Request) string {

	if mode == ".json" {

		if strings.Contains(content, "[[JARVIS.twitch]]") {
			// TODO: We should make data have a function to turn it into a map?
			responseMap := make(map[string]interface{})

			responseMap["LastFollower"] = m.data.LastFollower
			responseMap["LastFollowers"] = m.data.LastFollowers
			responseMap["LastSubscriber"] = m.data.LastSubscriber
			responseMap["LastSubscribers"] = m.data.LastSubscribers
			responseMap["ChannelFollowers"] = m.data.ChannelFollowers
			responseMap["ChannelViewers"] = m.data.ChannelViewers
			responseMap["ChannelGame"] = m.data.ChannelGame
			responseMap["StreamPreviewURL"] = m.data.StreamPreviewURL
			responseMap["Username"] = strings.Replace(m.settings.Channel, "#", "", -1)

			outputJSON, _ := json.Marshal(responseMap)
			content = strings.Replace(content, "[[JARVIS.twitch]]", string(outputJSON), -1)
		}

		if strings.Contains(content, "[[JARVIS.twitch.viewers]]") {
			outputJSON, _ := json.Marshal(m.data.Viewers)
			content = strings.Replace(content, "[[JARVIS.twitch.viewers]]", string(outputJSON), -1)
		}

	} else {

		if strings.Contains(content, "[[JARVIS.twitch.username]]") {
			content = strings.Replace(content, "[[JARVIS.twitch.username]]", strings.Replace(m.settings.Channel, "#", "", -1), -1)
		}
	}
	return content
}
