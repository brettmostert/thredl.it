package ui

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/brettmostert/thredl.it/internal/server"
)

type GlobalState struct {
	Count int
}

type ui struct {
	ctx            *server.AppContext
	router         *http.ServeMux
	state          *GlobalState
	sessionManager *scs.SessionManager
}

func New(ctx *server.AppContext, r *http.ServeMux) *ui {
	s := &ui{
		ctx:            ctx,
		router:         r,
		state:          &GlobalState{},
		sessionManager: scs.New(),
	}

	// Todo - Add to config
	s.sessionManager.Lifetime = 2 * time.Minute
	s.routes()
	return s
}
