package youtube

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"

	Core "../../core"
)

var (
	ch = make(chan *oauth2.Token)
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

	tok, err := m.youtubeOAuth.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}

	fmt.Fprintf(w, "Login Completed! Please close this tab/window.")
	ch <- tok
}

// // getClient uses a Context and Config to retrieve a Token
// // then generate a Client. It returns the generated Client.
// func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
// 	cacheFile, err := tokenCacheFile()
// 	if err != nil {
// 		log.Fatalf("Unable to get path to cached credential file. %v", err)
// 	}
// 	tok, err := tokenFromFile(cacheFile)
// 	if err != nil {
// 		tok = getTokenFromWeb(config)
// 		saveToken(cacheFile, tok)
// 	}
// 	return config.Client(ctx, tok)
// }

// // getTokenFromWeb uses Config to request a Token.
// // It returns the retrieved Token.
// func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
// 	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
// 	fmt.Printf("Go to the following link in your browser then type the "+
// 		"authorization code: \n%v\n", authURL)

// 	var code string
// 	if _, err := fmt.Scan(&code); err != nil {
// 		log.Fatalf("Unable to read authorization code %v", err)
// 	}

// 	tok, err := config.Exchange(oauth2.NoContext, code)
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve token from web %v", err)
// 	}
// 	return tok
// }

// // tokenCacheFile generates credential file path/filename.
// // It returns the generated credential path/filename.
// func tokenCacheFile() (string, error) {
// 	usr, err := user.Current()
// 	if err != nil {
// 		return "", err
// 	}
// 	tokenCacheDir := filepath.Join(usr.HomeDir, ".jarvis-credentials")
// 	os.MkdirAll(tokenCacheDir, 0700)
// 	return filepath.Join(tokenCacheDir,
// 		url.QueryEscape("youtube-go-quickstart.json")), err
// }

// // tokenFromFile retrieves a Token from a given file path.
// // It returns the retrieved Token and any read error encountered.
// func tokenFromFile(file string) (*oauth2.Token, error) {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		return nil, err
// 	}
// 	t := &oauth2.Token{}
// 	err = json.NewDecoder(f).Decode(t)
// 	defer f.Close()
// 	return t, err
// }

// // saveToken uses a file path to create a file and store the
// // token in it.
// func saveToken(file string, token *oauth2.Token) {
// 	fmt.Printf("Saving credential file to: %s\n", file)
// 	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
// 	if err != nil {
// 		log.Fatalf("Unable to cache oauth token: %v", err)
// 	}
// 	defer f.Close()
// 	json.NewEncoder(f).Encode(token)
// }

// func handleError(err error, message string) {
// 	if message == "" {
// 		message = "Error making API call"
// 	}
// 	if err != nil {
// 		log.Fatalf(message+": %v", err.Error())
// 	}
// }
