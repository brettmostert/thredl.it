package api

import (
	"encoding/json"
	"net/http"
)

func (a *api) routes() {
	a.router.HandleFunc("/info", a.handleInfo())
	a.router.Use(contentTypeApplicationJsonMiddleware)
}

type Message struct {
	Text string
}

func (a *api) handleInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := Message{Text: "Moo"}
		res, _ := json.Marshal(payload)
		w.Write(res)
	}
}
