package main

import (
	"io"
	"log"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"syscall"
	"time"

	"fmt"

	Core "./core"
	Modules "./modules"
	"github.com/chosenken/twitch2go"
	"github.com/fatih/color"
	"github.com/zmb3/spotify"
)

var (
	spotifyClient *spotify.Client
	twitchClient  *twitch2go.Client
	logFile       *os.File

	spotifyTicker *time.Ticker
	twitchTicker  *time.Ticker
)

// Version Number
const Version string = "0.1.1"

func main() {

	// Start Logging
	dir, pathError := filepath.Abs(filepath.Dir(os.Args[0]))
	if pathError != nil {
		log.Fatal("Unable to determine path of application")
	}
	logPath := path.Join(dir, "jarvis.log")
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

	// Startup
	Core.Log("SYSTEM", "LOG", "Just A Rather Very Intelligent System "+color.BlueString("v"+Version))

	// Load Config
	var config = Core.ReadConfig()

	// Pathing Check
	os.MkdirAll(filepath.Dir(config.General.OutputPath), 0755)

	// Quit Routine
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		fmt.Println("")
		Core.Log("SYSTEM", "LOG", "Shutting Down ...")
		if spotifyTicker != nil {
			spotifyTicker.Stop()
		}
		if twitchTicker != nil {
			twitchTicker.Stop()
		}
		close(quit)
		os.Exit(1)
	}()
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	Core.Log("SYSTEM", "LOG", "Starting Up ...")

	// Initialize Webserver
	Core.InitializeWebServer(config.General.ServerPort)

	// Initialize Modules
	spotifyClient = Modules.InitializeSpotify(&config)
	twitchClient = Modules.InitializeTwitch(&config)

	// Create Our Tickers (for Channeling)
	spotifyPollingFrequency, spotifyPollingError := time.ParseDuration(config.Spotify.PollingFrequency)
	if spotifyPollingError != nil {
		spotifyPollingFrequency, _ = time.ParseDuration("5s")
	}
	spotifyTicker := time.NewTicker(spotifyPollingFrequency)

	twitchPollingFrequency, twitchPollingError := time.ParseDuration(config.Twitch.PollingFrequency)
	if twitchPollingError == nil {
		twitchPollingFrequency, _ = time.ParseDuration("10s")
	}
	twitchTicker := time.NewTicker(twitchPollingFrequency)

	// Get Initial Values
	Modules.PollSpotify(spotifyClient, &config)
	Modules.PollTwitch(twitchClient, &config)

	// Lets do this!
	Core.Log("SYSTEM", "LOG", "Ready")

	// Infinite Loop
	for {
		select {
		case <-spotifyTicker.C:
			Modules.PollSpotify(spotifyClient, &config)
		case <-twitchTicker.C:
			Modules.PollTwitch(twitchClient, &config)
		}
	}
}
