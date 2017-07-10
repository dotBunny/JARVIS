package command

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"path"
	"strings"
)

func (m *Module) setupEndpoints() {
	m.j.WebServer.RegisterEndpoint("/command/", m.endpointCommand)
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
	commandLine := path.Join(m.scriptsPath, baseScript)

	// Handle Arguments
	var arguments = r.FormValue("arg")
	arguments, argErr := url.PathUnescape(arguments)
	if argErr != nil {
		m.j.Log.Error("SYSTEM", argErr.Error())
	}
	var argumentSplit = strings.Split(arguments, ",")
	for key, value := range argumentSplit {
		if value == "\"\"" {
			argumentSplit[key] = "Clear Layer"
		}
		if value == "" {
			argumentSplit[key] = "Clear Layer"
		}
	}

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
	case 6:
		commandInstance = exec.Command(commandLine, argumentSplit[0], argumentSplit[1], argumentSplit[2], argumentSplit[3], argumentSplit[4], argumentSplit[5])
		break
	default:
		commandInstance = exec.Command(commandLine)
		break
	}

	// Execute Command
	// err := commandInstance.Run()

	var out bytes.Buffer
	var stderr bytes.Buffer
	commandInstance.Stdout = &out
	commandInstance.Stderr = &stderr
	err := commandInstance.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())

	if err != nil {
		m.j.Log.Error("SYSTEM", err.Error())
	}
	// Handle CLion Build Counter
	if err == nil && baseScript == "CLion.appleScript" {
		m.j.WebServer.TouchEndpoint("/stats/builds/plus/")
	}
}
