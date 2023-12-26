package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/brettmostert/thredl.it/internal/api"
	"github.com/brettmostert/thredl.it/internal/ui"
	"github.com/brettmostert/thredl.it/server"
)

func main() {
	// todo - move this to internal/cmd to keep this clean...
	ctx := &server.AppContext{}
	s := server.New(ctx)

	api.New(ctx, s.NewSubRoute("/api/v1"))
	ui.New(ctx, s.NewSubRoute(""))

	http := http.Server{
		Addr:         "127.0.0.1:3000",
		Handler:      s,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	http.ListenAndServe()

	fmt.Println("service has shut down")
}
