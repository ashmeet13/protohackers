package main

import (
	"os"

	"github.com/ashmeet13/protohackers/source/echo"
	"github.com/ashmeet13/protohackers/source/server"
	"github.com/ashmeet13/protohackers/source/utils"
)

func main() {
	logger := utils.GetLogger()

	command := os.Getenv("CMD")

	switch command {
	case "echo":
		server.StartListener(echo.NewEcho())
	default:
		logger.WithField("command", command).Error("Unknown command provided, exiting")
	}
}
