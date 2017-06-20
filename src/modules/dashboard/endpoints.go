package dashboard

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (m *Module) setupEndpoints() {
	m.j.WebServer.RegisterEndpoint("/dashboard", m.endpointDashboard)
	m.j.WebServer.RegisterEndpoint("/dashboard/", m.endpointDashboard)
}

func (m *Module) endpointDashboard(w http.ResponseWriter, r *http.Request) {

	pageData, error := ioutil.ReadFile(m.dashboardPath)
	if error != nil {
		m.j.Log.Error("Overlay", "Unable to read base HTML page ("+m.dashboardPath+") from www folder.")
	}

	if len(pageData) <= 0 {
		m.j.Log.Error("Overlay", "No data to serve. Length is off.")
		fmt.Fprintf(w, "No Overlay Found")
	} else {
		w.Header().Set("Content-Length", strconv.Itoa(len(pageData)))
		fmt.Fprintf(w, string(pageData))
	}
}
