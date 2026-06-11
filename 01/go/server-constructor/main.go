package main

import (
	"fmt"
	"os"
	"time"
)

type Server struct {
	addr              string
	port              int
	maxIdleConnection time.Duration
	maxConns          int
}

func SetPort(p int) Option {
	return func(srv *Server) {
		srv.port = p
	}
}

func SetIdleConnections(t time.Duration) Option {
	return func(srv *Server) {
		srv.maxIdleConnection = t
	}
}

func SetMaxConns(conns int) Option {
	return func(srv *Server) {
		srv.maxConns = conns
	}
}

type Option func(*Server)

func main() {
	srv, err := NewServer("localhost", SetPort(11), SetIdleConnections(5), SetMaxConns(10))
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(srv)
}

func NewServer(addr string, opts ...Option) (*Server, error) {
	srv := &Server{addr: addr}

	for _, opt := range opts {
		opt(srv)
	}

	return srv, nil
}
