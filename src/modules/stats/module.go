package stats

import (
	Core "../../core"
)

// Module Class
type Module struct {
	outputs *Outputs
	data    *Data
	j       *Core.JARVIS
}

// Initialize the Stats Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS) {
	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance

	m.j.Log.RegisterChannel("Stats", "red")

	m.setupOutputs()
	m.setupData()

	m.setupEndpoints()
	m.setupCommands()
}
