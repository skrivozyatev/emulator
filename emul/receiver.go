package emulator

import (
	"log"
	"net"
)

type receiver struct {
	listener net.Listener
}

func CreateReceiver(url string) *receiver {
	receiver := new(receiver)
	listener, err := net.Listen("tcp", url)
	log.Print(listener)
	if err != nil {
		log.Fatalf("Error listening url %s: %s", url, err.Error())
	}
	return receiver
}
