package ui

import (
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/brettmostert/thredl.it/internal/server"
	"github.com/gorilla/mux"
)

type GlobalState struct {
	Count int
}

type ui struct {
	ctx            *server.AppContext
	router         *mux.Router
	state          *GlobalState
	sessionManager *scs.SessionManager
}

func New(ctx *server.AppContext, r *mux.Router) *ui {
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
