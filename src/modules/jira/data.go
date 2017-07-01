package jira

import (
	Core "../../core"
	"github.com/andygrunwald/go-jira"
)

// Data Structure
type Data struct {
	LastIssues       []jira.Issue
	LastIssuesString string
	LastNotifyText   string
	LastNotifyIcon   string
}

func (m *Module) setupData() {
	m.data = new(Data)
}

func (m *Module) outputLastIssues() {

	m.data.LastIssuesString = ""
	for _, issue := range m.data.LastIssues {
		m.data.LastIssuesString = m.data.LastIssuesString + issue.Fields.Type.Name + "," + issue.Fields.Summary + "\n"
	}
	Core.SaveFile([]byte(m.data.LastIssuesString), m.outputs.IssuesPath)
}
