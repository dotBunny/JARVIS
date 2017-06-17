package modules

// To get your userID
// curl -H 'Accept: application/vnd.twitchtv.v5+json' -H 'Client-ID: <CLIENT ID>' -X GET https://api.twitch.tv/kraken/users?login=<USERNAME>

import (
	"encoding/json"
	"fmt"
	// 	"net/http"
	// 	"path/filepath"
	// 	"strconv"
	// 	"time"

	// 	"strings"

	// 	"bytes"

	"time"

	Core "../core"
	"github.com/chosenken/twitch2go"
	irc "github.com/thoj/go-ircevent"
	// 	"github.com/fatih/color"
	// 	"github.com/thoj/go-ircevent"
)

const server string = "irc.chat.twitch.tv:6667"
const jarvisMessagePrefix string = "VaultBoy "

// TwitchConfig elements
type TwitchConfig struct {
	Channel            string
	ChannelID          int
	ChatSync           bool
	ChatSyncChannel    string
	ClientID           string
	ClientSecret       string
	Enabled            bool
	LastFollowersCount int
	PollingFrequency   int
	RedirectURI        string
	Token              string
	Username           string
}

// TwitchMessage is used to pass data around
type TwitchMessage struct {
	// Author  string
	// Command string
	// Content string
	// Raw     *discordgo.MessageCreate
}

// TwitchModule Class
type TwitchModule struct {
	// LastFollower   string
	// LastSubscriber string
	// LastFollowers  string

	// ChannelFollowers   uint
	// ChannelViews       uint
	// ChannelDisplayName string
	// CurrentViewers     uint
	// CurrentGame        string

	Ticker *time.Ticker

	// latestFollowerPath     string
	// latestFollowersPath    string
	// latestSubscriberPath   string
	// currentGamePath        string
	// currentViewersPath     string
	// currentDisplayNamePath string
	// channelViewsPath       string
	// channelFollowersPath   string
	authenticated bool
	irc           *irc.Connection
	client        *twitch2go.Client
	discord       *DiscordModule
	settings      *TwitchConfig
	j             *Core.JARVIS
}

