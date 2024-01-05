package server

import (
	"net/http"
)

type AppContext struct{}

type server struct {
	ctx    *AppContext
	router *http.ServeMux
}

func New(ctx *AppContext) *server {
	s := &server{
		ctx:    ctx,
		router: http.NewServeMux(),
	}

	return s
}

func (s *server) Router() *http.ServeMux {
	// TODO: Check for duplicate routes?
	return s.router
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	print("server.serveHttp\n")
	s.router.ServeHTTP(w, r)
}
