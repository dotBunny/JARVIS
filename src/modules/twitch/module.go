package twitch

// twitch.connect (reconnect - reauth)

// To get your userID
// curl -H 'Accept: application/vnd.twitchtv.v5+json' -H 'Client-ID: <CLIENT ID>' -X GET https://api.twitch.tv/kraken/users?login=<USERNAME>

import (
	"golang.org/x/oauth2"
	// 	"net/http"
	// 	"path/filepath"
	// 	"strconv"
	// 	"time"

	// 	"strings"

	// 	"bytes"
	"time"

	Core "../../core"
	irc "./irc"
	// 	"github.com/fatih/color"
	// 	"github.com/thoj/go-ircevent"

	"net/http"
	"strings"
)

// TwitchMessage is used to pass data around
type TwitchMessage struct {
	// Author  string
	// Command string
	// Content string
	// Raw     *discordgo.MessageCreate
}

// Module Class
type Module struct {
	ticker *time.Ticker

	authenticated bool
	irc           *irc.Client

	discord          *Core.DiscordCore
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

	m.setupData()
	m.setupOutputs()
	m.setupEndpoints()

	// Some cached settings
	m.twitchStreamName = strings.TrimLeft(m.settings.Channel, "#")

	m.setupPolling()
}

// Connect to Twitch
func (m *Module) Connect() {

	if !m.IsEnabled() {
		return
	}
	// Make sure flag is toggled off
	m.authenticated = false

	// Start OAuth2 Procedure
	m.authenticate()
	//go m.getFollowers()

	// Create Poller
	// twitchPollingFrequency, _ := time.ParseDuration(fmt.Sprintf("%ds", m.settings.PollingFrequency))
	// m.Ticker = time.NewTicker(twitchPollingFrequency)

	// Dont connect twich
	//m.connectIRC()
}

func (m *Module) getResponse(url string) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		m.j.Log.Error("Twitch", "Unable to create request: "+url+", "+err.Error())
		return nil, nil
	}

	req.Header.Set("User-Agent", "JARVIS")
	req.Header.Set("Client-ID", m.settings.ClientID)
	req.Header.Set("Authorization", m.twitchToken)
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
		if m.ticker != nil {
			m.ticker.Stop()
		}

		// if m.irc != nil {
		// 	m.irc.Disconnect()
		// }
	}
}

// func (m *Module) ircMessage(e *irc.Event) {
// 	message := e.Message()

// 	message = strings.Replace(message, m.config.Twitch.ChatName, m.coloredName, -1)

// 	if strings.HasPrefix(e.Arguments[0], "#") {
// 		Core.Log("TWITCH", "LOG", m.openBracket+e.Nick+m.closeBracket+" "+message)

// 		// TODO This is where we'd build out chat command recognition and all that sort of feature line
// 		if message == "!spotify" {
// 			m.SendMessageToChannel(jarvisMessagePrefix + m.spotify.GetCurrentlyPlayingMessage())
// 		}

// 	} else {
// 		Core.Log("TWITCH", "IMPORTANT", "[DM] "+m.openBracket+e.Nick+m.closeBracket+" "+message)
// 	}
// }

// func (m *Module) consoleBan(input string) {
// 	splitLocation := strings.Index(input, " ")
// 	var user string
// 	var message = "Bye Bye!"
// 	if splitLocation > 0 {
// 		user = input[:splitLocation]
// 		message = strings.Trim(input[(splitLocation+1):len(input)], " ")
// 		if len(message) <= 0 {
// 			message = "Bye Bye!"
// 		}
// 	} else {
// 		user = input
// 	}

// 	m.irc.SendRaw("CLEARCHAT " + m.config.Twitch.ChatChannel + " :" + user + " @ban-duration=;ban-reason=" + message)
// 	m.irc.SendRaw("CLEARCHAT " + m.config.Twitch.ChatChannel + " :" + user)
// 	Core.Log("TWITCH", "IMPORTANT", "Banned @"+user+" ("+message+")")
// }

// func (m *Module) consoleTimeout(input string) {
// 	splitLocation := strings.Index(input, " ")
// 	var user string
// 	var timeout = "30"
// 	if splitLocation > 0 {
// 		user = input[:splitLocation]
// 		timeout = strings.Trim(input[(splitLocation+1):len(input)], " ")
// 		if len(timeout) <= 0 {
// 			timeout = "30"
// 		}
// 	} else {
// 		user = input
// 	}

// 	m.irc.SendRaw("CLEARCHAT " + m.config.Twitch.ChatChannel + " :" + user + " @ban-duration=" + timeout + ";ban-reason=You\\'re on a break!")
// 	m.irc.SendRaw("CLEARCHAT " + m.config.Twitch.ChatChannel + " :" + user)
// 	Core.Log("TWITCH", "IMPORTANT", "Timedout @"+user+" ("+timeout+" seconds)")
// }
// func (m *Module) consoleKick(input string) {
// 	splitLocation := strings.Index(input, " ")
// 	var user string
// 	var message = "Bye Bye!"
// 	if splitLocation > 0 {
// 		user = input[:splitLocation]
// 		message = strings.Trim(input[(splitLocation+1):len(input)], " ")
// 		if len(message) <= 0 {
// 			message = "Bye Bye!"
// 		}
// 	} else {
// 		user = input
// 	}
// 	m.irc.Kick(input, m.config.Twitch.ChatChannel, message)
// 	Core.Log("TWITCH", "IMPORTANT", "Kicked @"+user+" ("+message+")")
// }

// func (m *Module) consoleStats(input string) {
// 	Core.Log("TWITCH", "LOG", "Current Viewers: "+fmt.Sprint(m.ChannelViewers)+"\tFollowers: "+fmt.Sprint(m.ChannelFollowers))
// }
// func (m *Module) consoleUpdate(input string) {
// 	m.Poll()
// 	Core.Log("TWITCH", "LOG", "Force Update")
// }

// func (m *Module) consoleWhisper(input string) {

// 	splitLocation := strings.Index(input, " ")
// 	var user string
// 	var message string
// 	if splitLocation > 0 {
// 		user = input[:splitLocation]
// 		message = input[(splitLocation + 1):len(input)]

// 		m.irc.Privmsg(user, message)
// 		Core.Log("TWITCH", "IMPORTANT", "[DM] @"+user+" "+message)
// 	} else {
// 		Core.Log("TWITCH", "LOG", "A message is required when whispering someone.")
// 	}

// }

// // SendMessageToChannel on Twitch
// func (m *Module) SendMessageToChannel(input string) {
// 	m.irc.Privmsg(m.config.Twitch.ChatChannel, input)
// 	Core.Log("TWITCH", "LOG", m.openBracket+m.config.Twitch.ChatName+m.closeBracket+" "+input)
// }
