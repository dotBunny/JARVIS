package stats

import (
	Core "../../core"
)

// StatsModule Class
type StatsModule struct {
	outputs *StatsOutputs
	data    *StatsData
	j       *Core.JARVIS
}

// Initialize the Stats Module
func (m *StatsModule) Initialize(jarvisInstance *Core.JARVIS) {
	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance

	m.setupOutputs()
	m.setupData()

	m.setupEndpoints()
	m.setupCommands()
}
