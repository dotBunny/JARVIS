package modules

// To get your userID
// curl -H 'Accept: application/vnd.twitchtv.v5+json' -H 'Client-ID: <CLIENT ID>' -X GET https://api.twitch.tv/kraken/users?login=<USERNAME>

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"strings"

	"bytes"

	Core "../core"
	"github.com/chosenken/twitch2go"
	"github.com/fatih/color"
	"github.com/thoj/go-ircevent"
)

const server string = "irc.chat.twitch.tv:6667"

// TwitchModule Class
type TwitchModule struct {
	LastFollower   string
	LastSubscriber string
	LastFollowers  string

	ChannelFollowers   uint
	ChannelViews       uint
	ChannelDisplayName string
	CurrentViewers     uint
	CurrentGame        string

	Ticker *time.Ticker

	latestFollowerPath     string
	latestFollowersPath    string
	latestSubscriberPath   string
	currentGamePath        string
	currentViewersPath     string
	currentDisplayNamePath string
	channelViewsPath       string
	channelFollowersPath   string

	irc    *irc.Connection
	client *twitch2go.Client
	config *Core.Config

	coloredName  string
	openBracket  string
	closeBracket string
}

// Init Module
func (m *TwitchModule) Init(config *Core.Config, console *ConsoleModule) {

	// Assing Config
	m.config = config

	// Only do this if we are going to write files
	if m.config.Twitch.Output {

		// Create our output paths
		m.latestFollowerPath = filepath.Join(m.config.General.OutputPath, "Twitch_LatestFollower.txt")
		m.latestFollowersPath = filepath.Join(m.config.General.OutputPath, "Twitch_LatestFollowers.txt")
		m.latestSubscriberPath = filepath.Join(m.config.General.OutputPath, "Twitch_LatestSubscriber.txt")
		m.currentGamePath = filepath.Join(m.config.General.OutputPath, "Twitch_CurrentGame.txt")
		m.currentViewersPath = filepath.Join(m.config.General.OutputPath, "Twitch_CurrentViewers.txt")
		m.currentDisplayNamePath = filepath.Join(m.config.General.OutputPath, "Twitch_CurrentDisplayName.txt")
		m.channelViewsPath = filepath.Join(m.config.General.OutputPath, "Twitch_ChannelViews.txt")
		m.channelFollowersPath = filepath.Join(m.config.General.OutputPath, "Twitch_ChannelFollowers.txt")

		// Check latestFollowerPath
		Core.Touch(m.latestFollowerPath)
		Core.Touch(m.latestSubscriberPath)
		Core.Touch(m.currentGamePath)
		Core.Touch(m.currentViewersPath)
		Core.Touch(m.currentDisplayNamePath)
		Core.Touch(m.channelViewsPath)
		Core.Touch(m.channelFollowersPath)
	}

	// Load Saved WorkingOn
	savedLatestFollower, err := ioutil.ReadFile(m.latestFollowerPath)
	if err == nil {
		m.LastFollower = string(savedLatestFollower)
	}

	// TODO: Need to auth with scope for subscribers to work
	// channel_commercial, channel_editor, channel_subscriptions,
	// &scope=user_read+channel_read
	client := twitch2go.NewClient(config.Twitch.ClientID)

	// Add Endpoints
	Core.AddEndpoint("/twitch/follower/last", m.endpointLastFollower)

	m.client = client

	twitchPollingFrequency, twitchPollingError := time.ParseDuration(m.config.Twitch.PollingFrequency)
	if twitchPollingError == nil {
		twitchPollingFrequency, _ = time.ParseDuration("10s")
	}
	m.Ticker = time.NewTicker(twitchPollingFrequency)

	// IRC functionality
	if m.config.Twitch.ChatEnabled {
		m.irc = irc.IRC(config.Twitch.ChatName, "jarvis")

		// Twitch IRC Settings
		m.irc.UseTLS = false
		m.irc.Password = config.Twitch.ChatToken

		m.irc.AddCallback("001", func(e *irc.Event) { m.irc.Join(m.config.Twitch.ChatChannel) })
		m.irc.AddCallback("366", func(e *irc.Event) {})

		if m.config.Twitch.ChatEcho {
			m.irc.AddCallback("PRIVMSG", m.ircMessage)
			m.irc.AddCallback("NOTICE", m.ircNotice)
		}

		err := m.irc.Connect(server)
		if err != nil {
			Core.Log("TWITCH", "ERROR", "Unable to connect to Twitch IRC Server.")
			Core.Log("TWITCH", "ERROR", err.Error())
			return
		}

		// Create cached handle replacement
		m.coloredName = color.HiMagentaString(m.config.Twitch.ChatName)
		m.openBracket = color.BlueString("<")
		m.closeBracket = color.BlueString(">")

		go m.irc.Loop()

		// Setup Console Commands
		console.AddHandler("twitch.say", "Say something in the Twitch IRC channel", m.consoleChannelSay)
		console.AddAlias("t", "twitch.say")
		console.AddHandler("twitch.update", "Force polling Twitch for updates.", m.consoleUpdate)
		console.AddHandler("twitch.whisper", "Whisper someone on Twitch's IRC server.", m.consoleWhisper)
		console.AddAlias("w", "twitch.whisper")
	}
}

