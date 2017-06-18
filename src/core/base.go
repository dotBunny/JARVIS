package core

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

// JARVIS Instance
type JARVIS struct {
	applicationPath string
	configPath      string
	startTime       time.Time

	WebServer *WebServerCore
	Config    *ConfigCore
	Discord   *DiscordCore
	Log       *LogCore
}

// Version Number
const Version string = "1.0"

// HireJarvis to work for you!
func HireJarvis() *JARVIS {

	// Create it!
	j := new(JARVIS)

	// Set starting time
	j.startTime = time.Now()

	// Set the application path, or try too
	j.SetApplicationPath(os.Args[0])

	// Lets check for a second argument, and use it as our path to the config
	if len(os.Args) >= 2 {
		j.SetConfigPath(os.Args[1])
	} else {
		j.SetConfigPath(path.Join(j.GetApplicationPath(), "jarvis.json"))
	}

	// Load Config
	j.Config.Initialize(j)

	// Initialize Logging Module
	j.Log.Initialize(j)
	j.Log.Message("System", "Version: v"+Version)

	// Initialize the WebServer
	j.WebServer.Initialize(j)

	// Initialize Discord
	j.Discord.Initialize(j)

	// Tell Discord to Connect
	j.Discord.Connect()

	// Send it back
	j.Log.Message("System", "Jarvis Hired!")
	return j
}

// GetApplicationPath returns the found application path
func (m *JARVIS) GetApplicationPath() string {
	return m.applicationPath
}

// Shutdown JARVIS instance
func (m *JARVIS) Shutdown() {

	// Stop Loggin
	m.Log.Shutdown()
}

// SetApplicationPath to the application directory
func (m *JARVIS) SetApplicationPath(application string) {

	applicationPath, pathError := filepath.Abs(filepath.Dir(application))
	if pathError != nil {
		log.Fatal("[System]\tUnable to determine path of application.")
	}
	log.Println("[System]\tStarting application in " + applicationPath)
	m.applicationPath = applicationPath
}

// SetConfigPath to the absoulte file
func (m *JARVIS) SetConfigPath(configPath string) {

	log.Println("[System]\tUsing configuration @ " + configPath)
	_, err := os.Stat(configPath)
	if err != nil {
		log.Fatal("[System]\tUnable to access config. Please correct your path, or leave it empty.")
	}
	m.configPath = configPath
}
