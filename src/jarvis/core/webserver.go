package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
)

// DiscordFunc for IRC
type WebServerParser func(string, string, *http.Request) string

type DashboardLink struct {
	Name   string
	URL    string
	Icon   string
	Target string
}

// WebServerConfig Settings
type WebServerConfig struct {
	Enabled        bool
	IPAddress      string
	ListenPort     int
	Prefix         string
	DashboardLinks []DashboardLink
}

// WebServerCore facilitates the callback/web related hosting
type WebServerCore struct {
	settings *WebServerConfig
	webPath  string
	parsers  map[string]WebServerParser
	proxies  map[string]string

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

	if m.j.Config.GetProxyStatus() {
		// If we are the local host (not proxy), then provide localhost as the IP address, otherwise use set
		return m.settings.IPAddress
	}
	return "localhost"
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
	m.proxies = make(map[string]string)
	m.RegisterParser("webserver", m.ParseWebContent)

	// Register Log Channel
	m.j.Log.RegisterChannel("WebServer", "blue", m.settings.Prefix)

	m.j.Config.LoadConfig("webserver.json", "WebServer")

	// Web Server Config
	m.settings.ListenPort = 8080
	m.settings.Prefix = ":go:"

	// TODO: Get default IP

	// Check Raw Data
	if m.j.Config.IsInitialized() {
		if !m.j.Config.IsValidKey("WebServer") {
			m.j.Log.Message("WebServer", "Unable to find \"WebServer\" config section. Using defaults.")
		} else {

			errorCheck := json.Unmarshal(m.j.Config.GetConfigData("WebServer"), &m.settings)
			if errorCheck != nil {
				m.j.Log.Message("Config", "Unable to properly parse WebServer Config, somethings may be wonky.")
				m.j.Status.ErrorCount++
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

	m.RegisterEndpoint("/proxy/register", m.endpointRegisterProxy)
	m.RegisterEndpoint("/proxy/unregister", m.endpointUnregisterProxy)
	m.RegisterEndpoint("/proxy/register/", m.endpointRegisterProxy)
	m.RegisterEndpoint("/proxy/unregister/", m.endpointUnregisterProxy)

	// Create Quick Links From Tray
	if len(m.settings.DashboardLinks) > 0 {

		var maximumLength = 0
		for _, link := range m.settings.DashboardLinks {
			if len(link.Name) > maximumLength {
				maximumLength = len(link.Name)
			}
		}

		// Create spacer
		spacer := systray.AddMenuItem(times("-", maximumLength+3), "")
		spacer.Disable()
	}

	for _, link := range m.settings.DashboardLinks {
		newLink := systray.AddMenuItem(link.Name, link.URL)
		go func() {
			for {
				select {
				case <-newLink.ClickedCh:
					open.Run(link.URL)
				}
			}
		}()
	}

	// Start Server
	go http.ListenAndServe(":"+strconv.Itoa(m.settings.ListenPort), nil)

	m.j.Log.Message("WebServer", "Initialized")
}

// IsEnabled for usage
func (m *WebServerCore) IsEnabled() bool {
	return m.settings.Enabled
}

// ParseContent data as string and replace in variables
func (m *WebServerCore) ParseContent(originalData []byte, mode string, r *http.Request) []byte {

	workingContent := string(originalData[:len(originalData)])

	for _, parser := range m.parsers {
		workingContent = parser(workingContent, mode, r)
	}

	return []byte(workingContent)
}

// RegisterEndpoint to WebServer
func (m *WebServerCore) RegisterEndpoint(endpoint string, function http.HandlerFunc) {
	http.HandleFunc(endpoint, function)
}

// TouchEndpoint of our API without returning anyhting
func (m *WebServerCore) TouchEndpoint(endpoint string) {
	go http.Get("http://" + m.GetIPAddress() + ":" + strconv.Itoa(m.settings.ListenPort) + endpoint)

	if len(m.proxies[endpoint]) > 0 {
		go http.Get(m.proxies[endpoint])
	}

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
		m.j.Status.ErrorCount++
		return
	}

	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		m.j.Log.Warning("WebServer", "Unable to read file: "+filePath)
		fmt.Fprintf(w, "Content Not Readable")
		m.j.Status.ErrorCount++
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
		w.Header().Set("Content-Type", "application/javascript")
		//parsableContent = true
		break
	case ".css":
		w.Header().Set("Content-Type", "text/css")
		break
	case ".jpg":
		w.Header().Set("Content-Type", "image/jpeg")
		break
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
		w.Header().Set("Content-Type", "text/html")
		parsableContent = true
	case ".htm":
		w.Header().Set("Content-Type", "text/html")
		parsableContent = true
		break
	case ".eot":
		w.Header().Set("Content-Type", "application/vnd.ms-fontobject")
		break
	case ".otf":
		w.Header().Set("Content-Type", "application/font-sfnt")
		break
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
		break
	}

	if parsableContent {
		fileData = m.ParseContent(fileData, ext, r)
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

func (m *WebServerCore) endpointRegisterProxy(w http.ResponseWriter, r *http.Request) {

	// ID is the callback / value is where to hit
	var id = r.FormValue("id")
	var address = r.FormValue("callback")

	if len(m.proxies[id]) > 0 {
		m.j.Log.Warning("WEB", "Duplicate proxy registration for '"+id+"', ignoring latest.")
		m.j.Status.WarningCount++
		return
	}
	m.proxies[id] = address
}
func (m *WebServerCore) endpointUnregisterProxy(w http.ResponseWriter, r *http.Request) {
	var id = r.FormValue("id")
	if len(m.proxies[id]) > 0 {
		delete(m.proxies, id)
	}
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
		m.j.Status.WarningCount++
		return
	}
	m.parsers[key] = function
}

// GetBaseURI returns the complete server web address
func (m *WebServerCore) GetBaseURI() string {
	return "http://" + m.GetIPAddress() + ":" + strconv.Itoa(m.settings.ListenPort)
}

func (m *WebServerCore) GetLocalURI() string {
	return "http://localhost:" + strconv.Itoa(m.settings.ListenPort)
}

func (m *WebServerCore) ParseWebContent(content string, mode string, r *http.Request) string {

	var buffer bytes.Buffer

	if len(m.settings.DashboardLinks) > 0 {
		buffer.WriteString("<li class=\"nav-spacer\"></li>")
	}
	for _, link := range m.settings.DashboardLinks {
		buffer.WriteString("<li><a href=\"" + link.URL + "\" target=\"" + link.Target + "\"><i class=\"fa fa-fw " + link.Icon + "\"></i> " + link.Name + "</a></li>")
	}

	// Quick Links
	content = strings.Replace(content, "[[JARVIS.WebServer.QuickLinks]]", buffer.String(), -1)

	if m.IsLocal(r) {
		return strings.Replace(content, "[[JARVIS.address]]", m.GetLocalURI(), -1)
	}
	return strings.Replace(content, "[[JARVIS.address]]", m.GetBaseURI(), -1)
}

// IsLocal Call
func (m *WebServerCore) IsLocal(r *http.Request) bool {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	if ip == "::1" {
		return true
	}
	return false
}
