package command

import (
	"os/exec"
	"path"

	Core "../../core"
	Stats "../stats"
)

// Module Class
type Module struct {
	scriptsPath string
	j           *Core.JARVIS

	statsModule *Stats.Module
}

// Initialize the Command Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS, statsModule *Stats.Module) {
	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance
	m.statsModule = statsModule
	m.scriptsPath = path.Join(m.j.GetResourcePath(), "scripts")

	m.setupEndpoints()
}

func (m *Module) Wirecast(layer string, shot string) {
	commandLine := path.Join(m.scriptsPath, "Wirecast.appleScript")
	commandInstance := exec.Command(commandLine, "layer", layer, shot)
	// Execute Command
	err := commandInstance.Run()

	if err != nil {
		m.j.Log.Error("SYSTEM", err.Error())
	}
}
