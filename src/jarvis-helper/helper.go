package main

import (
	"os"

	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"

	Resources "./resources"
)

var (
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
		open.Run("http://192.168.1.250:8080/dashboard.html")
	}()

	go func() {

		// Start loading process, indicate via icon
		systray.SetIcon(Resources.TrayIconReady)

	}()
}

// Shutdown JARVIS
func Shutdown() {

	// Close application
	os.Exit(1)
}
