package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

// WebServerConfig Settings
type WebServerConfig struct {
	Enabled    bool
	IPAddress  string
	ListenPort int
	Prefix     string
}

// WebServerCore facilitates the callback/web related hosting
type WebServerCore struct {
	externalIP  string
	settings    *WebServerConfig
	contentPath string
	pagePath    string

	j *JARVIS
}

func (m *WebServerCore) endpointContent(w http.ResponseWriter, r *http.Request) {

	// Santize query (just incase)
	var query = strings.Replace(r.URL.RawQuery, "..", "/", -1)
	// Build File Path
	filePath := path.Join(m.contentPath, query)

	// Check Existence
	_, err := os.Stat(filePath)
	if err != nil {
		m.j.Log.Error("WebServer", "Unable to find file: "+filePath)
		fmt.Fprintf(w, "Content Not Found")
		return
	}

	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		m.j.Log.Error("WebServer", "Unable to read file: "+filePath)
		fmt.Fprintf(w, "Content Not Readable")
		return
	}

	// No need to cache locally
	w.Header().Set("Cache-Control", "no-cache, must-revalidate")

	// Check MIME Type
	last3 := filePath[len(filePath)-3:]
	switch last3 {
	case "png":
		w.Header().Set("Content-Type", "image/png")
		break
	case "gif":
		w.Header().Set("Content-Type", "image/gif")
		break
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
		break
	case "css":
		w.Header().Set("Content-Type", "text/css")
		break
	case "jpg":
	case "peg":
		w.Header().Set("Content-Type", "image/jpeg")
		break
	default:
		w.Header().Set("Content-Type", "text/plain")
		break
	}

	w.Header().Set("Content-Length", strconv.Itoa(len(fileData)))
	if _, err := w.Write(fileData); err != nil {
		m.j.Log.Error("WebServer", "Unable to  serve file: "+filePath+", "+err.Error())
	}
}

// DefaultEndpoint to use
func (m *WebServerCore) endpointDefault(w http.ResponseWriter, r *http.Request) {
	// Dashboard redirect?
	//m.j.Log.Message("WebServer", "Request Received\n"+r.URL.String())
}

// TODO PAGE SERVING
func (m *WebServerCore) endpointPage(w http.ResponseWriter, r *http.Request) {

	// Sanatize
	var query = strings.Replace(r.URL.RawQuery, "..", "/", -1)

	// Build File Path
	filePath := path.Join(m.j.WebServer.GetPagePath(), query)

	// Check Existence
	_, err := os.Stat(filePath)
	if err != nil {
		m.j.Log.Error("WebServer", "Unable to find file: "+filePath)
		fmt.Fprintf(w, "Content Not Found")
		return
	}

	pageData, err := ioutil.ReadFile(filePath)
	if err != nil {
		m.j.Log.Error("WebServer", "Unable to read file: "+filePath)
		fmt.Fprintf(w, "Content Not Readable")
		return
	}

	if len(pageData) <= 0 {
		m.j.Log.Error("WebServer", "No data to serve. Length is off.")
		fmt.Fprintf(w, "No Overlay Found")
	} else {
		w.Header().Set("Content-Length", strconv.Itoa(len(pageData)))
		fmt.Fprintf(w, string(pageData))
	}

}

// GetContentPath for pathing
func (m *WebServerCore) GetContentPath() string {
	return m.contentPath
}

// GetPagePath for pathing
func (m *WebServerCore) GetPagePath() string {
	return m.pagePath
}

// GetIPAddress server is listening on
func (m *WebServerCore) GetIPAddress() string {
	return m.settings.IPAddress
}

// GetPort server is listening on
func (m *WebServerCore) GetPort() string {
	return strconv.Itoa(m.settings.ListenPort)
}

// GetPrefix for webserver
func (m *WebServerCore) GetPrefix() string {
	return m.settings.Prefix
}

// Initialize the Logging Module
func (m *WebServerCore) Initialize(jarvisInstance *JARVIS) {

	// Create instance of Config Core
	m = new(WebServerCore)

	// Assign JARVIS (circle!)
	jarvisInstance.WebServer = m
	m.j = jarvisInstance

	// Create default general settings
	m.settings = new(WebServerConfig)

	// Register Log Channel
	m.j.Log.RegisterChannel("WebServer", "blue", m.settings.Prefix)

	// Web Server Config
	m.settings.ListenPort = 8080
	m.settings.Prefix = ":go: "

	// TODO: Get default IP

	// Check Raw Data
	if m.j.Config.IsInitialized() {
		if !m.j.Config.IsValidKey("WebServer") {
			m.j.Log.Message("WebServer", "Unable to find \"WebServer\" config section. Using defaults.")
		} else {

			errorCheck := json.Unmarshal(*m.j.Config.GetConfigData("WebServer"), &m.settings)
			if errorCheck != nil {
				m.j.Log.Message("Config", "Unable to properly parse WebServer Config, somethings may be wonky.")
			}
		}
	}

	m.pagePath = path.Join(m.j.GetApplicationPath(), "www")
	m.contentPath = path.Join(m.pagePath, "content")

	// Register default endpoint
	m.RegisterEndpoint("/", m.endpointDefault)
	m.RegisterEndpoint("/content", m.endpointContent)
	m.RegisterEndpoint("/content/", m.endpointContent)
	m.RegisterEndpoint("/page", m.endpointContent)
	m.RegisterEndpoint("/page/", m.endpointContent)

	// Start Server
	go http.ListenAndServe(":"+strconv.Itoa(m.settings.ListenPort), nil)

	m.j.Log.Message("WebServer", "Initialized")
}

// IsEnabled for usage
func (m *WebServerCore) IsEnabled() bool {
	return m.settings.Enabled
}

// RegisterEndpoint to WebServer
func (m *WebServerCore) RegisterEndpoint(endpoint string, function http.HandlerFunc) {
	http.HandleFunc(endpoint, function)
}
