package command

import (
	"net/http"
	"net/url"
	"os/exec"
	"path"
	"strings"
)

func (m *Module) setupEndpoints() {
	m.j.WebServer.RegisterEndpoint("/command/", m.endpointCommand)
	m.j.WebServer.RegisterEndpoint("/command", m.endpointCommand)
}

func (m *Module) endpointCommand(w http.ResponseWriter, r *http.Request) {

	// Handle Command
	var script = r.FormValue("script")
	var baseScript string
	if script == "clion" {
		baseScript = "CLion.appleScript"
	} else if script == "wirecast" {
		baseScript = "Wirecast.appleScript"
	} else if script == "itunes" {
		// TODO: Consolidate iTunes and implement
	}

	// Handle Arguments
	var arguments = r.FormValue("arg")
	var argumentSplit = strings.Split(arguments, ",")
	for index, element := range argumentSplit {
		argumentSplit[index], _ = url.PathUnescape(element)
	}
	commandLine := path.Join(m.scriptsPath, baseScript)

	var commandInstance *exec.Cmd
	switch len(argumentSplit) {
	case 1:
		commandInstance = exec.Command(commandLine, argumentSplit[0])
		break
	case 2:
		commandInstance = exec.Command(commandLine, argumentSplit[0], argumentSplit[1])
		break
	case 3:
		commandInstance = exec.Command(commandLine, argumentSplit[0], argumentSplit[1], argumentSplit[2])
		break
	case 4:
		commandInstance = exec.Command(commandLine, argumentSplit[0], argumentSplit[1], argumentSplit[2], argumentSplit[3])
		break
	case 5:
		commandInstance = exec.Command(commandLine, argumentSplit[0], argumentSplit[1], argumentSplit[2], argumentSplit[3], argumentSplit[4])
		break
	default:
		commandInstance = exec.Command(commandLine)
		break
	}
	// Execute Command
	err := commandInstance.Run()

	// Handle CLion Build Counter
	if err == nil && baseScript == "CLion.appleScript" {
		m.statsModule.IncrementBuildCount()
	}
}
