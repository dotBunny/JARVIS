package tasks

import (
	"encoding/json"
	"strings"

	Core "../../core"
	JIRA "./jira"
)

// Module Class
type Module struct {
	UseJIRAForWork bool
	outputs        *Outputs
	data           *Data
	settings       *Config

	j            *Core.JARVIS
	jiraInstance *JIRA.Module
}

// Initialize the Stats Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS) {
	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance
	jiraInstance := new(JIRA.Module)
	m.jiraInstance = jiraInstance

	m.jiraInstance.Initialize(m.j, m.SetWorkingOnJIRA, m.GetWorkingOn)

	m.loadConfig()
	m.j.Log.RegisterChannel("Tasks", "yellow", m.j.Config.GetPrefix())

	m.setupOutputs()
	m.setupData()

	m.setupCommands()

	// Register Parser With Webserver
	m.j.WebServer.RegisterParser("tasks", m.ParseWebContent)
}

func (m *Module) ParseWebContent(content string, mode string) string {

	if mode == ".json" {

		responseMap := make(map[string]interface{})

		responseMap["Text"] = m.data.WorkingOn
		responseMap["JIRA"] = m.UseJIRAForWork

		if m.UseJIRAForWork {
			responseMap["Icon"] = m.jiraInstance.GetData().LastNotifyIcon
		}
		// responseMap["link"] =

		outputJSON, _ := json.Marshal(responseMap)

		content = strings.Replace(content, "[[JARVIS.tasks]]", string(outputJSON), -1)

	} else {

	}
	return content
}
