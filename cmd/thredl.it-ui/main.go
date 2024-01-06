package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/brettmostert/thredl.it/internal/server"
	"github.com/brettmostert/thredl.it/internal/ui"
)

func main() {
	// TODO: Move this to internal/cmd to keep this clean...
	ctx := &server.AppContext{}
	router := http.NewServeMux()
	ui.New(ctx, router)

	http := http.Server{
		Addr:         "127.0.0.1:3000",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	http.ListenAndServe()

	fmt.Println("api service has shut down")
}
