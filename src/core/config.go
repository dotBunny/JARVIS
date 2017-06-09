package core

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// SpotifyConfig elements
type SpotifyConfig struct {
	ClientID         string
	ClientSecret     string
	CurrentInfoPath  string
	CurrentImagePath string

	LastInfoData  []byte
	LastImageData []byte
}

// Config is an external config type
type Config struct {
	Spotify SpotifyConfig
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
