package emulator

import (
	"../logger"
	"net"
)

type sender struct {
	conn net.Conn
}

func CreateSender(url string) *sender {
	sender := new(sender)
	var err error
	sender.conn, err = net.Dial("tcp", url)
	if err != nil {
		logger.Fatalf("Error connecting %s: %s", url, err.Error())
	}
	return sender
}

func (s *sender) Send(data []byte) {
	n, err := s.conn.Write(stxEtxEncode(data))
	if err != nil {
		logger.Infof("Error sending data %s", err.Error())
	} else {
		logger.Infof("[%s] %d bytes sent", data, n)
	}
}

func (s *sender) SendString(message string) {
	n, err := s.conn.Write(stxEtxEncode([]byte(message)))
	if err != nil {
		logger.Infof("Error sending data %s", err.Error())
	} else {
		logger.Infof("[%s] %d bytes sent", message, n)
	}
}

func stxEtxEncode(message []byte) []byte {
	data := make([]byte, len(message)+2)
	data[0] = 2
	for i := 0; i < len(message); i++ {
		data[i+1] = message[i]
	}
	data[len(data)-1] = 3
	return data
}
