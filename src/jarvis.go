package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"syscall"

	Core "./core"
	Modules "./modules"
	"github.com/fatih/color"
)

var (
	spotifyModule   *Modules.SpotifyModule
	twitchModule    *Modules.TwitchModule
	overlayModule   *Modules.OverlayModule
	consoleModule   *Modules.ConsoleModule
	workingOnModule *Modules.WorkingOnModule
	logFile         *os.File
	config          Core.Config
	quit            chan os.Signal
)

// Version Number
const Version string = "0.2.0"

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
	config := Core.ReadConfig()

	// Pathing Check
	os.MkdirAll(filepath.Dir(config.General.OutputPath), 0755)

	// Quit Routine
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		Exit("")
	}()
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	Core.Log("SYSTEM", "LOG", "Starting Up ...")

	// Initialize Console
	var consoleModule Modules.ConsoleModule
	consoleModule.Init(&config)
	consoleModule.AddHandler("quit", "Quit the application", Exit)
	consoleModule.AddAlias("exit", "quit")
	consoleModule.AddAlias("x", "quit")

	// Initialize Webserver
	Core.InitializeWebServer(config.General.ServerPort)

	// Initialize Modules
	var overlayModule Modules.OverlayModule
	overlayModule.Init(&config, &consoleModule)

	// Initialize Spotify
	var spotifyModule Modules.SpotifyModule
	if config.Spotify.Enabled {
		spotifyModule.Init(&config, &consoleModule)
		spotifyModule.Poll()
		go spotifyModule.Loop()
	}

	// Initialize Twitch
	var twitchModule Modules.TwitchModule
	if config.Twitch.Enabled {
		twitchModule.Init(&config, &consoleModule)
		twitchModule.Poll()
		go twitchModule.Loop()
	}

	// Initialize WorkingOn
	var workingOnModule Modules.WorkingOnModule
	if config.WorkingOn.Enabled {
		workingOnModule.Init(&config, &consoleModule)
	}

	// Lets do this!
	Core.Log("SYSTEM", "LOG", "Ready")

	// Activate Console
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		consoleModule.Handle(scanner.Text())
	}
}

// Exit the application
func Exit(input string) {
	fmt.Println("")
	Core.Log("SYSTEM", "LOG", "Shutting Down ...")

	spotifyModule.Shutdown()
	twitchModule.Shutdown()

	if quit != nil {
		close(quit)
	}
	os.Exit(1)
}
