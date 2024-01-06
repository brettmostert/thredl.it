package ui

import (
	"net/http"

	"github.com/brettmostert/thredl.it/internal/ui/components"
)

func (u *ui) handleInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			u.postHandler(w, r)
			return
		}
		u.getHandler(w, r)
	}
}

func (u *ui) getHandler(w http.ResponseWriter, r *http.Request) {
	userCount := u.sessionManager.GetInt(r.Context(), "count")
	component := components.Page(u.state.Count, userCount)
	component.Render(r.Context(), w)
}

func (u *ui) getPartial(w http.ResponseWriter, r *http.Request) {
	userCount := u.sessionManager.GetInt(r.Context(), "count")
	component := components.Counts(u.state.Count, userCount)
	component.Render(r.Context(), w)
}

func (u *ui) postHandler(w http.ResponseWriter, r *http.Request) {
	// Update state.
	r.ParseForm()

	// Check to see if the global button was pressed.
	if r.Form.Has("global") {
		u.state.Count++
	}
	if r.Form.Has("session") {
		currentCount := u.sessionManager.GetInt(r.Context(), "count")
		u.sessionManager.Put(r.Context(), "count", currentCount+1)
	}

	// Display the form.
	u.getPartial(w, r)
}
