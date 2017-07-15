package stats

import (
	"fmt"
	"net/http"
	"strconv"
)

func (m *Module) setupEndpoints() {

	// TODO: LEGACY
	m.j.WebServer.RegisterEndpoint("/stats/workingon", m.endpointWorkingOn)
	m.j.WebServer.RegisterEndpoint("/stats/workingon/", m.endpointWorkingOn)

	m.j.WebServer.RegisterEndpoint("/stats/", m.endpointStats)
	m.j.WebServer.RegisterEndpoint("/stats", m.endpointStats)

}

func (m *Module) endpointStats(w http.ResponseWriter, r *http.Request) {

	// Handle Command
	var mode = r.FormValue("mode")
	var key = r.FormValue("key")

	if mode == "set" {
		var value = r.FormValue("value")
		iValue, err := strconv.Atoi(value)

		if err != nil {
			m.j.Log.Warning("STATS", "Invalid value attempting to be parsed: "+value)
			return
		}

		m.ChangeData(key, iValue, false)
	} else if mode == "increase" {
		m.ChangeData(key, m.stats[key].Value+1, true)
	} else if mode == "decrease" {
		m.ChangeData(key, m.stats[key].Value-1, false)
	}
}

func (m *Module) endpointWorkingOn(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)

	if m.UseJIRAForWork {
		output := m.data.WorkingOnIcon + "," + m.data.WorkingOn
		w.Header().Set("Content-Length", strconv.Itoa(len(output)))
		fmt.Fprintf(w, output)
	} else {
		w.Header().Set("Content-Length", strconv.Itoa(len(m.data.WorkingOn)))
		fmt.Fprintf(w, string(m.data.WorkingOn))
	}
}
