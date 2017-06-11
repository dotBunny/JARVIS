package modules

import (
	"fmt"
	"net/http"

	Core "../core"
)

// InitializeOverlay Module
func InitializeOverlay(config *Core.Config) {
	Core.AddEndpoint("/overlay", overlayRender)
}

func overlayRender(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Overlay Content")
}
