package internal

import (
	"log"
	"net"
)

func IsNetworkAvailable() bool {

	_, err := net.Dial("tcp", "8.8.8.8:53")
	if err != nil {
		log.Println("no network available")
		return false
	}
	return true
}
