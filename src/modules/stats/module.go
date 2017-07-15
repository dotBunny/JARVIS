package stats

import (
	Core "../../core"
	Command "../command"
)

// Module Class
type Module struct {
	UseJIRAForWork bool
	outputs        *Outputs
	data           *Data
	settings       *Config

	stats map[string]Stat

	commandModule *Command.Module
	j             *Core.JARVIS
}

// Initialize the Stats Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS, commandModule *Command.Module) {
	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance
	m.commandModule = commandModule

	m.loadConfig()
	m.j.Log.RegisterChannel("Stats", "red", m.j.Config.GetPrefix())

	m.setupOutputs()
	m.setupData()

	m.setupEndpoints()
	m.setupCommands()
}
