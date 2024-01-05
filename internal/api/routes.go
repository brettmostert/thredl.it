package api

import (
	"encoding/json"
	"net/http"
)

const baseUrl = "/api/v1"

func (a *api) routes() {
	a.router.HandleFunc(baseUrl+"/info", contentTypeApplicationJsonMiddleware(a.handleInfo()))
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