// Initialize the Logging Module
func (m *TwitchModule) Initialize(jarvisInstance *Core.JARVIS, discordInstance *DiscordModule) {

	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance
	m.discord = discordInstance

	// Create default general settings
	m.settings = new(TwitchConfig)

	// TWitch Default Config
	m.settings.Channel = "#reapazor"
	m.settings.ChannelID = 21139969
	m.settings.ChatSync = true
	m.settings.ChatSyncChannel = "#twitch"
	m.settings.ClientID = "You need to set your ClientID"
	m.settings.ClientSecret = "You need to set your ClientSecret"
	m.settings.Enabled = true
	m.settings.LastFollowersCount = 10
	m.settings.PollingFrequency = 7
	m.settings.RedirectURI = "/twitch/callback"
	m.settings.Token = "You need to set your Token"
	m.settings.Username = "JARVIS"

	// Check Raw Data
	if m.j.Config.IsInitialized() {
		if !m.j.Config.IsValidKey("Twitch") {
			m.j.Log.Message("Twitch", "Unable to find \"Twitch\" config section. Using defaults.")
		} else {

			errorCheck := json.Unmarshal(*m.j.Config.GetConfigData("Twitch"), &m.settings)
			if errorCheck != nil {
				m.j.Log.Message("Config", "Unable to properly parse Twitch Config, somethings may be wonky.")

				m.j.Log.Message("Config", "Twitch.Channel: "+m.settings.Channel)
				m.j.Log.Message("Config", "Twitch.ChannelID: "+fmt.Sprintf("%d", m.settings.ChannelID))
				if m.settings.ChatSync {
					m.j.Log.Message("Config", "Twitch.ChatSync: true")
				} else {
					m.j.Log.Message("Config", "Twitch.ChatSync: false")
				}
				m.j.Log.Message("Config", "Twitch.ChatSyncChannel: "+m.settings.ChatSyncChannel)
				m.j.Log.Message("Config", "Twitch.ClientID: "+m.settings.ClientID)
				m.j.Log.Message("Config", "Twitch.ClientSecret: "+m.settings.ClientSecret)
				if m.settings.Enabled {
					m.j.Log.Message("Config", "Twitch.Enabled: true")
				} else {
					m.j.Log.Message("Config", "Twitch.Enabled: false")
				}
				m.j.Log.Message("Config", "Twitch.LastFollowersCount: "+fmt.Sprintf("%d", m.settings.LastFollowersCount))
				m.j.Log.Message("Config", "Twitch.PollingFrequency: "+fmt.Sprintf("%d", m.settings.PollingFrequency))
				m.j.Log.Message("Config", "Twitch.RedirectURI: "+m.settings.RedirectURI)
				m.j.Log.Message("Config", "Twitch.Token: "+m.settings.Token)
				m.j.Log.Message("Config", "Twitch.Username: "+m.settings.Username)
			}
		}
	}

	// HANDLE OUT PUTS
	// 	// Only do this if we are going to write files
	// 	if m.config.Twitch.Output {

	// 		// Create our output paths
	// 		m.latestFollowerPath = filepath.Join(m.config.General.OutputPath, "Twitch_LatestFollower.txt")
	// 		m.latestFollowersPath = filepath.Join(m.config.General.OutputPath, "Twitch_LatestFollowers.txt")
	// 		m.latestSubscriberPath = filepath.Join(m.config.General.OutputPath, "Twitch_LatestSubscriber.txt")
	// 		m.currentGamePath = filepath.Join(m.config.General.OutputPath, "Twitch_CurrentGame.txt")
	// 		m.currentViewersPath = filepath.Join(m.config.General.OutputPath, "Twitch_CurrentViewers.txt")
	// 		m.currentDisplayNamePath = filepath.Join(m.config.General.OutputPath, "Twitch_CurrentDisplayName.txt")
	// 		m.channelViewsPath = filepath.Join(m.config.General.OutputPath, "Twitch_ChannelViews.txt")
	// 		m.channelFollowersPath = filepath.Join(m.config.General.OutputPath, "Twitch_ChannelFollowers.txt")

	// 		// Check latestFollowerPath
	// 		Core.Touch(m.latestFollowerPath)
	// 		Core.Touch(m.latestSubscriberPath)
	// 		Core.Touch(m.currentGamePath)
	// 		Core.Touch(m.currentViewersPath)
	// 		Core.Touch(m.currentDisplayNamePath)
	// 		Core.Touch(m.channelViewsPath)
	// 		Core.Touch(m.channelFollowersPath)
	// 	}

	// 	// Add Endpoints
	// 	Core.RegisterEndpoint("/twitch/follower/last", m.endpointLastFollower)
	// 	Core.RegisterEndpoint("/twitch/viewers/current", m.endpointCurrentViewers)

	m.j.Log.RegisterChannel("Twitch", "purple")
}

// Connect to Twitch
func (m *TwitchModule) Connect() {

	if !m.IsEnabled() {
		return
	}
	// Make sure flag is toggled off
	m.authenticated = false

	// TODO: Need to auth with scope for subscribers to work
	// channel_commercial, channel_editor, channel_subscriptions,
	// &scope=user_read+channel_read
	m.client = twitch2go.NewClient(m.settings.ClientID)

	// Create Poller
	twitchPollingFrequency, _ := time.ParseDuration(fmt.Sprintf("%ds", m.settings.PollingFrequency))
	m.Ticker = time.NewTicker(twitchPollingFrequency)

	// Create IRC Objects
	m.irc = irc.IRC(m.settings.Username, "jarvis")
	m.irc.UseTLS = false
	m.irc.Password = m.settings.Token

	// Set IRC Connection Callback
	m.irc.AddCallback("001", m.handleConnected)
	m.irc.AddCallback("366", func(e *irc.Event) {})
	m.irc.AddCallback("PRIVMSG", m.handleMessage)
	m.irc.AddCallback("NOTICE", m.handleNotice)

	errorCheck := m.irc.Connect(server)
	if errorCheck != nil {
		m.j.Log.Error("Twitch", "Unable to connect ot Twitch IRC Server.c"+errorCheck.Error())
		return
	}

	// Set auth'd
	m.authenticated = true

	// Go off and do your thing IRC connection!
	go m.irc.Loop()
}

