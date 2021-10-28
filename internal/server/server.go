package server

import (
	"fmt"
	"net"
	"net/http"
)

// Server is an HTTP server.
type Server struct {
	nl net.Listener
	r  *http.ServeMux
	h  *http.Server
	cr ClientRegisterer
}

// NewServer returns an HTTP server configured to run on the given host and port.
func NewServer(port, host string, cr ClientRegisterer) (s *Server, err error) {
	s = &Server{
		r:  http.NewServeMux(),
		cr: cr,
	}

	addr := net.JoinHostPort(host, port)
	s.nl, err = net.Listen("tcp4", addr)
	if err != nil {
		return nil, fmt.Errorf("net listen: %w", err)
	}

	s.routes()

	s.h = &http.Server{
		Addr:    addr,
		Handler: s.r,
	}

	return s, nil
}

// Serve accepts incoming connections.
func (s *Server) Serve() error {
	return s.h.Serve(s.nl)
}

// Close closes the HTTP server.
func (s *Server) Close() error {
	return s.h.Close()
}

func (s *Server) routes() {
	s.r.HandleFunc("/register", s.register)
}
