package clion

import (
	"path"

	Core "../../core"
	Stats "../stats"
)

// Module Class
type Module struct {
	buildScriptPath string
	j               *Core.JARVIS
	statsModule     *Stats.Module
}

// Initialize the Dashboard Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS, statsModule *Stats.Module) {
	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance
	m.statsModule = statsModule
	m.buildScriptPath = path.Join(m.j.GetResourcePath(), "scripts", "build-mac.appleScript")

	m.setupEndpoints()
}
