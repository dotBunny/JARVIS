package modules

// To get your userID
// curl -H 'Accept: application/vnd.twitchtv.v5+json' -H 'Client-ID: <CLIENT ID>' -X GET https://api.twitch.tv/kraken/users?login=<USERNAME>

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"bytes"

	Core "../core"
	"github.com/chosenken/twitch2go"
)

// TwitchLogin initilizas twitch configuration
func TwitchLogin(config *Core.TwitchConfig) *twitch2go.Client {

	// channel_commercial, channel_editor, channel_subscriptions,
	// &scope=user_read+channel_read

	client := twitch2go.NewClient(config.ClientID) //, "channel_commercial+channel_editor+channel_subscriptions")

	// Pathing Check
	os.MkdirAll(filepath.Dir(config.LatestFollowerPath), 0755)

	// Info Path
	if _, err := os.Stat(config.LatestFollowerPath); os.IsNotExist(err) {
		ioutil.WriteFile(config.LatestFollowerPath, nil, 0755)
	}

	return client
}

// TwitchPoll For Current Playing
func TwitchPoll(client *twitch2go.Client, config *Core.TwitchConfig) {
	var buffer bytes.Buffer

	followers, error := client.GetChannelFollows(config.ChannelID, "", 1, "DESC")
	if error != nil {
		Core.Log("Twitch", error.Error())
		return
	}

	if followers.Total > 0 {
		if followers.Follows[0].User.DisplayName != config.LastFollower {

			buffer.WriteString(followers.Follows[0].User.DisplayName)
			Core.SaveFile(buffer.Bytes(), config.LatestFollowerPath)
			config.LastFollower = buffer.String()
			buffer.Reset()

			buffer.WriteString("New Follower ")
			buffer.WriteString(followers.Follows[0].User.DisplayName)
			Core.Log("Twitch", buffer.String())
		}
	}

	// ** The logic is sound to do subscribers, but without oauth it wont work for now
	// TODO: ADD OAUTH FIX

	// subscribers, error := client.GetChannelSubscribers(config.ChannelID, config.OAuth, 1, 0, "DESC")

	// if error != nil {
	// 	Core.Log("Twitch", error.Error())
	// 	return
	// }

	// if subscribers.Total > 0 {
	// 	if subscribers.Subscriptions[0].User.Name != config.LastSubscriber {

	// 		buffer.Reset()

	// 		buffer.WriteString(subscribers.Subscriptions[0].User.Name)
	// 		Core.SaveFile(buffer.Bytes(), config.LatestSubscriberPath)
	// 		config.LastSubscriber = buffer.String()
	// 		buffer.Reset()

	// 		buffer.WriteString("New SUBSCRIBER ")
	// 		buffer.WriteString(subscribers.Subscriptions[0].User.Name)
	// 		Core.Log("Twitch", buffer.String())
	// 	}
	// }
}
