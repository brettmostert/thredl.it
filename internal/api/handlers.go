package api

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Text string
}

func (a *API) handleInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := Message{Text: "Moo"}
		res, _ := json.Marshal(payload)
		_, err := w.Write(res)
		if err != nil {
			// TODO: Implement nicer error handling
			panic("oh no")
		}
	}
}
