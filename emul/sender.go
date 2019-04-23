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

func (s *sender) Send(data []byte) []byte {
	n, err := s.conn.Write(stxEtxEncode(data))
	if err != nil {
		logger.Infof("Error sending data %s", err.Error())
		return []byte("Error sending data " + err.Error())
	}
	logger.Infof("[%s] %d bytes sent", string(data), n)
	buf := make([]byte, 1024)
	readBytes, readError := s.conn.Read(buf)
	if readError != nil {
		logger.Infof("Error receiving reply %s", readError.Error())
		return []byte("Error receiving reply " + readError.Error())
	}
	return buf[1 : readBytes-1]
}

func (s *sender) SendString(message string) string {
	return string(s.Send([]byte(message)))
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
