package dashboard

import (
	Core "../../core"
)

// Module Class
type Module struct {
	j *Core.JARVIS
}

// Initialize the Dashboard Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS) {
	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance

	m.j.Log.RegisterChannel("Twitch", "blue")
	m.setupEndpoints()
}
