package modules

import (
	"fmt"
	"net/http"
	"path"

	"io/ioutil"

	Core "../core"
)

var (
	basePage string
)

// InitializeOverlay Module
func InitializeOverlay(config *Core.Config) {
	// Setup endpoint
	Core.AddEndpoint("/overlay", overlayRender)

	// Cache HTML
	basePath := path.Join(config.AppDir, "resources", "overlay", "index.html")
	basePageData, error := ioutil.ReadFile(basePath)
	basePage = string(basePageData)

	if error != nil {
		Core.Log("OVERLAY", "ERROR", "Unable to read base HTML page ("+basePath+") from resources folder.")
	}
	if len(basePage) <= 0 {
		Core.Log("OVERLAY", "ERROR", "No data to serve for overlay.")
	}
}

func overlayRender(w http.ResponseWriter, r *http.Request) {
	Core.Log("none", "none", basePage)
	if len(basePage) > 0 {
		fmt.Fprintf(w, basePage)
	} else {
		fmt.Fprintf(w, "No Overlay Found")
	}

	// all requests outside of the
	// folder need to look at the resources folder as its base
}
