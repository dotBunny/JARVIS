package core

import (
	"net/http"
	"strconv"
)

var (
	port int
)

// InitializeWebServer for callbacks
func InitializeWebServer(listeningPort int) {
	port = listeningPort

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			Log("SYSTEM", "LOG", "Got request for: "+r.URL.String())
		})

	go http.ListenAndServe(":"+strconv.Itoa(listeningPort), nil)
}

// AddEndpoint to webserver
func AddEndpoint(endpoint string, function http.HandlerFunc) {
	http.HandleFunc(endpoint, function)
}
