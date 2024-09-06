package v1

import (
	"fmt"
	"log/slog"
	"net/http"
	"sso/internal/app"
	"sso/internal/handler/v1/request"
	"sso/internal/handler/v1/response"
)

type user struct {
	lg       *slog.Logger
	userSrvc UserSrvc
}

func newUser(
	mux *http.ServeMux,
	app *app.App,
	prefix string,
	userSrvc UserSrvc,
) {
	prefix += "/users"
	u := user{
		lg:       app.Lg,
		userSrvc: userSrvc,
	}

	mux.HandleFunc("GET "+prefix, u.list)
	mux.HandleFunc("POST "+prefix, u.create)
	mux.HandleFunc("GET "+prefix+"/{id}", u.show)
	mux.HandleFunc("PATCH "+prefix+"/{id}", u.update)
	mux.HandleFunc("DELETE "+prefix+"/{id}", u.delete)
}

func (u *user) list(w http.ResponseWriter, r *http.Request) {
	const op = "v1.user.list"

	search := request.GetQuerySearch(r)

	users, pagination, err := u.userSrvc.List(
		r.Context(),
		search.Pagination.Page,
		search.Pagination.Count,
		search.Filters,
		search.Sorts,
	)
	if err != nil {
		response.JsonFail(w, fmt.Errorf("%s: %w", op, err))
		return
	}

	response.JsonList(w, users, pagination)
}

func (u *user) create(w http.ResponseWriter, r *http.Request) {
	const op = "v1.user.create"

	var req request.UserCreate

	err := request.ParseBody(r, &req)
	if err != nil {
		response.JsonFail(w, fmt.Errorf("%s: %w", op, err))
		return
	}

	user, err := u.userSrvc.Create(
		r.Context(),
		req.Email,
		req.Name,
		req.Password,
	)
	if err != nil {
		response.JsonFail(w, fmt.Errorf("%s: %w", op, err))
		return
	}

	response.JsonSuccess(w, http.StatusCreated, user)
}

func (u *user) show(w http.ResponseWriter, r *http.Request) {
	const op = "v1.user.show"

	id := r.PathValue("id")

	user, err := u.userSrvc.Show(r.Context(), id)
	if err != nil {
		response.JsonFail(w, fmt.Errorf("%s: %w", op, err))
		return
	}

	response.JsonSuccess(w, http.StatusOK, user)
}

func (u *user) update(w http.ResponseWriter, r *http.Request) {
	const op = "v1.user.update"

	id := r.PathValue("id")
	var req request.UserUpdate

	err := request.ParseBody(r, &req)
	if err != nil {
		response.JsonFail(w, fmt.Errorf("%s: %w", op, err))
		return
	}

	user, err := u.userSrvc.Update(r.Context(), id, req.Name)
	if err != nil {
		response.JsonFail(w, fmt.Errorf("%s: %w", op, err))
		return
	}

	response.JsonSuccess(w, http.StatusOK, user)
}

func (u *user) delete(w http.ResponseWriter, r *http.Request) {
	const op = "v1.user.delete"

	id := r.PathValue("id")

	err := u.userSrvc.Delete(r.Context(), id)
	if err != nil {
		response.JsonFail(w, fmt.Errorf("%s: %w", op, err))
		return
	}

	response.JsonSuccess(w, http.StatusNoContent, nil)
}
