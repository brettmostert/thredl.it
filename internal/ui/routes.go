package ui

import (
	"net/http"
)

func (u *UI) routes() {
	u.router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("public/assets"))))
	u.router.Handle("/", u.sessionManager.LoadAndSave(u.handleInfo()))
}

func (u *UI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u.router.ServeHTTP(w, r)
}
