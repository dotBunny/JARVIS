package google

// twitch.connect (reconnect - reauth)

// To get your userID
// curl -H 'Accept: application/vnd.twitchtv.v5+json' -H 'Client-ID: <CLIENT ID>' -X GET https://api.twitch.tv/kraken/users?login=<USERNAME>

import (
	"time"

	Core "../../core"
)

// Module Class
type Module struct {
	ticker *time.Ticker

	settings *Config
	outputs  *Outputs
	data     *Data

	j *Core.JARVIS
}

// Initialize the Logging Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS) {

	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance

	// Load Configuration
	m.loadConfig()

	m.j.Log.RegisterChannel("Google", "blue", m.settings.Prefix)

	m.setupOutputs()
	m.setupData()
}
