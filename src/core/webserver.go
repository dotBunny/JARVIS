package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// WebServerConfig Settings
type WebServerConfig struct {
	ListenPort int
}

// WebServerCore facilitates the callback/web related hosting
type WebServerCore struct {
	Settings *WebServerConfig

	j *JARVIS
}

// DefaultEndpoint to use
func (m *WebServerCore) DefaultEndpoint(w http.ResponseWriter, r *http.Request) {
	m.j.Log.Message("WebServer", "Request Received\n"+r.URL.String())
}

// Initialize the Logging Module
func (m *WebServerCore) Initialize(jarvisInstance *JARVIS) {

	// Create instance of Config Core
	m = new(WebServerCore)

	// Assign JARVIS (circle!)
	jarvisInstance.WebServer = m
	m.j = jarvisInstance

	// Register Log Channel
	m.j.Log.RegisterChannel("WebServer", "blue")

	// Create default general settings
	m.Settings = new(WebServerConfig)

	// Web Server Config
	m.Settings.ListenPort = 8080

	// Check Raw Data
	if m.j.Config.IsInitialized() {
		if !m.j.Config.IsValidKey("WebServer") {
			m.j.Log.Message("WebServer", "Unable to find \"WebServer\" config section. Using defaults.")
		} else {

			errorCheck := json.Unmarshal(*m.j.Config.GetConfigData("WebServer"), &m.Settings)
			if errorCheck != nil {
				m.j.Log.Message("Config", "Unable to properly parse WebServer Config, somethings may be wonky.")
				m.j.Log.Message("Config", "WebServer.ListenPort: "+fmt.Sprintf("%d", m.Settings.ListenPort))
			}
		}
	}

	// Register default endpoint
	m.RegisterEndpoint("/", m.DefaultEndpoint)

	// Start Server
	go http.ListenAndServe(":"+strconv.Itoa(m.Settings.ListenPort), nil)

	m.j.Log.Message("WebServer", "Initialized")
}

// RegisterEndpoint to WebServer
func (m *WebServerCore) RegisterEndpoint(endpoint string, function http.HandlerFunc) {
	http.HandleFunc(endpoint, function)
}
