package server

import (
	"net"

	"github.com/ashmeet13/protohackers/source/utils"
)

func StartListener(handler Handler) {
	switch handler.GetConnectionNetwork() {
	case "tcp":
		StartTCPListener(handler.Handle)
	}
}

func StartTCPListener(handle func(connection net.Conn)) {
	logger := utils.GetLogger()
	listener, err := net.Listen("tcp", "0.0.0.0:10000")
	if err != nil {
		logger.WithError(err).Error("Error in building listener")
	}

	logger.WithField("port", "10000").Info("Listening for tcp connections")

	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			logger.WithError(err).Error("Error in accepting connection")
		}
		go handle(connection)
	}
}
