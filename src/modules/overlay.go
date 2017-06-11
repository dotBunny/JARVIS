package modules

import (
	"fmt"
	"net/http"
	"path"
	"strconv"

	"io/ioutil"

	Core "../core"
)

var (
	basePage     string
	resourceBase string
	imageBase    string
	cssBase      string
	jsBase       string
)

// InitializeOverlay Module
func InitializeOverlay(config *Core.Config) {
	// Setup endpoint
	Core.AddEndpoint("/overlay", overlayRender)
	Core.AddEndpoint("/overlay/img", overlayGetImage)
	Core.AddEndpoint("/overlay/css", overlayGetCSS)
	Core.AddEndpoint("/overlay/js", overlayGetJS)

	// Cache HTML
	basePath := path.Join(config.AppDir, "resources", "overlay", "index.html")
	resourceBase := path.Join(config.AppDir, "resources", "overlay", "content")
	imageBase = path.Join(resourceBase, "img")
	cssBase = path.Join(resourceBase, "css")
	jsBase = path.Join(resourceBase, "js")

	basePageData, error := ioutil.ReadFile(basePath)
	basePage = string(basePageData)

	if error != nil {
		Core.Log("OVERLAY", "ERROR", "Unable to read base HTML page ("+basePath+") from resources folder.")
	}
	if len(basePage) <= 0 {
		Core.Log("OVERLAY", "ERROR", "No data to serve for overlay.")
	}
}

func overlayGetImage(w http.ResponseWriter, r *http.Request) {

	// Get Image File
	filePath := r.FormValue("path")

	fileData, err := ioutil.ReadFile(path.Join(imageBase, filePath))

	if err != nil {
		Core.Log("OVERLAY", "ERROR", err.Error())
	}

	// r.FormValue("path

	last3 := filePath[len(filePath)-3:]
	switch last3 {
	case "png":
		w.Header().Set("Content-Type", "image/png")
		break
	case "gif":
		w.Header().Set("Content-Type", "image/gif")
		break
	default:
		w.Header().Set("Content-Type", "image/jpeg")
		break
	}

	w.Header().Set("Content-Length", strconv.Itoa(len(fileData)))
	if _, err := w.Write(fileData); err != nil {
		Core.Log("SPOTIFY", "ERROR", "Unable to write image")
	}
}
func overlayGetCSS(w http.ResponseWriter, r *http.Request) {

	filePath := r.FormValue("path")

	fileData, err := ioutil.ReadFile(path.Join(cssBase, filePath))

	if err != nil {
		Core.Log("OVERLAY", "ERROR", err.Error())
	}
	w.Header().Set("Content-Type", "text/css")

	w.Header().Set("Content-Length", strconv.Itoa(len(fileData)))
	if _, err := w.Write(fileData); err != nil {
		Core.Log("SPOTIFY", "ERROR", "Unable to write image")
	}

}
func overlayGetJS(w http.ResponseWriter, r *http.Request) {

	filePath := r.FormValue("path")

	fileData, err := ioutil.ReadFile(path.Join(jsBase, filePath))

	if err != nil {
		Core.Log("OVERLAY", "ERROR", err.Error())
	}
	w.Header().Set("Content-Type", "application/javascript")

	w.Header().Set("Content-Length", strconv.Itoa(len(fileData)))
	if _, err := w.Write(fileData); err != nil {
		Core.Log("SPOTIFY", "ERROR", "Unable to write image")
	}
}

func overlayRender(w http.ResponseWriter, r *http.Request) {
	if len(basePage) > 0 {
		fmt.Fprintf(w, basePage)
	} else {
		fmt.Fprintf(w, "No Overlay Found")
	}

	// all requests outside of the
	// folder need to look at the resources folder as its base
}
