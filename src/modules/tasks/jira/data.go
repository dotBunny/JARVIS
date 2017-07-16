package jira

import (
	"io/ioutil"

	"strings"

	Core "../../../core"
)

// Data Structure
type Data struct {
	LastIssues       []JSONItem
	LastIssuesString string
	LastNotifyText   string
	IssueType        string
}

func (m *Module) setupData() {
	m.data = new(Data)
	m.data.IssueType = "Task"

	savedIssues, errorIssuesLoading := ioutil.ReadFile(m.outputs.IssuesPath)
	if errorIssuesLoading == nil {

		m.data.LastIssues = nil
		existingData := strings.Split(string(savedIssues), "\n")

		for _, value := range existingData {
			if len(value) > 3 {
				m.data.LastIssues = append(m.data.LastIssues, m.LoadJSONItem(value))
			}
		}

	} else {
		m.outputLastIssues()
	}
}

func (m *Module) GetData() *Data {
	return m.data
}
func (m *Module) outputLastIssues() {

	m.data.LastIssuesString = ""
	for _, issue := range m.data.LastIssues {

		m.data.LastIssuesString = m.data.LastIssuesString + issue.ID + "]||[" + issue.Type + "]||[" + issue.Description + "\n"
	}
	Core.SaveFile([]byte(m.data.LastIssuesString), m.outputs.IssuesPath)
}
