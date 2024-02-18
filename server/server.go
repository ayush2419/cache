package server

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/ayush2419/cache/internal/command"
	"github.com/ayush2419/cache/pkg"
)

const (
	networkTypeTCP = "tcp"
)

type ServerConfig struct {
	ListenAddr string
	IsLeader   bool
}

type Server struct {
	ServerConfig
	handler *pkg.Handler
}

func NewServer(serverConfig ServerConfig, handler *pkg.Handler) *Server {
	return &Server{ServerConfig: serverConfig, handler: handler}
}

func (s *Server) Start() error {
	ln, err := net.Listen(networkTypeTCP, s.ListenAddr)
	if err != nil {
		return fmt.Errorf("listen error: %s", err)
	}

	log.Printf("server started on port: [%s]\n", s.ListenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("error accepting conn: %s", err)
			return err
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 2048)

	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatalf("conn read error: %s", err)
	}

	msg := buffer[:n]

	fmt.Printf("message received: %s", string(msg))

	err = s.getHandler(conn, string(msg))
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func (s *Server) getHandler(conn net.Conn, message string) (err error) {
	if message == "" {
		return fmt.Errorf("empty command")
	}

	parts := strings.Split(message, " ")

	if len(parts) < 2 {
		return fmt.Errorf("invalid command")
	}

	msg, err := command.ParseMessage(message)
	if err != nil {
		return
	}

	s.handler.HandleRequestMessage(conn, msg)
	return
}
