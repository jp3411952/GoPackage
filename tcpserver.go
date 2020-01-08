package main

import (
	"net"
	"sync"
)

type TcpServer struct {
	  listen  net.Listener
	  Connets sync.Map
}
