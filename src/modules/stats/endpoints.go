package stats

import (
	"fmt"
	"net/http"
)

func (m *Module) setupEndpoints() {
	m.j.WebServer.RegisterEndpoint("/stats/workingon", m.endpointWorkingOn)
	m.j.WebServer.RegisterEndpoint("/stats/coffee", m.endpointCoffee)
	m.j.WebServer.RegisterEndpoint("/stats/saves", m.endpointSaves)
	m.j.WebServer.RegisterEndpoint("/stats/workingon/", m.endpointWorkingOn)
	m.j.WebServer.RegisterEndpoint("/stats/coffee/", m.endpointCoffee)
	m.j.WebServer.RegisterEndpoint("/stats/saves/", m.endpointSaves)
}

func (m *Module) endpointCoffee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(m.data.CoffeeCount))
}
func (m *Module) endpointSaves(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(m.data.SavesCount))
}
func (m *Module) endpointWorkingOn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(m.data.WorkingOn))
}
