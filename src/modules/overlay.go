package modules

import (
	"fmt"
	"net/http"
	"strconv"

	Core "../core"
)

var (
	refreshHeader string
)

// InitializeOverlay Module
func InitializeOverlay(config *Core.Config) {

	// Cache header
	refreshHeader = config.Overlay.RefreshFrequency + "; url=" + "http://localhost:" + strconv.Itoa(config.General.ServerPort) + config.Overlay.Endpoint

	// Setup endpoint
	Core.AddEndpoint(config.Overlay.Endpoint, overlayRender)
}

func overlayRender(w http.ResponseWriter, r *http.Request) {

	// Use Normal Refresh
	w.Header().Add("Refresh", refreshHeader)

	fmt.Fprintf(w, "Overlay Content")

}
