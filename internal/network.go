package internal

import (
	"log"
	"net"
	"net/http"
)

func IsNetworkAvailable() bool {
	_, err := net.Dial("tcp", "8.8.8.8:53")
	if err != nil {
		log.Println("no network available")
		return false
	}
	return true
}

func CheckIsControlPlaneServerReachable() bool {
	_, err := http.Get(FetchConfig().ServerHost + "/health")

	if err != nil {
		log.Println("control plane not reachable")
		return false
	}
	return true
}
