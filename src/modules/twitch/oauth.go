package twitch

import (
	"net/http"
	"time"

	"github.com/skratchdot/open-golang/open"

	"fmt"

	Core "../../core"

	"golang.org/x/oauth2"
)

var (
	ch = make(chan string)
)

// Initialize the Logging Module
func (m *Module) authenticate() {

	// Create callback endpoint
	m.j.WebServer.RegisterEndpoint("/twitch/callback", m.callbackAuthenticate)

	// OAuth Setup
	m.twitchOAuth = oauth2.Config{
		ClientID:     m.settings.ClientID,
		ClientSecret: m.settings.ClientSecret,
		RedirectURL:  "http://" + m.j.WebServer.GetIPAddress() + ":" + m.j.WebServer.GetPort() + "/twitch/callback",
		Scopes: []string{
			"channel_check_subscription",
			"channel_commercial",
			"channel_editor",
			"channel_feed_edit",
			"channel_feed_read",
			"channel_subscriptions",
			"chat_login",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL: twitchRootURL + "oauth2/authorize/",
		},
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := m.twitchOAuth.AuthCodeURL("state", oauth2.AccessTypeOffline)

	Core.CopyToClipboard(url)
	open.Run(url)

	// Wait for authentication
	temp := <-ch
	m.twitchToken = temp

	m.j.Log.Message("Twitch", "OAuth Complete.")
}

func (m *Module) callbackAuthenticate(w http.ResponseWriter, r *http.Request) {

	code := r.FormValue("code")

	if len(code) == 0 {
		m.j.Log.Warning("Twitch", "Unable to find OAuth code on return.")
		return
	}

	m.twitchClient = new(http.Client)
	m.twitchClient.Timeout = time.Second * 2

	fmt.Fprintf(w, "Login Completed! Please close this tab/window.")

	ch <- code
}
