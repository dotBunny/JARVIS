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
	ServerPort string
}

// SpotifyConfig elements
type SpotifyConfig struct {
	ClientID     string
	ClientSecret string
}

// TwitchConfig elements
type TwitchConfig struct {
	ClientID     string
	ClientSecret string
	OAuth        string
	ChannelID    string
}

// JIRAConfig elements
type JIRAConfig struct {
	URI string
}

// Config is an external config type
type Config struct {
	General GeneralConfig
	Spotify SpotifyConfig
	Twitch  TwitchConfig
}

// ReadConfig gets the local config file
func ReadConfig() Config {

	dir, pathError := filepath.Abs(filepath.Dir(os.Args[0]))
	if pathError != nil {
		log.Fatal("Unable to determine path of application")
	}

	configPath := path.Join(dir, "jarvis.toml")

	_, err := os.Stat(configPath)

	if err != nil {
		log.Fatal("Config file is missing: ", configPath)
	}

	var config Config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		log.Fatal(err)
	}

	return config
}
