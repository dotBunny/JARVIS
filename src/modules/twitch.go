package modules

// To get your userID
// curl -H 'Accept: application/vnd.twitchtv.v5+json' -H 'Client-ID: <CLIENT ID>' -X GET https://api.twitch.tv/kraken/users?login=<USERNAME>

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	Core "../core"
	"github.com/chosenken/twitch2go"
)

// TwitchModule Class
type TwitchModule struct {
	LastFollower   string
	LastSubscriber string
	OAuth          string

	twitchLatestFollowerPath   string
	twitchLatestSubscriberPath string

	client *twitch2go.Client
	config *Core.Config
}

// Init Module
func (m *TwitchModule) Init(config *Core.Config) {

	// Assing Config
	m.config = config

	// Only do this if we are going to write files
	if m.config.Twitch.Output {

		// Create our output paths
		m.twitchLatestFollowerPath = filepath.Join(m.config.General.OutputPath, "Twitch_LatestFollower.txt")
		m.twitchLatestSubscriberPath = filepath.Join(m.config.General.OutputPath, "Twitch_LatestSubscriber.txt")

		// Check twitchLatestFollowerPath
		if _, err := os.Stat(m.twitchLatestFollowerPath); os.IsNotExist(err) {
			ioutil.WriteFile(m.twitchLatestFollowerPath, nil, 0755)
		}

		// Check twitchLatestFollowerPath
		if _, err := os.Stat(m.twitchLatestSubscriberPath); os.IsNotExist(err) {
			ioutil.WriteFile(m.twitchLatestSubscriberPath, nil, 0755)
		}
	}

	// TODO: Need to auth with scope for subscribers to work
	// channel_commercial, channel_editor, channel_subscriptions,
	// &scope=user_read+channel_read
	client := twitch2go.NewClient(config.Twitch.ClientID)

	// Add Endpoints
	Core.AddEndpoint("/twitch/follower/last", m.lastFollowerEndpoint)

	m.client = client
}

// PollTwitch For Updates
func (m *TwitchModule) Poll() {
	m.pollFollowers()
}

func (m *TwitchModule) lastFollowerEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(m.LastFollower))
}

func (m *TwitchModule) pollFollowers() {

	followers, error := m.client.GetChannelFollows(strconv.Itoa(m.config.Twitch.ChannelID), "", 1, "DESC")
	if error != nil {
		Core.Log("TWITCH", "ERROR", error.Error())
		return
	}

	if followers.Total > 0 {
		if followers.Follows[0].User.DisplayName != m.LastFollower {

			if m.config.Twitch.Output {
				var buffer bytes.Buffer
				buffer.WriteString(followers.Follows[0].User.DisplayName)
				Core.SaveFile(buffer.Bytes(), m.twitchLatestFollowerPath)
			}

			m.LastFollower = followers.Follows[0].User.DisplayName
			Core.Log("TWITCH", "IMPORTANT", "New Follower "+followers.Follows[0].User.DisplayName)
		}
	}
}

func (m *TwitchModule) pollSubscribers() {

	subscribers, error := m.client.GetChannelSubscribers(strconv.Itoa(m.config.Twitch.ChannelID), m.OAuth, 1, 0, "DESC")
	if error != nil {
		Core.Log("TWITCH", "ERROR", error.Error())
	}

	if subscribers.Total > 0 {
		if subscribers.Subscriptions[0].User.Name != m.LastSubscriber {

			if m.config.Twitch.Output {
				var buffer bytes.Buffer
				buffer.WriteString(subscribers.Subscriptions[0].User.Name)
				Core.SaveFile(buffer.Bytes(), m.twitchLatestSubscriberPath)
			}

			m.LastSubscriber = subscribers.Subscriptions[0].User.Name
			Core.Log("TWITCH", "IMPORTANT", "New Subscriber "+subscribers.Subscriptions[0].User.Name)
		}
	}
}
