package youtube

//* You can determine your [ChannelID] by going to https://www.youtube.com/account_advanced and copying the string after "YouTube Channel ID".
import (
	"net/http"
	"time"

	"golang.org/x/oauth2"
	youtube "google.golang.org/api/youtube/v3"

	Core "../../core"
)

// Module Class
type Module struct {
	ticker *time.Ticker

	settings          *Config
	outputs           *Outputs
	data              *Data
	youtubeClient     *http.Client
	youtubeOAuth      oauth2.Config
	youtubeToken      *oauth2.Token
	youtubeService    *youtube.Service
	liveChatID        string
	liveChatPageToken string

	// twitchStreamName string
	j *Core.JARVIS
}

// Initialize the Logging Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS) {

	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance

	// Empty it
	m.liveChatID = ""

	// Load Configuration
	m.loadConfig()

	m.j.Log.RegisterChannel("YouTube", "red", m.settings.Prefix)

	m.setupOutputs()
	m.setupData()

	m.authenticate()

	// m.setupEndpoints()

	m.setupCommands()

	m.Start()
}

// func (m *Module) getResponse(url string) (*http.Response, error) {
// 	req, err := http.NewRequest(http.MethodGet, url, nil)
// 	if err != nil {
// 		m.j.Log.Error("Twitch", "Unable to create request: "+url+", "+err.Error())
// 		return nil, nil
// 	}

// 	req.Header.Set("User-Agent", "JARVIS")
// 	req.Header.Set("Client-ID", m.settings.ClientID)
// 	req.Header.Set("Authorization", "OAuth "+m.twitchToken)

// 	req.Header.Set("Accept", "application/vnd.twitchtv.v5+json")

// 	return m.twitchClient.Do(req)
// }

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
	m.setupPolling()
}

// Stop Twitch Polling / IRC
func (m *Module) Stop() {
	if m != nil {
		if m.ticker != nil {
			m.ticker.Stop()
		}
	}
}
