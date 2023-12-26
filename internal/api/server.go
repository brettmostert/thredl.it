package api

import (
	"net/http"

	"github.com/brettmostert/thredl.it/internal/server"
	"github.com/gorilla/mux"
)

type api struct {
	ctx    *server.AppContext
	router *mux.Router
}

func New(ctx *server.AppContext, r *mux.Router) *api {
	s := &api{
		ctx:    ctx,
		router: r,
	}
	s.routes()
	return s
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	print("api.serveHttp\n")
	a.router.ServeHTTP(w, r)
}
