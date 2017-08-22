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
	resourcePath    string
	configPath      string
	startTime       time.Time
	macBundle       bool

	CoreErrorCount  int
	CorWarningCount int

	Media     *MediaCore
	WebServer *WebServerCore
	Config    *ConfigCore
	Discord   *DiscordCore
	Log       *LogCore
	Notify    *NotifyCore
	Status    *StatusCore
}

// Version Number
const Version string = "0.9.1"

// HireJarvis to work for you!
func HireJarvis() *JARVIS {

	// Create it!
	j := new(JARVIS)

	// Set starting time
	j.startTime = time.Now()

	// Set the application path, or try too
	j.SetApplicationPath(os.Args[0])

	// Config FOlder
	_, besideConfig := os.Stat(path.Join(j.GetApplicationPath(), "config"))
	if besideConfig == nil {
		j.resourcePath = j.GetApplicationPath()
	} else {
		_, macBundleCheck := os.Stat(path.Join(j.GetApplicationPath(), "..", "Resources", "config", "general.json"))
		if macBundleCheck == nil {
			log.Println("[System]\tMac Bundle Detected")
			j.macBundle = true
			j.resourcePath = path.Join(j.GetApplicationPath(), "..", "Resources")
		} else {
			j.resourcePath = j.GetApplicationPath()
		}
	}

	// Set Config Path
	j.SetConfigPath(path.Join(j.GetResourcePath(), "config"))

	// Load Config
	j.Config.Initialize(j)

	// Initialize Logging Module
	j.Log.Initialize(j)
	j.Log.Message("System", "Version: v"+Version)

	// Setup Status Tracker
	j.Status.Initialize(j)

	// Start notification system system
	j.Notify.Initialize(j)

	// Start Media
	j.Media.Initialize(j)

	// Initialize the WebServer
	j.WebServer.Initialize(j)

	// Initialize Discord
	j.Discord.Initialize(j)

	// Tell Discord to Connect
	j.Discord.Connect()

	// Send it back
	if j.macBundle {
		j.Log.Message("System", "Jarvis Hired! (Mac Bundle)")
	} else {
		j.Log.Message("System", "Jarvis Hired!")
	}

	return j
}

// GetApplicationPath returns the found application path
func (m *JARVIS) GetApplicationPath() string {
	return m.applicationPath
}

// GetResourcePath returns the resources path
func (m *JARVIS) GetResourcePath() string {
	return m.resourcePath
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

	log.Println("[System]\tUsing configuration files @ " + configPath)
	_, err := os.Stat(configPath)
	if err != nil {
		log.Fatal("[System]\tUnable to access config files. Please correct your path, or leave it empty.")
	}
	m.configPath = configPath
}
