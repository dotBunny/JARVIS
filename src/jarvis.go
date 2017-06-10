package main

import (
	"io"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"fmt"

	Core "./core"
	Modules "./modules"
	"github.com/chosenken/twitch2go"
	"github.com/zmb3/spotify"
)

var (
	spotifyClient *spotify.Client
	twitchClient  *twitch2go.Client
	logFile       *os.File
)

func main() {

	// Start Logging
	logFile, err := os.OpenFile("jarvis.log", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

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

	Core.Log("SYSTEM", "LOG", "Active")

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
	fmt.Println("")
	Core.Log("SYSTEM", "LOG", "Shutting Down")
}
