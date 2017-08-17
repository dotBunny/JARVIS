package stats

import (
	"encoding/json"
	"strings"

	Core "../../core"
	Command "../command"
)

// Module Class
type Module struct {
	settings     *Config
	errorCount   int
	warningCount int

	stats                map[string]Stat
	dashboardDefinitions []DashboardCounterDefinition

	commandModule *Command.Module
	j             *Core.JARVIS
}

// Initialize the Stats Module
func (m *Module) Initialize(jarvisInstance *Core.JARVIS, commandModule *Command.Module) {
	// Assign JARVIS, the module is made we dont to create it like in core!
	m.j = jarvisInstance
	m.commandModule = commandModule

	m.loadConfig()
	m.j.Log.RegisterChannel("Stats", "red", m.j.Config.GetPrefix())

	// Register Status Updator
	m.j.Status.RegisterUpdator("stats", m.StatusUpdate)

	m.setupData()

	m.setupEndpoints()
	m.setupCommands()

	// Register Parser With Webserver
	m.j.WebServer.RegisterParser("stats", m.ParseWebContent)

}

func (m *Module) StatusUpdate() (int, int) {
	return m.warningCount, m.errorCount
}
func (m *Module) ParseWebContent(content string, mode string) string {

	if mode == ".json" {

		if strings.Contains(content, "[[JARVIS.stats]]") {
			responseMap := make(map[string]interface{})

			for _, stat := range m.stats {
				responseMap[stat.Key] = stat.Value
			}
			outputJSON, _ := json.Marshal(responseMap)

			content = strings.Replace(content, "[[JARVIS.stats]]", string(outputJSON), -1)
		}

		if strings.Contains(content, "[[JARVIS.stats.details]]") {
			outputJSON, _ := json.Marshal(m.dashboardDefinitions)
			content = strings.Replace(content, "[[JARVIS.stats.details]]", string(outputJSON), -1)
		}

	} else {

	}
	return content
}
