package api

import (
	"net/http"

	"github.com/brettmostert/thredl.it/server"
	"github.com/gorilla/mux"
)

type api struct {
	ctx    *server.AppContext
	router *mux.Router
}

func New(ctx *server.AppContext) *api {
	s := &api{
		ctx:    ctx,
		router: mux.NewRouter(),
	}

	s.routes()

	return s
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
