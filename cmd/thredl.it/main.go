package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/brettmostert/thredl.it/internal/api"
	"github.com/brettmostert/thredl.it/internal/server"
	"github.com/brettmostert/thredl.it/internal/ui"
)

func main() {
	ctx := &server.AppContext{}
	server := server.New(ctx)

	api.New(ctx, server.Router())
	ui.New(ctx, server.Router())

	http := http.Server{
		Addr:         "127.0.0.1:3000",
		Handler:      server,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err := http.ListenAndServe()
	if err != nil {
		// TODO: Implement nicer error handling
		panic("oh no")
	}

	fmt.Println("service has shut down")
}
