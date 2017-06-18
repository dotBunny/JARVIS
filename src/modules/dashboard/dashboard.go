package dashboard

import (
	Core "../../core"
)

// DashboardModule Class
type DashboardModule struct {
	j *Core.JARVIS
}

// Initialize the Dashboard Module
func (m *DashboardModule) Initialize(jarvisInstance *Core.JARVIS) {
	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance

	m.setupEndpoints()
}

// import (
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"path"
// 	"strconv"

// 	"io/ioutil"

// 	Core "../core"
// )

// // OverlayModule Class
// type OverlayModule struct {
// 	baseDir      string
// 	basePage     string
// 	pageBasePath string
// 	basePath     string
// 	resourceBase string

// 	config *Core.Config
// }

// // Init  Module
// func (m *OverlayModule) Init(config *Core.Config, console *ConsoleModule) {

// 	// Assing Config
// 	m.config = config

// 	// Setup endpoint
// 	Core.RegisterEndpoint("/overlay", m.endpointOverlay)
// 	Core.RegisterEndpoint("/overlay/resource", m.endpointOverlayResource)
// 	Core.RegisterEndpoint("/overlay/page", m.endpointOverlayPage)

// 	Core.Log("OVERLAY", "IMPORTANT", "Your default overlay can be accessed at: http://localhost:"+strconv.Itoa(m.config.General.ServerPort)+"/overlay")

// 	m.baseDir = m.config.AppDir
// 	m.basePath = path.Join(m.config.AppDir, "resources", "overlay", "index.html")
// 	m.resourceBase = path.Join(m.config.AppDir, "resources", "overlay", "content")
// 	m.pageBasePath = path.Join(m.config.AppDir, "resources", "overlay")
// }

// func (m *OverlayModule) endpointOverlayPage(w http.ResponseWriter, r *http.Request) {

// 	// Build File Path
// 	filePath := path.Join(m.pageBasePath, r.URL.RawQuery)

// 	// Check Existence
// 	_, err := os.Stat(filePath)
// 	if err != nil {
// 		Core.Log("OVERLAY", "ERROR", "Unable to find file: "+filePath)
// 		fmt.Fprintf(w, "Resource Not Found")
// 		return
// 	}

// 	pageData, error := ioutil.ReadFile(filePath)
// 	if error != nil {
// 		Core.Log("OVERLAY", "ERROR", "Unable to read base HTML page ("+filePath+") from resources folder.")
// 	}

// 	if len(pageData) <= 0 {
// 		Core.Log("OVERLAY", "ERROR", "No data to serve for overlay.")
// 		fmt.Fprintf(w, "No Overlay Found")
// 	} else {
// 		fmt.Fprintf(w, string(pageData))
// 	}
// }

// func (m *OverlayModule) endpointOverlayResource(w http.ResponseWriter, r *http.Request) {

// 	// Build File Path
// 	filePath := path.Join(m.resourceBase, r.URL.RawQuery)

// 	// Check Existence
// 	_, err := os.Stat(filePath)
// 	if err != nil {
// 		Core.Log("OVERLAY", "ERROR", "Unable to find file: "+filePath)
// 		fmt.Fprintf(w, "Resource Not Found")
// 		return
// 	}

// 	fileData, err := ioutil.ReadFile(filePath)
// 	if err != nil {
// 		Core.Log("OVERLAY", "ERROR", err.Error())
// 		fmt.Fprintf(w, "Invalid Resource")
// 		return
// 	}

// 	// No need to cache locally
// 	w.Header().Set("Cache-Control", "no-cache, must-revalidate")

// 	// Check MIME Type
// 	last3 := filePath[len(filePath)-3:]
// 	switch last3 {
// 	case "png":
// 		w.Header().Set("Content-Type", "image/png")
// 		break
// 	case "gif":
// 		w.Header().Set("Content-Type", "image/gif")
// 		break
// 	case ".js":
// 		w.Header().Set("Content-Type", "application/javascript")
// 		break
// 	case "css":
// 		w.Header().Set("Content-Type", "text/css")
// 		break
// 	case "jpg":
// 	case "peg":
// 		w.Header().Set("Content-Type", "image/jpeg")
// 		break
// 	default:
// 		w.Header().Set("Content-Type", "text/plain")
// 		break
// 	}

// 	w.Header().Set("Content-Length", strconv.Itoa(len(fileData)))
// 	if _, err := w.Write(fileData); err != nil {
// 		Core.Log("SPOTIFY", "ERROR", "Unable to write resource stream")
// 		Core.Log("SPOTIFY", "ERROR", err.Error())
// 	}
// }
