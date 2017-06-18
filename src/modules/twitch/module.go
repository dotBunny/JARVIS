package twitch

// twitch.connect (reconnect - reauth)

// To get your userID
// curl -H 'Accept: application/vnd.twitchtv.v5+json' -H 'Client-ID: <CLIENT ID>' -X GET https://api.twitch.tv/kraken/users?login=<USERNAME>

import (
	"encoding/json"
	"io/ioutil"
	"log"

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
	irc           *irc.Client

	discord          *Core.DiscordCore
	settings         *Config
	twitchClient     *http.Client
	twitchOAuth      oauth2.Config
	twitchToken      string
	twitchStreamName string
	j                *Core.JARVIS
}

// Initialize the Logging Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS, discordInstance *Core.DiscordCore) {

	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance
	m.discord = discordInstance

	// Load Configuration
	m.loadConfig()

	m.j.Log.RegisterChannel("Twitch", "purple", m.settings.Prefix)

	// Some cached settings
	m.twitchStreamName = strings.TrimLeft(m.settings.Channel, "#")

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

func (m *Module) getJSON(url string) map[string]*json.RawMessage {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "JARVIS")
	req.Header.Set("Client-ID", m.settings.ClientID)
	req.Header.Set("Accept", "application/vnd.twitchtv.v5+json")

	res, getErr := m.twitchClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var objmap map[string]*json.RawMessage

	json.Unmarshal(body, &objmap)

	return objmap
}

func (m *Module) getFollowers() {

	data := m.getJSON(twitchRootURL + "channels/" + m.settings.ChannelID + "/follows/?limit=1")

	var followerCount int
	err := json.Unmarshal(*data["_total"], &followerCount)
	if err != nil {
		m.j.Log.Warning("Twitch", "Failed to update follower count.")
	}
	//log.Println("Followers: " + fmt.Sprintf("%d", followerCount))

}

// IsEnabled for usage
func (m *Module) IsEnabled() bool {
	return m.settings.Enabled
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

// // Loop awaiting ticker
// func (m *Module) Loop() {
// 	for {
// 		select {
// 		case <-m.Ticker.C:
// 			m.Poll()
// 		}
// 	}
// }

// // Poll For Updates
// func (m *Module) Poll() {
// 	m.pollFollowers()
// 	m.pollStream()
// }

// // Shutdown Module
// func (m *Module) Shutdown() {
// 	if m != nil {
// 		if m.Ticker != nil {
// 			m.Ticker.Stop()
// 		}

// 		if m.irc != nil {
// 			m.irc.Disconnect()
// 		}
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
// 	Core.Log("TWITCH", "LOG", "Current Viewers: "+fmt.Sprint(m.CurrentViewers)+"\tFollowers: "+fmt.Sprint(m.ChannelFollowers))
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

// func (m *Module) endpointLastFollower(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, string(m.LastFollower))
// }
// func (m *Module) endpointCurrentViewers(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, string(m.CurrentViewers))

// }

// func (m *Module) pollFollowers() {

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

// func (m *Module) pollStream() {
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
// func (m *Module) SendMessageToChannel(input string) {
// 	m.irc.Privmsg(m.config.Twitch.ChatChannel, input)
// 	Core.Log("TWITCH", "LOG", m.openBracket+m.config.Twitch.ChatName+m.closeBracket+" "+input)
// }

// // func (m *Module) pollSubscribers() {

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
