package emulator

import (
	"../config"
	"../logger"
	"fmt"
	"net"
	"sync"
)

func StartReceiver(wg *sync.WaitGroup) {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", config.ReplyPort))
	if err != nil {
		logger.Fatalf("Error listen for the port %d: %s", config.ReplyPort, err.Error())
	}
	logger.Infof("Start listening on localhost:%d", config.ReplyPort)
	go mainLoop(listener, wg)
}

func mainLoop(listener net.Listener, wg *sync.WaitGroup) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Fatalf("Receiver accept error: %s", err.Error())
		}
		wg.Add(1)
		go handle(conn, wg)
	}
	logger.Infof("Stop listening on localhost:%d", config.ReplyPort)
}

func handle(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		logger.Fatal("Error read reply:", err.Error())
	}
	logger.Info(string(buf[1 : n-1]))
}
