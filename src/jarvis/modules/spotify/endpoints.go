package spotify

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func (m *Module) setupEndpoints() {
	// Add Endpoints
	m.j.WebServer.RegisterEndpoint("/spotify/track", m.endpointTrack)
	m.j.WebServer.RegisterEndpoint("/spotify/image", m.endpointImage)
	m.j.WebServer.RegisterEndpoint("/spotify/track/", m.endpointTrack)
	m.j.WebServer.RegisterEndpoint("/spotify/image/", m.endpointImage)
	m.j.WebServer.RegisterEndpoint("/spotify/track/next", m.endpointNextTrack)
	m.j.WebServer.RegisterEndpoint("/spotify/track/next/", m.endpointNextTrack)
	m.j.WebServer.RegisterEndpoint("/spotify/play", m.endpointPausePlay)
	m.j.WebServer.RegisterEndpoint("/spotify/play/", m.endpointPausePlay)
	m.j.WebServer.RegisterEndpoint("/spotify/pause", m.endpointPause)
	m.j.WebServer.RegisterEndpoint("/spotify/pause/", m.endpointPause)
	m.j.WebServer.RegisterEndpoint("/spotify/info", m.endpointInfo)
	m.j.WebServer.RegisterEndpoint("/spotify/info/", m.endpointInfo)
}

func (m *Module) endpointImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(m.data.LastImageData)))
	if _, err := w.Write(m.data.LastImageData); err != nil {
		m.j.Log.Error("Spotify", "Unable to serve image via endpoint. "+err.Error())
	}
}

func (m *Module) endpointTrack(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)
	w.Header().Set("Content-Length", strconv.Itoa(len(m.data.CurrentlyPlayingTrack)))
	fmt.Fprintf(w, m.data.CurrentlyPlayingTrack)
}

func (m *Module) endpointNextTrack(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)
	m.spotifyClient.Next()
	m.j.Log.Message("Spotify", "Skipping Track")
}

func (m *Module) endpointPausePlay(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)
	if m.data.CurrentlyPlaying {
		m.spotifyClient.Pause()
		m.data.CurrentlyPlaying = false
		m.j.Log.Message("Spotify", "Pause Track")
	} else {
		m.spotifyClient.Play()
		m.data.CurrentlyPlaying = true
		m.j.Log.Message("Spotify", "Play Track")
	}
}
func (m *Module) endpointPause(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)
	m.spotifyClient.Pause()
	m.data.CurrentlyPlaying = false
	m.j.Log.Message("Spotify", "Pause Track")
}
func (m *Module) endpointPlay(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)
	m.spotifyClient.Play()
	m.data.CurrentlyPlaying = true
	m.j.Log.Message("Spotify", "Play Track")
}

func (m *Module) endpointInfo(w http.ResponseWriter, r *http.Request) {
	m.j.Discord.AnnoucementEmbed(&discordgo.MessageEmbed{
		Type:      "rich",
		Title:     "Playing On Spotify",
		URL:       m.data.CurrentlyPlayingURL,
		Color:     1947988,
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: m.data.TrackThumbnailURL},
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   m.data.TrackName,
				Value:  m.data.ArtistLine,
				Inline: true},
		},
	})
}
