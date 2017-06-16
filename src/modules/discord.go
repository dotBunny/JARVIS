//https://discordapp.com/oauth2/authorize?&client_id=YOUR_CLIENT_ID_HERE&scope=bot&permissions=0

package modules

import (
	"encoding/json"
	"fmt"

	"strings"

	Core "../core"
	"github.com/bwmarrin/discordgo"
)

type DiscordFunc func(*DiscordMessage)
type DiscordMessage struct {
	Command string
	Content string

	Author string

	Raw *discordgo.MessageCreate
}

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
	botID          string
	connected      bool
	commands       map[string]DiscordFunc
	commandAliases map[string]string
	commandCache   []string
	descriptions   map[string]string

	settings *DiscordConfig
	session  *discordgo.Session
	user     *discordgo.User
	j        *Core.JARVIS
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

	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance

	// Create command index
	m.commands = make(map[string]DiscordFunc)
	m.descriptions = make(map[string]string)

	// Create default general settings
	m.settings = new(DiscordConfig)

	// Web Server Config
	m.settings.ClientID = 0
	m.settings.ClientSecret = "You must enter a ClientID/ClientSecret."
	m.settings.RedirectURI = "/discord/callback"
	m.settings.Token = "You must enter a Token."
	m.settings.Username = "JARVIS"

	// Check Raw Data
	if m.j.Config.IsInitialized() {
		if !m.j.Config.IsValidKey("Discord") {
			m.j.Log.Message("Discord", "Unable to find \"Discord\" config section. Using defaults.")
		} else {

			errorCheck := json.Unmarshal(*m.j.Config.GetConfigData("Discord"), &m.settings)
			if errorCheck != nil {
				m.j.Log.Message("Config", "Unable to properly parse Discord Config, somethings may be wonky.")

				m.j.Log.Message("Config", "Discord.ClientID: "+fmt.Sprintf("%d", m.settings.ClientID))
				m.j.Log.Message("Config", "Discord.ClientSecret: "+m.settings.ClientSecret)
				m.j.Log.Message("Config", "Discord.RedirectURI: "+m.settings.RedirectURI)
				m.j.Log.Message("Config", "Discord.Token: "+m.settings.Token)
				m.j.Log.Message("Config", "Discord.Username: "+m.settings.Username)
			}
		}
	}

	m.j.Log.RegisterChannel("Discord", "purple")
}

// IsConnected to Discord?
func (m *DiscordModule) IsConnected() bool {
	return m.connected
}

// RegisterAlias for a command
func (m *DiscordModule) RegisterAlias(alias string, command string) {
	m.commandAliases[command] = alias
}

// RegisterCommand to use with bot
func (m *DiscordModule) RegisterCommand(command string, function DiscordFunc, description string) {

	// Sanitize
	command = strings.ToLower(command)

	// Check for command
	if m.commands[command] != nil {
		m.j.Log.Warning("Discord", "Duplicate command registration for '"+command+"', ignoring latest.")
		return
	}

	// Add to command buffer and save description
	m.commands[command] = function
	m.descriptions[command] = description

	// Add to command cache for easier lookup
	m.commandCache = append(m.commandCache, command)
}

// messageHandler handles stuff
func (m *DiscordModule) messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {

	// Dont process bots own messages
	if message.Author.ID == m.botID {
		return
	}

	contentSplit := strings.Split(message.Content, ",")
	command := strings.ToLower(contentSplit[0])

	// Assign the command we theorehtically will be processing
	var targetCommand = command

	// Check Alias
	_, alias := m.commandAliases[command]
	if alias {
		targetCommand = m.commandAliases[command]
	}

	if m.commands[command] != nil {

		// Create new Discord transport message
		var newMessage DiscordMessage
		newMessage.Author = message.Author.Username
		newMessage.Command = targetCommand
		newMessage.Content = strings.TrimLeft(message.Content, command)
		newMessage.Raw = message

		// Reference command
		execCommand := m.commands[targetCommand]

		// Execute it!
		execCommand(&newMessage)
	}
}
