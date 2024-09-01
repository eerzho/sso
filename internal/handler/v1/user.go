package v1

import (
	"fmt"
	"net/http"
	"sso/internal/app"
	"sso/internal/handler/v1/response"
)

type user struct {
}

func newUser(mux *http.ServeMux, app *app.App, prefix string) {
	u := user{}
	prefix += "/users"

	mux.HandleFunc("GET "+prefix, u.list)
	mux.HandleFunc("POST "+prefix, u.create)
	mux.HandleFunc("GET "+prefix+"/{id}", u.show)
	mux.HandleFunc("PATCH "+prefix+"/{id}", u.update)
	mux.HandleFunc("DELETE "+prefix+"/{id}", u.delete)
}

func (u *user) list(w http.ResponseWriter, r *http.Request) {
	response.JsonSuccess(w, http.StatusOK, "hello from list")
}

func (u *user) create(w http.ResponseWriter, r *http.Request) {
	response.JsonSuccess(w, http.StatusCreated, "hello from create")
}

func (u *user) show(w http.ResponseWriter, r *http.Request) {
	response.JsonSuccess(w, http.StatusOK, fmt.Sprintf("hello from show %s", r.PathValue("id")))
}

func (u *user) update(w http.ResponseWriter, r *http.Request) {
	response.JsonSuccess(w, http.StatusOK, fmt.Sprintf("hello from update %s", r.PathValue("id")))
}

func (u *user) delete(w http.ResponseWriter, r *http.Request) {
	response.JsonSuccess(w, http.StatusOK, fmt.Sprintf("hello from delete %s", r.PathValue("id")))
}
