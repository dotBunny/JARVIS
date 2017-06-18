package main

import (
	"os"

	"github.com/getlantern/systray"

	"log"

	Core "./core"

	Spotify "./modules/spotify"
	Twitch "./modules/twitch"
	Resources "./resources"
)

var (
	j *Core.JARVIS

	spotifyModule *Spotify.SpotifyModule
	twitchModule  *Twitch.TwitchModule

	quit chan os.Signal
)

func main() {
	systray.Run(onReady)

}

func onReady() {
	systray.SetIcon(Resources.Data)
	mQuit := systray.AddMenuItem("Quit", "Shutdown JARVIS")
	go func() {
		<-mQuit.ClickedCh
		Shutdown()
	}()

	go func() {
		// Create new Jarvis instance
		j = Core.HireJarvis()

		// // Initialize Twitch
		// twitchModule := new(Twitch.TwitchModule)
		// twitchModule.Initialize(j, discordModule)
		// go twitchModule.Connect()

		// Spotify
		spotifyModule := new(Spotify.SpotifyModule)
		spotifyModule.Initialize(j)

		// // Initialize WorkingOn
		// var workingOnModule Modules.WorkingOnModule
		// if config.WorkingOn.Enabled {
		// 	workingOnModule.Init(&config, &consoleModule)
		// }

		// // Initialize Modules
		// var overlayModule Modules.OverlayModule
		// overlayModule.Init(&config, &consoleModule)

		j.Log.Message("System", "Ready")
	}()
}

// Shutdown JARVIS
func Shutdown() {

	spotifyModule.Shutdown()

	j.Shutdown()
	log.Println("[SYSTEM]\tShutdown.")

	// Close application
	os.Exit(1)
}
