package modules

// To get your userID
// curl -H 'Accept: application/vnd.twitchtv.v5+json' -H 'Client-ID: <CLIENT ID>' -X GET https://api.twitch.tv/kraken/users?login=<USERNAME>

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	Core "../core"
	"github.com/chosenken/twitch2go"
)

// TwitchData contains the live data we compare against
type TwitchData struct {
	LastFollower   string
	LastSubscriber string
}

var (
	twitchData                 TwitchData
	twitchLatestFollowerPath   string
	twitchLatestSubscriberPath string
)

// InitializeTwitch Module
func InitializeTwitch(config *Core.Config) *twitch2go.Client {

	// Create our output paths
	twitchLatestFollowerPath = filepath.Join(config.General.OutputPath, "Twitch_LatestFollower.txt")
	twitchLatestSubscriberPath = filepath.Join(config.General.OutputPath, "Twitch_LatestSubscriber.txt")

	// Check twitchLatestFollowerPath
	if _, err := os.Stat(twitchLatestFollowerPath); os.IsNotExist(err) {
		ioutil.WriteFile(twitchLatestFollowerPath, nil, 0755)
	}

	// Check twitchLatestFollowerPath
	if _, err := os.Stat(twitchLatestSubscriberPath); os.IsNotExist(err) {
		ioutil.WriteFile(twitchLatestSubscriberPath, nil, 0755)
	}

	// TODO: Need to auth with scope for subscribers to work
	// channel_commercial, channel_editor, channel_subscriptions,
	// &scope=user_read+channel_read
	client := twitch2go.NewClient(config.Twitch.ClientID)

	return client
}

// PollTwitch For Updates
func PollTwitch(client *twitch2go.Client, config *Core.Config) {
	twitchFollowers(client, config)
}

func twitchFollowers(client *twitch2go.Client, config *Core.Config) bool {

	followers, error := client.GetChannelFollows(config.Twitch.ChannelID, "", 1, "DESC")
	if error != nil {
		Core.Log("Twitch", error.Error())
		return false
	}

	if followers.Total > 0 {
		if followers.Follows[0].User.DisplayName != twitchData.LastFollower {
			var buffer bytes.Buffer
			buffer.WriteString(followers.Follows[0].User.DisplayName)
			Core.SaveFile(buffer.Bytes(), twitchLatestFollowerPath)
			twitchData.LastFollower = followers.Follows[0].User.DisplayName

			// Alert
			Core.Log("Twitch", "New Follower "+followers.Follows[0].User.DisplayName)
		}
	}

	return true
}

func twitchSubscribers(client *twitch2go.Client, config *Core.Config) bool {

	subscribers, error := client.GetChannelSubscribers(config.Twitch.ChannelID, config.Twitch.OAuth, 1, 0, "DESC")
	if error != nil {
		Core.Log("Twitch", error.Error())
		return false
	}

	if subscribers.Total > 0 {
		if subscribers.Subscriptions[0].User.Name != twitchData.LastSubscriber {

			var buffer bytes.Buffer
			buffer.WriteString(subscribers.Subscriptions[0].User.Name)
			Core.SaveFile(buffer.Bytes(), twitchLatestSubscriberPath)
			twitchData.LastSubscriber = subscribers.Subscriptions[0].User.Name

			// Alert
			Core.Log("Twitch", "New SUBSCRIBER "+subscribers.Subscriptions[0].User.Name)
		}
	}

	return true
}

// TwitchRender component
func TwitchRender() {
}
