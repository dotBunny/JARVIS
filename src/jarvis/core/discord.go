// USE to authorize bots: https://discordapp.com/oauth2/authorize?&client_id=YOUR_CLIENT_ID_HERE&scope=bot&permissions=0

package core

import (
	"encoding/json"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// DiscordPermissions Set
type DiscordPermissions struct {
	Moderator []string
	Admin     []string
}

// DiscordConfig Settings
type DiscordConfig struct {
	ClientID             uint
	ClientSecret         string
	RedirectURI          string
	Token                string
	Username             string
	ChatChannelID        string
	LogChannelID         string
	Prefix               string
	AnnouncementChannels []string
	Permissions          DiscordPermissions
}

func (m *DiscordCore) loadConfig() {

	m.j.Config.LoadConfig("discord.json", "Discord")

	// Create default general settings
	m.settings = new(DiscordConfig)

	// Discord Default Config
	m.settings.ClientID = 0
	m.settings.ClientSecret = "You must enter a ClientID/ClientSecret."
	m.settings.RedirectURI = "/discord/callback"
	m.settings.Token = "You must enter a Token."
	m.settings.Username = "JARVIS"
	m.settings.LogChannelID = "325784977415340043"
	m.settings.ChatChannelID = "324979047396540418"
	m.settings.Prefix = ":discord:"

	// TODO Add default annoucnemet channels

	// Check Raw Data
	if m.j.Config.IsInitialized() {
		if !m.j.Config.IsValidKey("Discord") {
			m.j.Log.Message("Discord", "Unable to find \"Discord\" config section. Using defaults.")

			// TODO Disable at ths point?
		} else {

			errorCheck := json.Unmarshal(m.j.Config.GetConfigData("Discord"), &m.settings)
			if errorCheck != nil {
				m.j.Log.Message("Config", "Unable to properly parse Discord Config, somethings may be wonky.")
				m.j.Status.ErrorCount++
			}
		}
	}
}

// DiscordFunc for IRC
type DiscordFunc func(*DiscordMessage)

// DiscordMessage is used to pass data around
type DiscordMessage struct {
	Author     string
	ChannelID  string
	Command    string
	CommandKey string
	Content    string
	Raw        *discordgo.MessageCreate
}

// DiscordCore facilitates the callback/web related hosting
type DiscordCore struct {
	botID               string
	channelCache        map[string]int
	commandAliases      map[string]string
	commandAccessLevels map[string]int
	commandKeys         map[string]string
	commandCache        []string
	commands            map[string]DiscordFunc
	connected           bool
	descriptions        map[string]string
	settings            *DiscordConfig
	session             *discordgo.Session
	user                *discordgo.User
	privateUsers        []string
	j                   *JARVIS
}

// Connect to Discord Server
func (m *DiscordCore) Connect() {

	var errorCheck error

	// Reset status
	m.connected = false

	// Create a new Discord session using the provided bot token.
	m.session, errorCheck = discordgo.New("Bot " + m.settings.Token)
	if errorCheck != nil {
		m.j.Log.Error("Discord", "Unable to create new Discord session. "+errorCheck.Error())
		m.j.Status.ErrorCount++
		return
	}

	// Get the account information.
	errorCheck = nil
	m.user, errorCheck = m.session.User("@me")
	if errorCheck != nil {
		m.j.Log.Warning("Discord", "Unable to obtain account details, "+errorCheck.Error())
		m.j.Status.WarningCount++
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
		m.j.Status.ErrorCount++
	} else {

		m.j.Log.Message("Discord", "Connected")
		// We're connected
		m.connected = true
	}
}

// GetSession of Discord
func (m *DiscordCore) GetSession() *discordgo.Session {
	return m.session
}

// GetChatChannelID for sync
func (m *DiscordCore) GetChatChannelID() string {
	return m.settings.ChatChannelID
}

// GetLogChannelID for logging
func (m *DiscordCore) GetLogChannelID() string {
	return m.settings.LogChannelID
}

// Initialize the Discord Module
func (m *DiscordCore) Initialize(jarvisInstance *JARVIS) {

	// Setup References
	m = new(DiscordCore)
	jarvisInstance.Discord = m
	m.j = jarvisInstance

	// Create command index
	m.commands = make(map[string]DiscordFunc)
	m.descriptions = make(map[string]string)
	m.commandKeys = make(map[string]string)
	m.channelCache = make(map[string]int)
	m.commandAccessLevels = make(map[string]int)

	m.loadConfig()

	m.j.Log.RegisterChannel("Discord", "purple", m.settings.Prefix)
}

// IsConnected to Discord?
func (m *DiscordCore) IsConnected() bool {
	return m.connected
}

// RegisterAlias for a command
func (m *DiscordCore) RegisterAlias(alias string, command string) {
	m.commandAliases[command] = alias
}

// Announcement sends a message to all channels flagged in settings
func (m *DiscordCore) Announcement(message string) {
	for _, element := range m.settings.AnnouncementChannels {
		m.session.ChannelMessageSend(element, message)
	}
}

// AnnoucementEmbed sends an embed message to all channels flagged in settings
func (m *DiscordCore) AnnoucementEmbed(message *discordgo.MessageEmbed) {
	for _, element := range m.settings.AnnouncementChannels {
		m.session.ChannelMessageSendEmbed(element, message)
	}
}

// RegisterCommand to use with bot
func (m *DiscordCore) RegisterCommand(command string, function DiscordFunc, description string, accessLevel int, key string) {

	// Sanitize
	command = strings.ToLower(command)

	// Check for command
	if m.commands[command] != nil {
		m.j.Log.Warning("Discord", "Duplicate command registration for '"+command+"', ignoring latest.")
		m.j.Status.WarningCount++
		return
	}

	// Add to command buffer and save description
	m.commands[command] = function
	m.descriptions[command] = description
	m.commandAccessLevels[command] = accessLevel
	m.commandKeys[command] = key

	// Add to command cache for easier lookup
	m.commandCache = append(m.commandCache, command)
}

// messageHandler handles stuff
func (m *DiscordCore) messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {

	// Dont process bots own messages
	if message.Author.ID == m.botID {
		return
	}

	contentSplit := strings.Split(message.Content, " ")
	command := strings.ToLower(contentSplit[0])

	// Assign the command we theorehtically will be processing
	var targetCommand = command

	// Check Alias
	_, alias := m.commandAliases[command]
	if alias {
		targetCommand = m.commandAliases[command]
	}

	if execCommand, ok := m.commands[targetCommand]; ok {
		// Log Channel Access Only
		var accessLevelCheck = true
		if m.commandAccessLevels[targetCommand] == CommandAccessLog && message.ChannelID != m.GetLogChannelID() {

			accessLevelCheck = false
		}

		// Check Permission level Command
		if m.commandAccessLevels[targetCommand] == CommandAccessModerator {
			accessLevelCheck = false
			for i := range m.settings.Permissions.Moderator {
				if m.settings.Permissions.Moderator[i] == message.Author.ID {
					accessLevelCheck = true
				}
			}
		}

		// Check Permission level Command
		if m.commandAccessLevels[targetCommand] == CommandAccessAdmin {
			accessLevelCheck = false
			for i := range m.settings.Permissions.Admin {
				if m.settings.Permissions.Admin[i] == message.Author.ID {
					accessLevelCheck = true
				}
			}
		}

		if !accessLevelCheck {
			return
		}

		// Create new Discord transport message
		newMessage := DiscordMessage{
			Author:     message.Author.Username,
			ChannelID:  message.ChannelID,
			Command:    targetCommand,
			CommandKey: m.commandKeys[targetCommand],
			Content:    strings.TrimLeft(strings.TrimLeft(message.Content, command), " "),
			Raw:        message}

		execCommand(&newMessage)
	}
}
