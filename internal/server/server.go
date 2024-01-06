package server

import (
	"net/http"
)

type AppContext struct{}

type Server struct {
	ctx    *AppContext
	router *http.ServeMux
}

func New(ctx *AppContext) *Server {
	s := &Server{
		ctx:    ctx,
		router: http.NewServeMux(),
	}

	return s
}

func (s *Server) Router() *http.ServeMux {
	// TODO: Check for duplicate routes?
	return s.router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	print("server.serveHttp\n")
	s.router.ServeHTTP(w, r)
}
