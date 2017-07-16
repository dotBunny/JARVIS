package jira

import (
	Core "../../../core"
	"github.com/andygrunwald/go-jira"
)

// Data Structure
type Data struct {
	LastIssues       []jira.Issue
	LastIssuesString string
	LastNotifyText   string
	IssueType        string
}

func (m *Module) setupData() {
	m.data = new(Data)
	m.data.IssueType = "Task"
}

func (m *Module) GetData() *Data {
	return m.data
}
func (m *Module) outputLastIssues() {

	m.data.LastIssuesString = ""
	for _, issue := range m.data.LastIssues {

		m.data.LastIssuesString = m.data.LastIssuesString + issue.Key + "," + issue.Fields.Type.Name + "," + issue.Fields.Summary + "\n"
	}
	Core.SaveFile([]byte(m.data.LastIssuesString), m.outputs.IssuesPath)
}
