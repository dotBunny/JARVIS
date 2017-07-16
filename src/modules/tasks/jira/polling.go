package jira

import (
	"strconv"
	"time"

	jira "github.com/andygrunwald/go-jira"
)

func (m *Module) setupPolling() {
	// Create Ticker
	jiraPollingFrequency, jiraPollingError := time.ParseDuration(strconv.Itoa(m.settings.PollingFrequency) + "s")
	if jiraPollingError != nil {
		jiraPollingFrequency, _ = time.ParseDuration("5s")
	}
	m.j.Log.Message("JIRA", "Starting polling at "+jiraPollingFrequency.String())
	m.ticker = time.NewTicker(jiraPollingFrequency)
	m.Poll(false)
	m.Polling = true
	go m.loop()
}

// Loop awaiting ticker
func (m *Module) loop() {
	for {
		select {
		case <-m.ticker.C:
			m.Poll(true)
		}
	}
}

// Poll For Updates
func (m *Module) Poll(notify bool) {

	m.pollIssues(notify)
}

func (m *Module) pollIssues(notify bool) {

	opt := &jira.SearchOptions{StartAt: 0, MaxResults: 10}
	issues, _, err := m.jiraClient.Issue.Search(m.settings.Query, opt)
	if issues == nil {
		m.j.Log.Message("JIRA", "No issues found")
	}
	if err != nil {
		m.j.Log.Error("JIRA", "An error occured fetching issues: "+err.Error())
	}

	// New issue set!
	if len(issues) > 0 {
		if m.workingOn.GetDataValue() != issues[0].Fields.Summary {

			if m.workingOn.ShouldUpdate() {
				m.data.IssueType = issues[0].Fields.Type.Name
				m.workingOn.SetDataValue(issues[0].Fields.Summary, true)
			}

			// Wipe Old
			m.data.LastIssues = nil

			// Make new data
			for _, issue := range issues {
				m.data.LastIssues = append(m.data.LastIssues, m.GetJSONItem(issue))
			}
			m.outputLastIssues()
		}
	} else {
		m.data.LastIssues = nil
		m.outputLastIssues()
	}
}
