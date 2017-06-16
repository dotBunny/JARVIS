package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	Core "./core"
	Modules "./modules"
)

var (
	// mode       string
	// configPath string

	j *Core.JARVIS
	// spotifyModule   *Modules.SpotifyModule
	// twitchModule    *Modules.TwitchModule
	// overlayModule   *Modules.OverlayModule
	// consoleModule   *Modules.ConsoleModule
	// workingOnModule *Modules.WorkingOnModule
	// logFile         *os.File
	// config          Core.Config

	discordModule *Modules.DiscordModule
	quit          chan os.Signal
)

func main() {

	// Create shutdown procedure
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		Shutdown()
	}()
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Create new Jarvis instance
	j = Core.HireJarvis()

	// Initialize Modules (referencing in their dependencies)

	// // Start Logging

	// // Startup

	// // Pathing Check

	// Core.Log("SYSTEM", "LOG", "Starting Up ...")

	// // Initialize Console
	// var consoleModule Modules.ConsoleModule
	// consoleModule.Init(&config)
	// consoleModule.AddHandler("/quit", "Quit the application", Exit)
	// consoleModule.AddAlias("/exit", "/quit")
	// consoleModule.AddAlias("/x", "/quit")

	// // Initialize Webserver
	// Core.InitializeWebServer(config.General.ServerPort)

	// // Initialize Modules
	// var overlayModule Modules.OverlayModule
	// overlayModule.Init(&config, &consoleModule)

	// // Initialize Spotify
	// var spotifyModule Modules.SpotifyModule
	// if config.Spotify.Enabled {
	// 	spotifyModule.Init(&config, &consoleModule)
	// 	spotifyModule.Poll()
	// 	go spotifyModule.Loop()
	// }

	// // Initialize Twitch
	// var twitchModule Modules.TwitchModule
	// if config.Twitch.Enabled {
	// 	twitchModule.Init(&config, &consoleModule, &spotifyModule)
	// 	twitchModule.Poll()
	// 	go twitchModule.Loop()
	// }

	// // Initialize WorkingOn
	// var workingOnModule Modules.WorkingOnModule
	// if config.WorkingOn.Enabled {
	// 	workingOnModule.Init(&config, &consoleModule)
	// }

	// // Start UI

	// // Lets do this!
	// Core.Log("SYSTEM", "LOG", "Ready")

	// // Activate Console
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	consoleModule.Handle(scanner.Text())
	// }
}

// Shutdown JARVIS
func Shutdown() {
	fmt.Println("")
	// Core.Log("SYSTEM", "LOG", "Shutting Down ...")

	// // Shutdown modules
	// spotifyModule.Shutdown()
	// twitchModule.Shutdown()

	// // Close log ile
	// logFile.Close()

	// Close any open channels
	if quit != nil {
		close(quit)
	}

	// Close application
	os.Exit(1)
}
