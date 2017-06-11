package core

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

//GeneralConfig elements
type GeneralConfig struct {
	OutputPath string
	ServerPort int
}

// SpotifyConfig elements
type SpotifyConfig struct {
	Enabled bool
	Output  bool

	PollingFrequency    string
	ClientID            string
	ClientSecret        string
	Callback            string
	TruncateTrackLength int
	TruncateTrackRunes  string
}

// TwitchConfig elements
type TwitchConfig struct {
	Enabled bool
	Output  bool

	PollingFrequency string
	ClientID         string
	ClientSecret     string
	OAuth            string
	ChannelID        int
	Callback         string
}

type OverlayConfig struct {
	CacheIndex bool
}

// JIRAConfig elements
type JIRAConfig struct {
	URI string
}

// Config is an external config type
type Config struct {
	AppDir  string
	General GeneralConfig
	Spotify SpotifyConfig
	Twitch  TwitchConfig
	Overlay OverlayConfig
}

// ReadConfig gets the local config file
func ReadConfig() Config {

	dir, pathError := filepath.Abs(filepath.Dir(os.Args[0]))
	if pathError != nil {
		Log("SYSTEM", "ERROR", "Odd, the application was not able to figure out its own path. No idea. You got any?")
		log.Fatal("Unable to determine path of application")
	}

	configPath := path.Join(dir, "jarvis.toml")

	_, err := os.Stat(configPath)

	if err != nil {
		log.Fatal("Config file is missing: ", configPath)
	}

	var config Config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		Log("SYSTEM", "ERROR", "Your jarvis.toml file seems BAD! You need to edit it OR fix what you messed up.")
		log.Fatal(err)
	}

	config.AppDir = dir

	return config
}
