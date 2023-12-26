package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type AppContext struct{}

type server struct {
	ctx    *AppContext
	router *mux.Router
}

func New(ctx *AppContext) *server {
	s := &server{
		ctx:    ctx,
		router: mux.NewRouter(),
	}

	return s
}

func (s *server) NewSubRoute(r string) *mux.Router {
	return s.router.PathPrefix(r).Subrouter()
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	print("server.serveHttp\n")
	s.router.ServeHTTP(w, r)
}
