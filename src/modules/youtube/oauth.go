package youtube

import (
	"fmt"
	"net/http"
	"time"

	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"

	Core "../../core"
)

var (
	ch = make(chan string)
)

// Initialize the Logging Module
func (m *Module) authenticate() {

	m.j.WebServer.RegisterEndpoint("/youtube/callback", m.callbackAuthenticate)

	// OAuth Setup
	m.youtubeOAuth = oauth2.Config{
		ClientID:     m.settings.ClientID,
		ClientSecret: m.settings.ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://" + m.j.WebServer.GetIPAddress() + ":" + m.j.WebServer.GetPort() + "/youtube/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/youtube",
			"https://www.googleapis.com/auth/youtube.readonly",
		},
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := m.youtubeOAuth.AuthCodeURL("state", oauth2.AccessTypeOffline)

	Core.CopyToClipboard(url)
	open.Run(url)

	// Wait for authentication
	temp := <-ch
	m.youtubeToken = temp

	var serviceCheck error
	m.youtubeService, serviceCheck = youtube.New(m.youtubeClient)
	if serviceCheck != nil {
		m.j.Log.Error("YouTube", "Service failed to create. "+serviceCheck.Error())
		return
	}

	m.j.Log.Message("YouTube", "OAuth Complete.")
}

func (m *Module) callbackAuthenticate(w http.ResponseWriter, r *http.Request) {

	code := r.FormValue("code")

	if len(code) == 0 {
		m.j.Log.Warning("YouTube", "Unable to find OAuth code on return.")
		return
	}

	m.youtubeClient = new(http.Client)
	m.youtubeClient.Timeout = time.Second * 2

	fmt.Fprintf(w, "Login Completed! Please close this tab/window.")

	ch <- code
}
