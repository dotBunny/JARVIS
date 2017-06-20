package main

import (
	"os"

	"github.com/getlantern/systray"

	"log"

	Core "./core"

	Dashboard "./modules/dashboard"
	Overlay "./modules/overlay"
	Spotify "./modules/spotify"
	Stats "./modules/stats"
	Twitch "./modules/twitch"
	Resources "./resources"
)

var (
	j *Core.JARVIS

	dashboardModule *Dashboard.Module
	spotifyModule   *Spotify.Module
	twitchModule    *Twitch.Module
	statsModule     *Stats.Module
	overlayModule   *Overlay.Module

	quit chan os.Signal
)

func main() {
	systray.Run(onReady)

}

func onReady() {
	systray.SetIcon(Resources.TrayIcon)
	mDashboard := systray.AddMenuItem("Dashboard", "Show Dashboard")
	mQuit := systray.AddMenuItem("Quit", "Shutdown JARVIS")
	go func() {
		<-mQuit.ClickedCh
		Shutdown()
	}()
	go func() {
		<-mDashboard.ClickedCh
		dashboardModule.Show()
	}()

	go func() {
		// Create new Jarvis instance
		j = Core.HireJarvis()

		// Dashboard Moudle
		dashboardModule := new(Dashboard.Module)
		dashboardModule.Initialize(j)

		// Stats Module
		statsModule := new(Stats.Module)
		statsModule.Initialize(j)

		// Spotify Module
		spotifyModule := new(Spotify.Module)
		spotifyModule.Initialize(j)

		// Twitch Module
		twitchModule := new(Twitch.Module)
		twitchModule.Initialize(j)

		// Overlay Module
		overlayModule := new(Overlay.Module)
		overlayModule.Initialize(j)

		// Ready to rock!
		j.Log.Message("System", "Ready")
	}()
}

// Shutdown JARVIS
func Shutdown() {

	spotifyModule.Shutdown()
	twitchModule.Shutdown()

	j.Shutdown()
	log.Println("[SYSTEM]\tShutdown.")

	// Close application
	os.Exit(1)
}
