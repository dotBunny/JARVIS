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

	m.loadConfig()
	m.j.Log.RegisterChannel("Tasks", "yellow", m.j.Config.GetPrefix())
	m.UseJIRAForWork = m.settings.UseJIRAByDefault

	m.setupOutputs()
	m.setupData()

	// Create our callback niceness
	jiraModifier := new(JIRAModifier)
	jiraModifier.tasks = m
	m.jiraInstance.Initialize(m.j, jiraModifier)

	m.setupCommands()

	// Register Parser With Webserver
	m.j.WebServer.RegisterParser("tasks", m.ParseWebContent)
}

func (m *Module) ParseWebContent(content string, mode string) string {

	if mode == ".json" {

		if strings.Contains(content, "[[JARVIS.tasks]]") {
			responseMap := make(map[string]interface{})

			responseMap["Text"] = m.data.WorkingOn
			responseMap["JIRA"] = m.UseJIRAForWork

			if m.UseJIRAForWork {
				responseMap["Type"] = m.jiraInstance.GetData().IssueType
			}
			// responseMap["link"] =
			outputJSON, _ := json.Marshal(responseMap)
			content = strings.Replace(content, "[[JARVIS.tasks]]", string(outputJSON), -1)
		}

		if strings.Contains(content, "[[JARVIS.tasks.jira]]") {

			//responseMap["Text"] = m.data.WorkingOn
			//responseMap["JIRA"] = m.UseJIRAForWork

			if m.jiraInstance.GetData().LastIssues != nil {
				//responseMap["Type"] = m.jiraInstance.GetData().IssueType
				//		responseMap := make(map[string]interface{})

				// for key, value := range m.jiraInstance.GetData().LastIssues {
				// 	keyValue := strconv.Itoa(key)
				// 	responseMap[keyValue] = value
				// }

				outputJSON, _ := json.Marshal(m.jiraInstance.GetLastIssues())
				content = strings.Replace(content, "[[JARVIS.tasks.jira]]", string(outputJSON), -1)
			} else {
				content = strings.Replace(content, "[[JARVIS.tasks.jira]]", "{}", -1)
			}
		}
	} else {

	}
	return content
}
