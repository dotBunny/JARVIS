package jira

import (
	"path/filepath"

	Core "../../core"
)

// Outputs Pathing
type Outputs struct {
	IssuesPath string
}

func (m *Module) setupOutputs() {

	m.outputs = new(Outputs)

	m.outputs.IssuesPath = filepath.Join(m.j.Config.GetOutputPath(), "JIRA_Issues.txt")

	// Touch Files
	Core.Touch(m.outputs.IssuesPath)

}
