package clion

import (
	"net/http"
	"os/exec"
)

func (m *Module) setupEndpoints() {
	m.j.WebServer.RegisterEndpoint("/clion/build", m.endpointBuild)
	m.j.WebServer.RegisterEndpoint("/clion/build/", m.endpointBuild)
}

func (m *Module) endpointBuild(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command(m.buildScriptPath)
	err := cmd.Run()

	if err == nil {
		m.statsModule.IncrementBuildCount()
	}
}
