package server

import "net"

type Handler interface {
	GetConnectionNetwork() string
	Handle(connection net.Conn)
}
