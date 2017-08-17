package main

import (
	"net"
	"net/http"
)

func Listen() {

	go http.ListenAndServe(":8080", nil)
	var ip = GetLocalIP()

	// %2F = /
	// %3F = ?
	// %3D = =
	// %3A = :
	// Register CLion
	go http.Get("http://192.168.1.250:8080/proxy/register/?id=%2Fcommand%2F%3Fscript%3Dclion&callback=http%3A%2F%2F" + ip + "%3A8080%2Fclion")

	// Build List
	http.HandleFunc("clion", BuildCLion)
}
func ShutdownListener() {
	go http.Get("http://192.168.1.250:8080/proxy/unregister/?id=%2Fcommand%2F%3Fscript%3Dclion")
}

func BuildCLion(w http.ResponseWriter, r *http.Request) {

}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
