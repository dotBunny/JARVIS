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

	"github.com/skratchdot/open-golang/open"
)

// DiscordFunc for IRC
type WebServerParser func(string, string) string

// WebServerConfig Settings
type WebServerConfig struct {
	Enabled    bool
	IPAddress  string
	ListenPort int
	Prefix     string
}

// WebServerCore facilitates the callback/web related hosting
type WebServerCore struct {
	settings *WebServerConfig
	webPath  string
	parsers  map[string]WebServerParser

	j *JARVIS
}

// DefaultHeader to be used
func (m *WebServerCore) DefaultHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
}

// GetPagePath for pathing
func (m *WebServerCore) GetPagePath() string {
	return m.webPath
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
	m.parsers = make(map[string]WebServerParser)

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

	m.webPath = path.Join(m.j.GetResourcePath(), "www")

	// Register default endpoint
	m.RegisterEndpoint("/", m.endpointBase)

	m.RegisterEndpoint("/media", m.endpointMedia)
	m.RegisterEndpoint("/media/", m.endpointMedia)
	m.RegisterEndpoint("/media/monitor", m.endpointMediaMonitor)
	m.RegisterEndpoint("/media/monitor/", m.endpointMediaMonitor)
	m.RegisterEndpoint("/media/fetch/", m.endpointMediaFetch)
	m.RegisterEndpoint("/media/fetch", m.endpointMediaFetch)

	// Start Server
	go http.ListenAndServe(":"+strconv.Itoa(m.settings.ListenPort), nil)

	m.j.Log.Message("WebServer", "Initialized")
}

// IsEnabled for usage
func (m *WebServerCore) IsEnabled() bool {
	return m.settings.Enabled
}

// ParseContent data as string and replace in variables
func (m *WebServerCore) ParseContent(originalData []byte, mode string) []byte {

	workingContent := string(originalData[:len(originalData)])

	for _, parser := range m.parsers {
		workingContent = parser(workingContent, mode)
	}

	return []byte(workingContent)
}

// RegisterEndpoint to WebServer
func (m *WebServerCore) RegisterEndpoint(endpoint string, function http.HandlerFunc) {
	http.HandleFunc(endpoint, function)
}

// TouchEndpoint of our API without returning anyhting
func (m *WebServerCore) TouchEndpoint(endpoint string) {
	go http.Get("http://" + m.settings.IPAddress + ":" + strconv.Itoa(m.settings.ListenPort) + endpoint)
}

func (m *WebServerCore) endpointBase(w http.ResponseWriter, r *http.Request) {

	// Santize query (just incase)
	var query = strings.Replace(r.URL.RequestURI(), "..", "/", -1)

	// Remove query
	var queryIndex = strings.Index(query, "?")
	if queryIndex > 0 {
		query = query[:queryIndex]
	}

	// Build File Path (safely)
	filePath := path.Join(m.webPath, query)

	// Check Existence
	_, err := os.Stat(filePath)
	if err != nil {
		m.j.Log.Warning("WebServer", "Unable to find file: "+filePath)
		fmt.Fprintf(w, "Content Not Found")
		return
	}

	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		m.j.Log.Warning("WebServer", "Unable to read file: "+filePath)
		fmt.Fprintf(w, "Content Not Readable")
		return
	}

	// No need to cache locally
	w.Header().Set("Cache-Control", "no-cache, must-revalidate")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Check MIME Type
	ext := filePath[strings.LastIndex(filePath, "."):]

	// Flag to see if the content is parsable (we look at the content type for this one)
	var parsableContent bool

	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		break
	case ".gif":
		w.Header().Set("Content-Type", "image/gif")
		break
	case ".js":
		// TODO: Maybe make this parse through? would be sick
		w.Header().Set("Content-Type", "application/javascript")
		break
	case ".css":
		w.Header().Set("Content-Type", "text/css")
		break
	case ".jpg":
	case ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
		break
	case ".json":
		w.Header().Set("Content-Type", "application/json")
		parsableContent = true
		break
	case ".svg":
		w.Header().Set("Content-Type", "image/svg+xml")
		break
	case ".html":
	case ".htm":
		w.Header().Set("Content-Type", "text/html")
		parsableContent = true
		break
	case ".eot":
		w.Header().Set("Content-Type", "application/vnd.ms-fontobject")
		break
	case ".otf":
	case ".ttf":
		w.Header().Set("Content-Type", "application/font-sfnt")
		break
	case ".xml":
		w.Header().Set("Content-Type", "text/xml")
		parsableContent = true
		break
	case ".woff":
		w.Header().Set("Content-Type", "application/font-woff")
		break
	case ".woff2":
		w.Header().Set("Content-Type", "font/woff2")
		break
	default:
		w.Header().Set("Content-Type", "text/plain")
		parsableContent = true
		break
	}

	if parsableContent {
		fileData = m.ParseContent(fileData, ext)
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(fileData)))

	if _, err := w.Write(fileData); err != nil {
		m.j.Log.Warning("WebServer", "Unable to  serve file: "+filePath+", "+err.Error())
	}

}

// Media player
func (m *WebServerCore) endpointMedia(w http.ResponseWriter, r *http.Request) {
	var filePath = r.FormValue("path")
	m.j.Log.Message("WebServer", "Playing media:"+filePath)
	m.j.Media.PlaySound(filePath)
}

func (m *WebServerCore) endpointMediaMonitor(w http.ResponseWriter, r *http.Request) {
	m.DefaultHeader(w)
	output := strconv.Itoa(m.j.Media.MediaLastVersion) + "," + m.GetBaseURI() + "/media/fetch"
	w.Header().Set("Content-Length", strconv.Itoa(len(output)))
	fmt.Fprintf(w, output)
}

func (m *WebServerCore) endpointMediaFetch(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "audio/wav")
	w.Write(m.j.Media.MediaLastData)
}

// OpenDashboard
func (m *WebServerCore) OpenDashboard() {
	open.Run(m.GetBaseURI() + "/dashboard.html")
}

func (m *WebServerCore) RegisterParser(key string, function WebServerParser) {

	key = strings.ToLower(key)

	// Check for command
	if m.parsers[key] != nil {
		m.j.Log.Warning("WEB", "Duplicate parser registration for '"+key+"', ignoring latest.")
		return
	}
	m.parsers[key] = function
}

// GetBaseURI returns the complete server web address
func (m *WebServerCore) GetBaseURI() string {
	return "http://" + m.settings.IPAddress + ":" + strconv.Itoa(m.settings.ListenPort)
}
