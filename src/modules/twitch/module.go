package twitch

// twitch.connect (reconnect - reauth)

// To get your userID
// curl -H 'Accept: application/vnd.twitchtv.v5+json' -H 'Client-ID: <CLIENT ID>' -X GET https://api.twitch.tv/kraken/users?login=<USERNAME>

import (
	"time"

	"golang.org/x/oauth2"

	Core "../../core"
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
}

// Initialize the Logging Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS) {

	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance

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
