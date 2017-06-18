package overlay

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (m *Module) setupEndpoints() {
	m.j.WebServer.RegisterEndpoint("/overlay", m.endpointOverlay)
	m.j.WebServer.RegisterEndpoint("/overlay/", m.endpointOverlay)
}

func (m *Module) endpointOverlay(w http.ResponseWriter, r *http.Request) {

	pageData, error := ioutil.ReadFile(m.overlayPath)
	if error != nil {
		m.j.Log.Error("Overlay", "Unable to read base HTML page ("+m.overlayPath+") from www folder.")
	}

	if len(pageData) <= 0 {
		m.j.Log.Error("Overlay", "No data to serve. Length is off.")
		fmt.Fprintf(w, "No Overlay Found")
	} else {
		w.Header().Set("Content-Length", strconv.Itoa(len(pageData)))
		fmt.Fprintf(w, string(pageData))
	}
}
