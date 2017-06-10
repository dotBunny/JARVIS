package main

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	Core "./core"
	Modules "./modules"
	"github.com/chosenken/twitch2go"
	"github.com/zmb3/spotify"
)

var (
	spotifyClient *spotify.Client
	twitchClient  *twitch2go.Client
)

func main() {

	// Load Config
	var config = Core.ReadConfig()

	// Pathing Check
	os.MkdirAll(filepath.Dir(config.General.OutputPath), 0755)

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		shutdown()
		os.Exit(1)
	}()

	Core.Log("System", "Active")

	// Initialize Webserver
	Core.InitializeWebServer(config.General.ServerPort)

	// Initialize Modules
	spotifyClient = Modules.InitializeSpotify(&config)
	twitchClient = Modules.InitializeTwitch(&config)

	for {
		// TODO: Make these agnostic of each other and have configurable polling rates
		Modules.PollSpotify(spotifyClient, &config)
		Modules.PollTwitch(twitchClient, &config)
		time.Sleep(5 * time.Second)
	}
}

func shutdown() {
	Core.Log("System", "Shutting Down")
}
