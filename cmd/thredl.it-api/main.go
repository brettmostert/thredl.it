package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/brettmostert/thredl.it/internal/api"
	"github.com/brettmostert/thredl.it/internal/server"
)

func main() {
	ctx := &server.AppContext{}
	router := http.NewServeMux()
	api.New(ctx, router)

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
