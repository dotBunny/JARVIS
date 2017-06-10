package core

import "net/http"

var (
	port string
)

// InitializeWebServer for callbacks
func InitializeWebServer(listeningPort string) {
	port = listeningPort

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			Log("Webserver", "Got request for: "+r.URL.String())
		})

	go http.ListenAndServe(":"+listeningPort, nil)
}

// AddEndpoint to webserver
func AddEndpoint(endpoint string, function http.HandlerFunc) {
	http.HandleFunc(endpoint, function)
}

// func WebserverStart(port string) {
// 	go http.ListenAndServe(":"+port, nil)
// }
