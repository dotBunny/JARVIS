package modules

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	Core "../core"
)

// WorkingOnModule Class
type WorkingOnModule struct {
	TextPath string
	Message  string

	config *Core.Config
}

// Init  Module
func (m *WorkingOnModule) Init(config *Core.Config, console *ConsoleModule) {

	// Assing Config
	m.config = config

	// Only do this if we are going to write files
	if m.config.WorkingOn.Output {
		m.TextPath = filepath.Join(m.config.General.OutputPath, "WorkingOn_Message.txt")

		// Check twitchLatestFollowerPath
		if _, err := os.Stat(m.TextPath); os.IsNotExist(err) {
			ioutil.WriteFile(m.TextPath, nil, 0755)
		}
	}

	// Setup Endpoints
	Core.AddEndpoint("/workingon", m.workingOnEndpoint)

	// Setup Console Commands
	console.AddHandler("workingon", "Set your currently working on text.", m.consoleWorkingOn)
	console.AddAlias("w", "workingon")
}

func (m *WorkingOnModule) consoleWorkingOn(input string) {
	m.Message = input
	if m.config.WorkingOn.Output {
		Core.SaveFile([]byte(input), m.TextPath)
	}
	Core.Log("WORKING", "LOG", "Set: "+input)
}

func (m *WorkingOnModule) workingOnEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(m.Message))
}
