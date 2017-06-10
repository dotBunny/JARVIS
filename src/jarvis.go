package main

import (
	"os"
	"os/signal"
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
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		shutdown()
		os.Exit(1)
	}()

	Core.Log("System", "Active")

	// Login to Spotify (Requires User Interaction)
	spotifyClient = Modules.SpotifyLogin(&config.Spotify)

	// Initialize Twitch
	twitchClient = Modules.TwitchLogin(&config.Twitch)

	for {

		// Poll
		Modules.SpotifyPoll(spotifyClient, &config.Spotify)
		Modules.TwitchPoll(twitchClient, &config.Twitch)

		// Our polling rate
		time.Sleep(5 * time.Second)
	}
}

func shutdown() {
	Core.Log("System", "Shutting Down")
}
