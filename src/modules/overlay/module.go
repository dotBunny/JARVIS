package overlay

import (
	"path"

	Core "../../core"
)

// Module Class
type Module struct {
	overlayPath string
	j           *Core.JARVIS
}

// Initialize the Dashboard Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS) {
	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance
	m.j.Log.RegisterChannel("Overlay", "pink", m.j.WebServer.GetPrefix())

	m.overlayPath = path.Join(m.j.WebServer.GetPagePath(), "overlay.html")

	m.setupEndpoints()
}