func (m *TwitchModule) handleConnected(event *irc.Event) {
	m.j.Log.Message("Twitch", "Joining channel "+m.settings.Channel)
	m.irc.Join(m.settings.Channel)
}

func (m *TwitchModule) handleMessage(event *irc.Event) {

	if m.settings.ChatSync {
		_, _ = m.discord.session.ChannelMessageSend(m.settings.ChatSyncChannel, Core.WrapNickname(event.Nick)+" "+event.Message())
	}
}
func (m *TwitchModule) handleNotice(event *irc.Event) {

	if m.settings.ChatSync {
		_, _ = m.discord.session.ChannelMessageSend(m.settings.ChatSyncChannel, "[NOTICE] "+Core.WrapNickname(event.Nick)+" "+event.Message())
	}
}

// IsEnabled for usage
func (m *TwitchModule) IsEnabled() bool {
	return m.settings.Enabled
}

// func (m *TwitchModule) ircMessage(e *irc.Event) {
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

// // Loop awaiting ticker
// func (m *TwitchModule) Loop() {
// 	for {
// 		select {
// 		case <-m.Ticker.C:
// 			m.Poll()
// 		}
// 	}
// }

// // Poll For Updates
// func (m *TwitchModule) Poll() {
// 	m.pollFollowers()
// 	m.pollStream()
// }

// // Shutdown Module
// func (m *TwitchModule) Shutdown() {
// 	if m != nil {
// 		if m.Ticker != nil {
// 			m.Ticker.Stop()
// 		}

// 		if m.irc != nil {
// 			m.irc.Disconnect()
// 		}
// 	}
// }

// func (m *TwitchModule) consoleBan(input string) {
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

// func (m *TwitchModule) consoleTimeout(input string) {
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
// func (m *TwitchModule) consoleKick(input string) {
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

// func (m *TwitchModule) consoleStats(input string) {
// 	Core.Log("TWITCH", "LOG", "Current Viewers: "+fmt.Sprint(m.CurrentViewers)+"\tFollowers: "+fmt.Sprint(m.ChannelFollowers))
// }
// func (m *TwitchModule) consoleUpdate(input string) {
// 	m.Poll()
// 	Core.Log("TWITCH", "LOG", "Force Update")
// }

// func (m *TwitchModule) consoleWhisper(input string) {

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

// func (m *TwitchModule) endpointLastFollower(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, string(m.LastFollower))
// }
// func (m *TwitchModule) endpointCurrentViewers(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, string(m.CurrentViewers))

// }

// func (m *TwitchModule) pollFollowers() {

// 	followers, error := m.client.GetChannelFollows(strconv.Itoa(m.config.Twitch.ChannelID), "", m.config.Twitch.LastFollowersCount, "DESC")
// 	if error != nil {
// 		Core.Log("TWITCH", "ERROR", error.Error())
// 		return
// 	}

// 	if followers.Total > 0 {

// 		// Handle Last Follower
// 		if followers.Follows[0].User.DisplayName != m.LastFollower {
// 			if m.config.Twitch.Output {
// 				Core.SaveFile([]byte(followers.Follows[0].User.DisplayName), m.latestFollowerPath)
// 			}

// 			m.LastFollower = followers.Follows[0].User.DisplayName
// 			Core.Log("TWITCH", "IMPORTANT", "New Follower "+followers.Follows[0].User.DisplayName)

// 			// Because the latest isn't the same we know the last 10 is not accurate

// 			var buffer bytes.Buffer
// 			items := len(followers.Follows)
// 			if items > m.config.Twitch.LastFollowersCount {
// 				items = m.config.Twitch.LastFollowersCount
// 			}
// 			for i := 0; i < items; i++ {
// 				buffer.WriteString(followers.Follows[i].User.DisplayName)
// 				buffer.WriteString("\n")
// 			}

