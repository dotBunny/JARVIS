package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// WebServerConfig Settings
type WebServerConfig struct {
	Enabled    bool
	IPAddress  string
	ListenPort int
}

// WebServerCore facilitates the callback/web related hosting
type WebServerCore struct {
	externalIP string
	settings   *WebServerConfig

	j *JARVIS
}

// DefaultEndpoint to use
func (m *WebServerCore) defaultEndpoint(w http.ResponseWriter, r *http.Request) {
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
	m.settings = new(WebServerConfig)

	// Web Server Config
	m.settings.ListenPort = 8080

	// Check Raw Data
	if m.j.Config.IsInitialized() {
		if !m.j.Config.IsValidKey("WebServer") {
			m.j.Log.Message("WebServer", "Unable to find \"WebServer\" config section. Using defaults.")
		} else {

			errorCheck := json.Unmarshal(*m.j.Config.GetConfigData("WebServer"), &m.settings)
			if errorCheck != nil {
				m.j.Log.Message("Config", "Unable to properly parse WebServer Config, somethings may be wonky.")
				m.j.Log.Message("Config", "WebServer.IPAddress: "+m.settings.IPAddress)
				m.j.Log.Message("Config", "WebServer.ListenPort: "+fmt.Sprintf("%d", m.settings.ListenPort))
			}
		}
	}

	// Register default endpoint
	m.RegisterEndpoint("/", m.defaultEndpoint)

	// Start Server
	go http.ListenAndServe(":"+strconv.Itoa(m.settings.ListenPort), nil)

	m.j.Log.Message("WebServer", "Initialized")
}

// RegisterEndpoint to WebServer
func (m *WebServerCore) RegisterEndpoint(endpoint string, function http.HandlerFunc) {
	http.HandleFunc(endpoint, function)
}
