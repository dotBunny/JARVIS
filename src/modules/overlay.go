package modules

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"

	"io/ioutil"

	Core "../core"
)

var (
	baseDir      string
	basePage     string
	basePath     string
	resourceBase string
)

// InitializeOverlay Module
func InitializeOverlay(config *Core.Config) {
	// Setup endpoint
	Core.AddEndpoint("/overlay", overlayRender)
	Core.AddEndpoint("/overlay/resource", overlayGetResource)

	baseDir = config.AppDir
	basePath = path.Join(config.AppDir, "resources", "overlay", "index.html")
	resourceBase = path.Join(config.AppDir, "resources", "overlay", "content")
}

func overlayGetResource(w http.ResponseWriter, r *http.Request) {

	// Build File Path
	filePath := path.Join(resourceBase, r.URL.RawQuery)

	// Check Existence
	_, err := os.Stat(filePath)
	if err != nil {
		Core.Log("OVERLAY", "ERROR", "Unable to find file: "+filePath)
		fmt.Fprintf(w, "Resource Not Found")
		return
	}

	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		Core.Log("OVERLAY", "ERROR", err.Error())
		fmt.Fprintf(w, "Invalid Resource")
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
		Core.Log("SPOTIFY", "ERROR", "Unable to write resource stream")
		Core.Log("SPOTIFY", "ERROR", err.Error())
	}
}

func overlayRender(w http.ResponseWriter, r *http.Request) {

	// Server Page Per Time
	basePageData, error := ioutil.ReadFile(basePath)
	if error != nil {
		Core.Log("OVERLAY", "ERROR", "Unable to read base HTML page ("+basePath+") from resources folder.")
	} else {
		basePage = string(basePageData)
	}

	if len(basePage) <= 0 {
		Core.Log("OVERLAY", "ERROR", "No data to serve for overlay.")
		fmt.Fprintf(w, "No Overlay Found")
	} else {
		fmt.Fprintf(w, basePage)
	}
}
