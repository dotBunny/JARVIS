package jira

import (
	"path/filepath"

	Core "../../../core"
	"github.com/andygrunwald/go-jira"
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

type JSONItem struct {
	ID          string
	Description string
	Type        string
}

func (m *Module) GetLastIssues() []JSONItem {
	return m.data.LastIssues
}

func (m *Module) GetJSONItem(issue jira.Issue) JSONItem {

	item := new(JSONItem)

	item.ID = issue.Key
	item.Description = issue.Fields.Summary
	item.Type = issue.Fields.Type.Name

	return *item
}
