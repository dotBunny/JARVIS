package jira

import (
	"fmt"
	"net/http"
	"strconv"
)

func (m *Module) setupEndpoints() {
	m.j.WebServer.RegisterEndpoint("/jira/inprogress", m.endpointInProgress)
	m.j.WebServer.RegisterEndpoint("/jira/inprogress/", m.endpointInProgress)
}

func (m *Module) endpointInProgress(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)
	w.Header().Set("Content-Length", strconv.Itoa(len(m.data.LastIssuesString)))
	fmt.Fprintf(w, m.data.LastIssuesString)
}
