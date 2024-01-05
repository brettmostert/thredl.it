package api

import (
	"net/http"

	"github.com/brettmostert/thredl.it/internal/server"
)

type api struct {
	ctx    *server.AppContext
	router *http.ServeMux
}

func New(ctx *server.AppContext, r *http.ServeMux) *api {
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