func (m *TwitchModule) ircNotice(e *irc.Event) {
	Core.Log("TWITCH", "IMPORTANT", "[NOTICE] <"+e.Nick+"> "+e.Message())
}
func (m *TwitchModule) ircMessage(e *irc.Event) {

	message := e.Message()
	message = strings.Replace(message, m.config.Twitch.ChatName, m.coloredName, -1)

	if strings.HasPrefix(e.Arguments[0], "#") {
		Core.Log("TWITCH", "LOG", m.openBracket+e.Nick+m.closeBracket+" "+message)
	} else {
		Core.Log("TWITCH", "IMPORTANT", "[DM] "+m.openBracket+e.Nick+m.closeBracket+" "+message)
	}
}

// Loop awaiting ticker
func (m *TwitchModule) Loop() {
	for {
		select {
		case <-m.Ticker.C:
			m.Poll()
		}
	}
}

// Poll For Updates
func (m *TwitchModule) Poll() {
	m.pollFollowers()
	m.pollStream()
}

// Shutdown Module
func (m *TwitchModule) Shutdown() {
	if m != nil {
		if m.Ticker != nil {
			m.Ticker.Stop()
		}

		if m.irc != nil {
			m.irc.Disconnect()
		}
	}
}

func (m *TwitchModule) consoleChannelSay(input string) {
	m.irc.Privmsg(m.config.Twitch.ChatChannel, input)
	Core.Log("TWITCH", "LOG", m.openBracket+m.config.Twitch.ChatName+m.closeBracket+" "+input)
}

func (m *TwitchModule) consoleUpdate(input string) {
	m.Poll()
	Core.Log("TWITCH", "LOG", "Force Update")
}

func (m *TwitchModule) consoleWhisper(input string) {

	splitLocation := strings.Index(input, " ")
	var user string
	var message string
	if splitLocation > 0 {
		user = input[:splitLocation]
		message = input[(splitLocation + 1):len(input)]

		m.irc.Privmsg(user, message)
		Core.Log("TWITCH", "IMPORTANT", "[DM] @"+user+" "+message)
	} else {
		Core.Log("TWITCH", "LOG", "A message is required when whispering someone.")
	}

}

func (m *TwitchModule) endpointLastFollower(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(m.LastFollower))
}

