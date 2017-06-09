package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	Core "./core"
	Modules "./modules"
	"github.com/zmb3/spotify"
)

var (
	spotifyClient *spotify.Client
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

	Core.Log("Active")

	spotifyClient = Modules.SpotifyLogin(&config.Spotify)

	for {

		// Poll Spotify
		Modules.SpotifyPoll(spotifyClient, &config.Spotify)

		// Our polling rate
		time.Sleep(5 * time.Second)
	}
}

func shutdown() {
	Core.Log("Shutting Down")
}
