package twitch

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (m *Module) setupEndpoints() {
	m.j.WebServer.RegisterEndpoint("/twitch/channel/game", m.endpointChannelGame)
	m.j.WebServer.RegisterEndpoint("/twitch/channel/name", m.endpointChannelName)
	m.j.WebServer.RegisterEndpoint("/twitch/channel/views", m.endpointChannelViews)
	m.j.WebServer.RegisterEndpoint("/twitch/channel/viewers", m.endpointChannelViewers)
	m.j.WebServer.RegisterEndpoint("/twitch/followers/last", m.endpointFollowersLast)
	m.j.WebServer.RegisterEndpoint("/twitch/followers/list", m.endpointFollowersList)
	m.j.WebServer.RegisterEndpoint("/twitch/followers/total", m.endpointFollowersTotal)
	m.j.WebServer.RegisterEndpoint("/twitch/subscribers/last", m.endpointSubscribersLast)
	m.j.WebServer.RegisterEndpoint("/twitch/info", m.endpointInfo)
	m.j.WebServer.RegisterEndpoint("/twitch/info/", m.endpointInfo)
}

func (m *Module) endpointChannelGame(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)
	w.Header().Set("Content-Length", strconv.Itoa(len(m.data.ChannelGame)))
	fmt.Fprintf(w, m.data.ChannelGame)
}
func (m *Module) endpointChannelName(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)
	w.Header().Set("Content-Length", strconv.Itoa(len(m.data.ChannelName)))
	fmt.Fprintf(w, m.data.ChannelName)
}
func (m *Module) endpointChannelViews(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)
	w.Header().Set("Content-Length", strconv.Itoa(len(fmt.Sprintf("%d", m.data.ChannelViews))))
	fmt.Fprintf(w, fmt.Sprintf("%d", m.data.ChannelViews))
}
func (m *Module) endpointChannelViewers(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)
	w.Header().Set("Content-Length", strconv.Itoa(len(fmt.Sprintf("%d", m.data.ChannelViewers))))
	fmt.Fprintf(w, fmt.Sprintf("%d", m.data.ChannelViewers))
}
func (m *Module) endpointFollowersLast(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)
	w.Header().Set("Content-Length", strconv.Itoa(len(m.data.LastFollower)))
	fmt.Fprintf(w, string(m.data.LastFollower))
}
func (m *Module) endpointFollowersList(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)
	followers := strings.Join(m.data.LastFollowers[:], ",")
	w.Header().Set("Content-Length", strconv.Itoa(len(followers)))
	fmt.Fprintf(w, followers)
}
func (m *Module) endpointFollowersTotal(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)
	w.Header().Set("Content-Length", strconv.Itoa(len(fmt.Sprintf("%d", m.data.ChannelFollowers))))
	fmt.Fprintf(w, fmt.Sprintf("%d", m.data.ChannelFollowers))
}
func (m *Module) endpointSubscribersLast(w http.ResponseWriter, r *http.Request) {
	m.j.WebServer.DefaultHeader(w)
	w.Header().Set("Content-Length", strconv.Itoa(len(m.data.LastSubscriber)))
	fmt.Fprintf(w, m.data.LastSubscriber)

}

func (m *Module) endpointInfo(w http.ResponseWriter, r *http.Request) {
	m.j.Discord.AnnoucementEmbed(&discordgo.MessageEmbed{
		Type:      "rich",
		Title:     "Streaming on Twitch",
		URL:       "https://www.twitch.tv/" + m.twitchStreamName,
		Color:     6570404,
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: m.data.StreamPreviewURL},
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Viewers",
				Value:  fmt.Sprintf("%d", m.data.ChannelViewers),
				Inline: true},
			&discordgo.MessageEmbedField{
				Name:   "Followers",
				Value:  fmt.Sprintf("%d", m.data.ChannelFollowers),
				Inline: true},
		},
	})
}
