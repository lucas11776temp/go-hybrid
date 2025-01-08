package server

import (
	"log"
	"strings"
)

type Server struct {
	IP_Address string
	Port       string
}

func Http(host string) Server {
	server := Server{
		IP_Address: host,
		Port:       "5555",
	}

	go func() {
		err := server.run()

		if err != nil {
			log.Fatal(err)
		}
	}()

	return server
}

func (ctx *Server) Address() string {
	return strings.Join([]string{ctx.IP_Address, ctx.Port}, ":")
}

func (ctx *Server) Host() string {
	return strings.Join([]string{"http", ctx.Address()}, "://")
}

func (s *Server) run() error {
	return nil
}
