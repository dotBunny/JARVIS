package stats

import (
	"fmt"
	"net/http"
)

func (m *StatsModule) setupEndpoints() {
	m.j.WebServer.RegisterEndpoint("/stats/workingon", m.endpointWorkingOn)
	m.j.WebServer.RegisterEndpoint("/stats/coffee", m.endpointCoffee)
	m.j.WebServer.RegisterEndpoint("/stats/saves", m.endpointSaves)
}

func (m *StatsModule) endpointCoffee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(m.data.CoffeeCount))
}
func (m *StatsModule) endpointSaves(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(m.data.SavesCount))
}
func (m *StatsModule) endpointWorkingOn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(m.data.WorkingOn))
}
