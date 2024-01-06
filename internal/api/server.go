package api

import (
	"net/http"

	"github.com/brettmostert/thredl.it/internal/server"
)

type API struct {
	ctx    *server.AppContext
	router *http.ServeMux
}

func New(ctx *server.AppContext, r *http.ServeMux) *API {
	s := &API{
		ctx:    ctx,
		router: r,
	}
	s.routes()
	return s
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	print("api.serveHttp\n")
	a.router.ServeHTTP(w, r)
}
