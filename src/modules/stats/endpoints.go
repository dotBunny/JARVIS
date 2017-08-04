package stats

import (
	"net/http"
	"strconv"
)

func (m *Module) setupEndpoints() {

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
			m.warningCount++
			return
		}

		m.ChangeData(key, iValue, false)
	} else if mode == "increase" {
		m.ChangeData(key, m.stats[key].Value+1, true)
	} else if mode == "decrease" {
		m.ChangeData(key, m.stats[key].Value-1, false)
	}
}
