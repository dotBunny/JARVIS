package command

import (
	"path"

	Core "../../core"
)

// Module Class
type Module struct {
	warningCount int
	errorCount   int
	scriptsPath  string
	j            *Core.JARVIS
}

// Initialize the Command Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS) {
	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance
	// Register Status Updator
	m.j.Status.RegisterUpdator("command", m.StatusUpdate)

	m.scriptsPath = path.Join(m.j.GetResourcePath(), "scripts")

	m.setupEndpoints()
}

func (m *Module) StatusUpdate() (int, int) {
	return m.warningCount, m.errorCount
}
