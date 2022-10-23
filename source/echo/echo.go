package echo

import (
	"io"
	"net"

	"github.com/ashmeet13/protohackers/source/utils"
)

func NewEcho() *Echo {
	return &Echo{}
}

type Echo struct{}

func (e *Echo) Handle(connection net.Conn) {
	logger := utils.GetLogger()
	logger.Info("Received new connection")

	defer connection.Close()

	_, err := io.Copy(connection, connection)
	if err != nil {
		logger.WithError(err).Error("Error in copy data")
	}
}

func (e *Echo) GetConnectionNetwork() string {
	return "tcp"
}
