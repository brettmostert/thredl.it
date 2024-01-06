package ui

import (
	"net/http"

	"github.com/brettmostert/thredl.it/internal/ui/components"
)

func (u *UI) handleInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			u.postHandler(w, r)
			return
		}
		u.getHandler(w, r)
	}
}

func (u *UI) getHandler(w http.ResponseWriter, r *http.Request) {
	userCount := u.sessionManager.GetInt(r.Context(), "count")
	component := components.Page(u.state.Count, userCount)
	component.Render(r.Context(), w)
}

func (u *UI) getPartial(w http.ResponseWriter, r *http.Request) {
	userCount := u.sessionManager.GetInt(r.Context(), "count")
	component := components.Counts(u.state.Count, userCount)
	component.Render(r.Context(), w)
}

func (u *UI) postHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Form == nil {
		// TODO: Implement nicer error handling
		panic("This should not be nil")
	}
	// Check to see if the global button was pressed.
	if r.Form.Has("global") {
		print("global")
		u.state.Count++
	}
	if r.Form.Has("session") {
		currentCount := u.sessionManager.GetInt(r.Context(), "count")
		u.sessionManager.Put(r.Context(), "count", currentCount+1)
	}

	// Display the form.
	u.getPartial(w, r)
}