func (m *TwitchModule) pollFollowers() {

	followers, error := m.client.GetChannelFollows(strconv.Itoa(m.config.Twitch.ChannelID), "", m.config.Twitch.LastFollowersCount, "DESC")
	if error != nil {
		Core.Log("TWITCH", "ERROR", error.Error())
		return
	}

	if followers.Total > 0 {

		// Handle Last Follower
		if followers.Follows[0].User.DisplayName != m.LastFollower {
			if m.config.Twitch.Output {
				Core.SaveFile([]byte(followers.Follows[0].User.DisplayName), m.latestFollowerPath)
			}

			m.LastFollower = followers.Follows[0].User.DisplayName
			Core.Log("TWITCH", "IMPORTANT", "New Follower "+followers.Follows[0].User.DisplayName)

			// Because the latest isn't the same we know the last 10 is not accurate

			var buffer bytes.Buffer
			items := len(followers.Follows)
			if items > m.config.Twitch.LastFollowersCount {
				items = m.config.Twitch.LastFollowersCount
			}
			for i := 0; i < items; i++ {
				buffer.WriteString(followers.Follows[i].User.DisplayName)
				buffer.WriteString("\n")
			}

			m.LastFollowers = buffer.String()
			if m.config.Twitch.Output {
				Core.SaveFile(buffer.Bytes(), m.latestFollowersPath)
			}
		}
	}

}

func (m *TwitchModule) pollStream() {
	stream, err := m.client.GetStreamByChannel(strconv.Itoa(m.config.Twitch.ChannelID))

	if err != nil {
		Core.Log("TWITCH", "ERROR", "Polling Stream Error - "+err.Error())
		return
	} else if stream == nil {
		Core.Log("TWITCH", "IMPORTANT", "Stream Offline")
		return
	} else if stream.Game == "" {
		Core.Log("TWITCH", "IMPORTANT", "Stream Offline")
		return
	}

	var workingString string

	if stream.Channel.Followers != m.ChannelFollowers {
		m.ChannelFollowers = stream.Channel.Followers
		if m.config.Twitch.Output {
			workingString = fmt.Sprint(m.ChannelFollowers)
			Core.SaveFile([]byte(workingString), m.channelFollowersPath)
		}
	}

	if stream.Channel.Views != m.ChannelViews {
		m.ChannelViews = stream.Channel.Views
		if m.config.Twitch.Output {
			workingString = fmt.Sprint(m.ChannelViews)
			Core.SaveFile([]byte(workingString), m.channelViewsPath)
		}
	}

	if stream.Viewers != m.CurrentViewers {
		m.CurrentViewers = stream.Viewers
		if m.config.Twitch.Output {
			workingString = fmt.Sprint(m.CurrentViewers)
			Core.SaveFile([]byte(workingString), m.currentViewersPath)
		}
	}

	if stream.Channel.DisplayName != m.ChannelDisplayName {
		m.ChannelDisplayName = stream.Channel.DisplayName
		if m.config.Twitch.Output {
			Core.SaveFile([]byte(m.ChannelDisplayName), m.currentDisplayNamePath)
		}
	}
	if stream.Game != m.CurrentGame {
		m.CurrentGame = stream.Game
		if m.config.Twitch.Output {
			Core.SaveFile([]byte(m.CurrentGame), m.currentGamePath)
		}
	}
}

// func (m *TwitchModule) pollSubscribers() {

// 	subscribers, error := m.client.GetChannelSubscribers(strconv.Itoa(m.config.Twitch.ChannelID), m.OAuth, 1, 0, "DESC")
// 	if error != nil {
// 		Core.Log("TWITCH", "ERROR", error.Error())
// 	}

// 	if subscribers.Total > 0 {
// 		if subscribers.Subscriptions[0].User.Name != m.LastSubscriber {

// 			if m.config.Twitch.Output {
// 				var buffer bytes.Buffer
// 				buffer.WriteString(subscribers.Subscriptions[0].User.Name)
// 				Core.SaveFile(buffer.Bytes(), m.latestSubscriberPath)
// 			}

// 			m.LastSubscriber = subscribers.Subscriptions[0].User.Name
// 			Core.Log("TWITCH", "IMPORTANT", "New Subscriber "+subscribers.Subscriptions[0].User.Name)
// 		}
// 	}
// }
