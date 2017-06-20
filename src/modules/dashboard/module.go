package dashboard

import (
	"path"

	Core "../../core"
	"github.com/skratchdot/open-golang/open"
)

// Module Class
type Module struct {
	dashboardPath string
	j             *Core.JARVIS
}

// Initialize the Dashboard Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS) {
	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance

	m.j.Log.RegisterChannel("Dashboard", "blue", m.j.WebServer.GetPrefix())
	m.dashboardPath = path.Join(m.j.WebServer.GetPagePath(), "dashboard.html")

	m.setupEndpoints()
}

// Show Dashboard
func (m *Module) Show() {
	open.Run("http://localhost:8080/dashboard")
}
