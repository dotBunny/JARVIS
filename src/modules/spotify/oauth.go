package spotify

import (
	"fmt"
	"log"
	"net/http"

	Core "../../core"
	"github.com/skratchdot/open-golang/open"
	"github.com/zmb3/spotify"
)

var (
	ch = make(chan *spotify.Client)
)

func (m *Module) authenticate() {

	m.spotifyOAuth = spotify.NewAuthenticator("http://"+m.j.WebServer.GetIPAddress()+":"+m.j.WebServer.GetPort()+"/spotify/callback",
		spotify.ScopeUserReadCurrentlyPlaying,
		spotify.ScopeUserReadRecentlyPlayed,
		spotify.ScopeUserModifyPlaybackState)

	m.spotifyOAuth.SetAuthInfo(m.settings.ClientID, m.settings.ClientSecret)

	// Register auth
	m.j.WebServer.RegisterEndpoint("/spotify/callback", m.callbackAuthenticate)

	// Create State
	m.stateHash = Core.RandomString(5)

	// Generate url
	url := m.spotifyOAuth.AuthURL(m.stateHash)

	// TODO: Disabled because of being on local machine, this will get added if we go remote?
	// _, _ = m.j.Discord.GetSession().ChannelMessageSendEmbed(m.j.Discord.GetPrivateChannelID(), &discordgo.MessageEmbed{
	// 	Type:        "rich",
	// 	Title:       "Spotify Login Required",
	// 	URL:         url,
	// 	Description: "An OAuth2 token is required for the Spotify Module to operate properly. You must login via the provided link, allowing the access requested.",
	// 	Color:       7005032,
	// })

	//m.j.Log.Warning("Spotify", "Please log in to Spotify (URL copied to your clipboard as well): "+url)
	Core.CopyToClipboard(url)
	open.Run(url)

	// wait for Auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser()
	if err != nil {
		m.errorCount++
		m.j.Log.Error("Spotify", err.Error())
	}
	m.j.Log.Message("Spotify", "You are logged in as: "+user.ID)

	// Assign Client
	m.spotifyClient = client
}

func (m *Module) callbackAuthenticate(w http.ResponseWriter, r *http.Request) {
	tok, err := m.spotifyOAuth.Token(m.stateHash, r)
	if err != nil {
		m.errorCount++
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != m.stateHash {
		m.errorCount++
		http.NotFound(w, r)
		m.j.Log.Error("Spotify", fmt.Sprintf("State mismatch: %s != %s\n", st, m.stateHash))
	}
	// use the token to get an authenticated client
	client := m.spotifyOAuth.NewClient(tok)

	fmt.Fprintf(w, "Login Completed! Please close this tab/window.")

	ch <- &client
}
