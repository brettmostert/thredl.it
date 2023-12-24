package ui

import (
	"net/http"

	"github.com/brettmostert/thredl.it/server"
	"github.com/gorilla/mux"
)

type ui struct {
	ctx    *server.AppContext
	router *mux.Router
}

func New(ctx *server.AppContext) *ui {
	s := &ui{
		ctx:    ctx,
		router: mux.NewRouter(),
	}
	// s.router.PathPrefix("/ui")
	s.routes()

	return s
}

func (a *ui) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
