package tasks

import (
	"path/filepath"

	Core "../../core"
)

// Outputs Pathing
type Outputs struct {
	WorkingOnPath string
}

func (m *Module) setupOutputs() {

	// TODO: CLEAN THIS UP
	m.outputs = new(Outputs)
	m.outputs.WorkingOnPath = filepath.Join(m.j.Config.GetOutputPath(), "TaskManager_WorkingOn.txt")

	// Touch Files
	Core.Touch(m.outputs.WorkingOnPath)
}
