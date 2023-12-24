package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/brettmostert/thredl.it/server"
	"github.com/brettmostert/thredl.it/server/api"
	"github.com/brettmostert/thredl.it/server/ui"
)

func main() {
	// todo - move this to internal/cmd to keep this clean...
	ctx := &server.AppContext{}
	api := api.New(ctx)

	apiServer := http.Server{
		Addr:         "127.0.0.1:3000",
		Handler:      api,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	ui := ui.New(ctx)
	uiServer := http.Server{
		Addr:         "127.0.0.1:3001",
		Handler:      ui,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		apiServer.ListenAndServe()
	}()

	go func() {
		defer wg.Done()
		uiServer.ListenAndServe()
	}()

	wg.Wait()

	fmt.Println("service has shut down")
}
