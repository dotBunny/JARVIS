//https://discordapp.com/oauth2/authorize?&client_id=YOUR_CLIENT_ID_HERE&scope=bot&permissions=0

package modules

import (
	Core "../core"
	"github.com/bwmarrin/discordgo"
)

// DiscordConfig Settings
type DiscordConfig struct {
	ClientID     uint
	ClientSecret string
	Username     string
	Token        string
	RedirectURI  string
}

// DiscordModule facilitates the callback/web related hosting
type DiscordModule struct {
	botID     string
	connected bool
	settings  *DiscordConfig
	session   *discordgo.Session
	user      *discordgo.User
	j         *Core.JARVIS
}

// Connect to Discord Server
func (m *DiscordModule) Connect() {

	var errorCheck error

	// Reset status
	m.connected = false

	// Create a new Discord session using the provided bot token.
	m.session, errorCheck = discordgo.New("Bot " + m.settings.Token)
	if errorCheck != nil {
		m.j.Log.Warning("Discord", "Unable to create new Discord session. "+errorCheck.Error())
	}

	// Get the account information.
	errorCheck = nil
	m.user, errorCheck = m.session.User("@me")
	if errorCheck != nil {
		m.j.Log.Warning("Discord", "Unable to obtain account details, "+errorCheck.Error())
	}

	// Store/cache the account ID for later use.
	m.botID = m.user.ID

	// Register our handler for new messageCreate events
	m.session.AddHandler(m.messageHandler)

	// Open the websocket and begin listening.
	errorCheck = nil
	errorCheck = m.session.Open()
	if errorCheck != nil {
		m.j.Log.Warning("Discord", "Error opening to Discord servers, "+errorCheck.Error())
	} else {

		// We're connected
		m.connected = true
	}
}

// Initialize the Logging Module
func (m *DiscordModule) Initialize(jarvisInstance *Core.JARVIS) {

	// Create instance of Config Core
	m = new(DiscordModule)

	// Assign JARVIS (circle!)
	//jarvisInstance. = m
	m.j = jarvisInstance

	m.j.Log.RegisterChannel("Discord", "purple")
}

// IsConnected to Discord?
func (m *DiscordModule) IsConnected() bool {
	return m.connected
}

// messageHandler handles stuff
func (m *DiscordModule) messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {

}
