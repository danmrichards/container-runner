package server

import "net/http"

func (s *Server) register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("registered with ID: " + s.cr.Register()))
}
