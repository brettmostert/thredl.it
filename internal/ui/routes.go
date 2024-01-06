package ui

import (
	"net/http"
)

func (u *ui) routes() {
	u.router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("public/assets"))))
	u.router.Handle("/", u.sessionManager.LoadAndSave(u.handleInfo()))
}

func (u *ui) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u.router.ServeHTTP(w, r)
}
