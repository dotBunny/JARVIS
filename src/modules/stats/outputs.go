package stats

import (
	"path/filepath"

	Core "../../core"
)

// Outputs Pathing
type Outputs struct {
	WorkingOnPath   string
	CoffeeCountPath string
	SavesCountPath  string
	CrashCountPath  string
}

func (m *Module) setupOutputs() {

	m.outputs = new(Outputs)

	m.outputs.WorkingOnPath = filepath.Join(m.j.Config.GetOutputPath(), "Stats_WorkingOn.txt")
	m.outputs.CoffeeCountPath = filepath.Join(m.j.Config.GetOutputPath(), "Stats_CoffeeCount.txt")
	m.outputs.SavesCountPath = filepath.Join(m.j.Config.GetOutputPath(), "Stats_SavesCount.txt")
	m.outputs.CrashCountPath = filepath.Join(m.j.Config.GetOutputPath(), "Stats_CrashCount.txt")

	// Touch Files
	Core.Touch(m.outputs.WorkingOnPath)
	Core.Touch(m.outputs.CoffeeCountPath)
	Core.Touch(m.outputs.SavesCountPath)
	Core.Touch(m.outputs.CrashCountPath)
}
