package tasks

import (
	"io/ioutil"

	Core "../../core"
)

// Data Structure
type Data struct {
	WorkingOn string
}

func (m *Module) GetWorkingOn() string {
	return m.data.WorkingOn
}
func (m *Module) setupData() {

	m.data = new(Data)

	// Default
	m.data.WorkingOn = "JARVIS"

	// Load WorkingOn Text
	savedWorkingOn, errorWorkingOn := ioutil.ReadFile(m.outputs.WorkingOnPath)
	if errorWorkingOn == nil {
		m.data.WorkingOn = string(savedWorkingOn)
	} else {
		Core.SaveFile([]byte(m.data.WorkingOn), m.outputs.WorkingOnPath)
	}
}
