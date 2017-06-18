package dashboard

func (m *DashboardModule) setupEndpoints() {
	// // Add Endpoints
	// m.j.WebServer.RegisterEndpoint("/spotify/track", m.endpointTrack)
	// m.j.WebServer.RegisterEndpoint("/spotify/image", m.endpointImage)
}

// func (m *SpotifyModule) endpointImage(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "image/jpeg")
// 	w.Header().Set("Content-Length", strconv.Itoa(len(m.data.LastImageData)))
// 	if _, err := w.Write(m.data.LastImageData); err != nil {
// 		m.j.Log.Error("Spotify", "Unable to serve image via endpoint. "+err.Error())
// 	}
// }

// func (m *OverlayModule) endpointOverlay(w http.ResponseWriter, r *http.Request) {

// 	// Server Page Per Time
// 	basePageData, error := ioutil.ReadFile(m.basePath)
// 	if error != nil {
// 		Core.Log("OVERLAY", "ERROR", "Unable to read base HTML page ("+m.basePath+") from resources folder.")
// 	} else {
// 		m.basePage = string(basePageData)
// 	}

// 	if len(m.basePage) <= 0 {
// 		Core.Log("OVERLAY", "ERROR", "No data to serve for overlay.")
// 		fmt.Fprintf(w, "No Overlay Found")
// 	} else {
// 		fmt.Fprintf(w, m.basePage)
// 	}
// }
