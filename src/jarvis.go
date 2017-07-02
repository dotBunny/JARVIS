package main

import (
	"os"
	"time"

	"github.com/getlantern/systray"

	"log"
	"math/rand"

	Core "./core"

	CLion "./modules/clion"
	Dashboard "./modules/dashboard"
	JIRA "./modules/jira"
	Overlay "./modules/overlay"
	Spotify "./modules/spotify"
	Stats "./modules/stats"
	Twitch "./modules/twitch"
	YouTube "./modules/youtube"
	Resources "./resources"
)

var (
	j *Core.JARVIS

	dashboardModule *Dashboard.Module
	spotifyModule   *Spotify.Module
	twitchModule    *Twitch.Module
	statsModule     *Stats.Module
	overlayModule   *Overlay.Module
	youtubeModule   *YouTube.Module
	jiraModule      *JIRA.Module
	clionModule     *CLion.Module

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

		// Start loading process, indicate via icon
		systray.SetIcon(Resources.TrayIconError)

		// Initialize Random
		rand.Seed(time.Now().Unix())

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

		// YouTube Module
		// STILL NOT WORKING
		// youtubeModule := new(YouTube.Module)
		// youtubeModule.Initialize(j)

		jiraModule := new(JIRA.Module)
		jiraModule.Initialize(j, statsModule)

		clionModule := new(CLion.Module)
		clionModule.Initialize(j, statsModule)
		// Overlay Module
		overlayModule := new(Overlay.Module)
		overlayModule.Initialize(j)

		// Ready to rock!
		j.Log.Message("System", "Ready")
		systray.SetIcon(Resources.TrayIconReady)

	}()
}

// Shutdown JARVIS
func Shutdown() {

	spotifyModule.Shutdown()
	twitchModule.Shutdown()
	youtubeModule.Shutdown()

	j.Shutdown()
	log.Println("[SYSTEM]\tShutdown.")

	// Close application
	os.Exit(1)
}
