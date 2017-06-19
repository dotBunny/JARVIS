package twitch

import (
	"fmt"
	"net/http"
	"strings"
)

func (m *Module) setupEndpoints() {
	m.j.WebServer.RegisterEndpoint("/twitch/channel/game", m.endpointChannelGame)
	m.j.WebServer.RegisterEndpoint("/twitch/channel/name", m.endpointChannelViews)
	m.j.WebServer.RegisterEndpoint("/twitch/channel/views", m.endpointChannelViews)
	m.j.WebServer.RegisterEndpoint("/twitch/followers/last", m.endpointFollowersLast)
	m.j.WebServer.RegisterEndpoint("/twitch/followers/list", m.endpointFollowersList)
	m.j.WebServer.RegisterEndpoint("/twitch/followers/total", m.endpointFollowersTotal)
	m.j.WebServer.RegisterEndpoint("/twitch/subscribers/last", m.endpointSubscribersLast)
}

func (m *Module) endpointChannelGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.data.ChannelGame)
}
func (m *Module) endpointChannelName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.data.ChannelName)
}
func (m *Module) endpointChannelViews(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintln("%d", m.data.ChannelViews))
}
func (m *Module) endpointChannelViewers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintln("%d", m.data.ChannelViewers))
}
func (m *Module) endpointFollowersLast(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(m.data.LastFollower))
}
func (m *Module) endpointFollowersList(w http.ResponseWriter, r *http.Request) {

	followers := strings.Join(m.data.LastFollowers[:], ",")
	fmt.Fprintf(w, string(followers))
}
func (m *Module) endpointFollowersTotal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintln("%d", m.data.ChannelFollowers))
}
func (m *Module) endpointSubscribersLast(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(m.data.LastSubscriber))
}
