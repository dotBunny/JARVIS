package stats

import (
	"fmt"
	"net/http"
	"strconv"
)

func (m *Module) setupEndpoints() {
	m.j.WebServer.RegisterEndpoint("/stats/workingon", m.endpointWorkingOn)

	m.j.WebServer.RegisterEndpoint("/stats/coffee", m.endpointCoffee)
	m.j.WebServer.RegisterEndpoint("/stats/coffee/", m.endpointCoffee)
	m.j.WebServer.RegisterEndpoint("/stats/coffee/plus", m.endpointCoffeePlus)
	m.j.WebServer.RegisterEndpoint("/stats/coffee/plus/", m.endpointCoffeePlus)
	m.j.WebServer.RegisterEndpoint("/stats/coffee/minus", m.endpointCoffeeMinus)
	m.j.WebServer.RegisterEndpoint("/stats/coffee/minus/", m.endpointCoffeeMinus)

	m.j.WebServer.RegisterEndpoint("/stats/saves", m.endpointSaves)
	m.j.WebServer.RegisterEndpoint("/stats/workingon/", m.endpointWorkingOn)

	m.j.WebServer.RegisterEndpoint("/stats/saves/", m.endpointSaves)
	m.j.WebServer.RegisterEndpoint("/stats/crashes/", m.endpointCrashes)
}

func (m *Module) endpointCoffee(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)
	w.Header().Set("Content-Length", strconv.Itoa(len(fmt.Sprintf("%d", m.data.CoffeeCount))))
	fmt.Fprintf(w, fmt.Sprintf("%d", m.data.CoffeeCount))
}
func (m *Module) endpointCoffeeMinus(w http.ResponseWriter, r *http.Request) {
	m.data.CoffeeCount--
}
func (m *Module) endpointCoffeePlus(w http.ResponseWriter, r *http.Request) {
	m.data.CoffeeCount++
}

func (m *Module) endpointSaves(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)
	w.Header().Set("Content-Length", strconv.Itoa(len(fmt.Sprintf("%d", m.data.SavesCount))))
	fmt.Fprintf(w, fmt.Sprintf("%d", m.data.SavesCount))
}
func (m *Module) endpointCrashes(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)
	w.Header().Set("Content-Length", strconv.Itoa(len(fmt.Sprintf("%d", m.data.CrashCount))))
	fmt.Fprintf(w, fmt.Sprintf("%d", m.data.CrashCount))
}
func (m *Module) endpointWorkingOn(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)
	w.Header().Set("Content-Length", strconv.Itoa(len(m.data.WorkingOn)))
	fmt.Fprintf(w, string(m.data.WorkingOn))
}