// 			m.LastFollowers = buffer.String()
// 			if m.config.Twitch.Output {
// 				Core.SaveFile(buffer.Bytes(), m.latestFollowersPath)
// 			}
// 		}
// 	}
// 	followers = nil
// }

// func (m *TwitchModule) pollStream() {
// 	stream, err := m.client.GetStreamByChannel(strconv.Itoa(m.config.Twitch.ChannelID))

// 	if err != nil {
// 		Core.Log("TWITCH", "ERROR", "Polling Stream Error - "+err.Error())
// 		return
// 	} else if stream == nil {
// 		Core.Log("TWITCH", "IMPORTANT", "Stream Offline")
// 		return
// 	} else if stream.Game == "" {
// 		Core.Log("TWITCH", "IMPORTANT", "Stream Offline")
// 		return
// 	}

// 	var workingString string

// 	if stream.Channel.Followers != m.ChannelFollowers {
// 		m.ChannelFollowers = stream.Channel.Followers
// 		if m.config.Twitch.Output {
// 			workingString = fmt.Sprint(m.ChannelFollowers)
// 			Core.SaveFile([]byte(workingString), m.channelFollowersPath)
// 		}
// 	}

// 	if stream.Channel.Views != m.ChannelViews {
// 		m.ChannelViews = stream.Channel.Views
// 		if m.config.Twitch.Output {
// 			workingString = fmt.Sprint(m.ChannelViews)
// 			Core.SaveFile([]byte(workingString), m.channelViewsPath)
// 		}
// 	}

// 	if stream.Viewers != m.CurrentViewers {
// 		m.CurrentViewers = stream.Viewers
// 		if m.config.Twitch.Output {
// 			workingString = fmt.Sprintf("%03d", m.CurrentViewers)
// 			Core.SaveFile([]byte(workingString), m.currentViewersPath)
// 		}
// 	}

// 	if stream.Channel.DisplayName != m.ChannelDisplayName {
// 		m.ChannelDisplayName = stream.Channel.DisplayName
// 		if m.config.Twitch.Output {
// 			Core.SaveFile([]byte(m.ChannelDisplayName), m.currentDisplayNamePath)
// 		}
// 	}
// 	if stream.Game != m.CurrentGame {
// 		m.CurrentGame = stream.Game
// 		if m.config.Twitch.Output {
// 			Core.SaveFile([]byte(m.CurrentGame), m.currentGamePath)
// 		}
// 	}

// 	stream = nil
// }

// // SendMessageToChannel on Twitch
// func (m *TwitchModule) SendMessageToChannel(input string) {
// 	m.irc.Privmsg(m.config.Twitch.ChatChannel, input)
// 	Core.Log("TWITCH", "LOG", m.openBracket+m.config.Twitch.ChatName+m.closeBracket+" "+input)
// }

// // func (m *TwitchModule) pollSubscribers() {

// // 	subscribers, error := m.client.GetChannelSubscribers(strconv.Itoa(m.config.Twitch.ChannelID), m.OAuth, 1, 0, "DESC")
// // 	if error != nil {
// // 		Core.Log("TWITCH", "ERROR", error.Error())
// // 	}

// // 	if subscribers.Total > 0 {
// // 		if subscribers.Subscriptions[0].User.Name != m.LastSubscriber {

// // 			if m.config.Twitch.Output {
// // 				var buffer bytes.Buffer
// // 				buffer.WriteString(subscribers.Subscriptions[0].User.Name)
// // 				Core.SaveFile(buffer.Bytes(), m.latestSubscriberPath)
// // 			}

// // 			m.LastSubscriber = subscribers.Subscriptions[0].User.Name
// // 			Core.Log("TWITCH", "IMPORTANT", "New Subscriber "+subscribers.Subscriptions[0].User.Name)
// // 		}
// // 	}
// // }
