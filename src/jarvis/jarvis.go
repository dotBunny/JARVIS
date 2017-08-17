package main

import (
	"os"
	"time"

	"github.com/getlantern/systray"

	"log"
	"math/rand"

	Core "./core"

	Command "./modules/command"
	Spotify "./modules/spotify"
	Stats "./modules/stats"
	Tasks "./modules/tasks"
	Twitch "./modules/twitch"
	YouTube "./modules/youtube"
	Resources "./resources"
)

var (
	j *Core.JARVIS

	spotifyModule *Spotify.Module
	twitchModule  *Twitch.Module
	tasksModule   *Tasks.Module
	youtubeModule *YouTube.Module
	commandModule *Command.Module

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
		j.WebServer.OpenDashboard()
	}()

	go func() {

		// Start loading process, indicate via icon
		systray.SetIcon(Resources.TrayIconError)

		// Initialize Random
		rand.Seed(time.Now().Unix())

		// Create new Jarvis instance
		j = Core.HireJarvis()

		// Start Command Module
		commandModule := new(Command.Module)
		commandModule.Initialize(j)

		// Stats Module
		statsModule := new(Stats.Module)
		statsModule.Initialize(j, commandModule)

		// Spotify Module
		spotifyModule := new(Spotify.Module)
		spotifyModule.Initialize(j)

		// Twitch Module
		twitchModule := new(Twitch.Module)
		twitchModule.Initialize(j, commandModule)

		twitchNotifier := new(Twitch.TwitchNotifier)
		twitchNotifier.Twitch = twitchModule
		j.Notify.ConnectTwitch(twitchNotifier)

		// YouTube Module
		// STILL NOT WORKING
		// youtubeModule := new(YouTube.Module)
		// youtubeModule.Initialize(j)

		tasksModule := new(Tasks.Module)
		tasksModule.Initialize(j)

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
