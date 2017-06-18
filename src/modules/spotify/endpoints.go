package spotify

import (
	"fmt"
	"net/http"
	"strconv"
)

func (m *Module) setupEndpoints() {
	// Add Endpoints
	m.j.WebServer.RegisterEndpoint("/spotify/track", m.endpointTrack)
	m.j.WebServer.RegisterEndpoint("/spotify/image", m.endpointImage)
	m.j.WebServer.RegisterEndpoint("/spotify/track/", m.endpointTrack)
	m.j.WebServer.RegisterEndpoint("/spotify/image/", m.endpointImage)
}

func (m *Module) endpointImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(m.data.LastImageData)))
	if _, err := w.Write(m.data.LastImageData); err != nil {
		m.j.Log.Error("Spotify", "Unable to serve image via endpoint. "+err.Error())
	}
}

func (m *Module) endpointTrack(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(m.data.CurrentlyPlayingTrack))
}