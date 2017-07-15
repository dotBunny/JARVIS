package stats

import (
	"encoding/json"
	"strings"

	Core "../../core"
	Command "../command"
)

// Module Class
type Module struct {
	UseJIRAForWork bool
	outputs        *Outputs
	data           *Data
	settings       *Config

	stats map[string]Stat

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

	m.setupOutputs()
	m.setupData()

	m.setupEndpoints()
	m.setupCommands()

	// Register Parser With Webserver
	m.j.WebServer.RegisterParser("stats", m.ParseWebContent)
}

func (m *Module) ParseWebContent(content string, mode string) string {

	if mode == ".json" {

		responseMap := make(map[string]interface{})

		for _, stat := range m.stats {
			responseMap[stat.Key] = stat.Value
		}
		outputJSON, _ := json.Marshal(responseMap)

		content = strings.Replace(content, "[[JARVIS.stats]]", string(outputJSON), -1)

	} else {

	}
	return content
}
