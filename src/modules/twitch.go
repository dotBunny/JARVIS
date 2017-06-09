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

func TwitchInit(config *Core.TwitchConfig) {
	// Pathing Check
	os.MkdirAll(filepath.Dir(config.LatestFollowerPath), 0755)

	// Info Path
	if _, err := os.Stat(config.LatestFollowerPath); os.IsNotExist(err) {
		ioutil.WriteFile(config.LatestFollowerPath, nil, 0755)
	}
}

// TwitchPoll For Current Playing
func TwitchPoll(client *twitch2go.Client, config *Core.TwitchConfig) {

	followers, error := client.GetChannelFollows(config.ChannelID, "", 1, "DESC")
	if error != nil {
		Core.Log("Twitch", "Unable to query channel for followers.")
		return
	}

	if followers.Follows[0].User.DisplayName != config.LastFollower {

		var buffer bytes.Buffer

		buffer.WriteString(followers.Follows[0].User.DisplayName)
		Core.SaveFile(buffer.Bytes(), config.LatestFollowerPath)
		config.LastFollower = buffer.String()
		buffer.Reset()

		buffer.WriteString("New Follower ")
		buffer.WriteString(followers.Follows[0].User.DisplayName)
		Core.Log("Twitch", buffer.String())

	}

	// TODO: To add Subscribers need to add OAUTH to config
}
