package v1

import (
	"fmt"
	"net/http"
	"sso/internal/handler/v1/request"
	"sso/internal/handler/v1/response"
)

type user struct {
	rp       *request.Parser
	rb       *response.Builder
	userSrvc UserSrvc
}

func newUser(
	mux *http.ServeMux,
	rp *request.Parser,
	rb *response.Builder,
	prefix string,
	userSrvc UserSrvc,
) {
	prefix += "/users"
	u := user{
		rp:       rp,
		rb:       rb,
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

	search := u.rp.GetQuerySearch(r)

	users, pagination, err := u.userSrvc.List(
		r.Context(),
		search.Pagination.Page,
		search.Pagination.Count,
		search.Filters,
		search.Sorts,
	)
	if err != nil {
		u.rb.JsonFail(r, w, fmt.Errorf("%s: %w", op, err))
		return
	}

	u.rb.JsonList(r, w, users, pagination)
}

func (u *user) create(w http.ResponseWriter, r *http.Request) {
	const op = "v1.user.create"

	var req request.UserCreate

	err := u.rp.ParseBody(r, &req)
	if err != nil {
		u.rb.JsonFail(r, w, fmt.Errorf("%s: %w", op, err))
		return
	}

	user, err := u.userSrvc.Create(
		r.Context(),
		req.Email,
		req.Name,
		req.Password,
	)
	if err != nil {
		u.rb.JsonFail(r, w, fmt.Errorf("%s: %w", op, err))
		return
	}

	u.rb.JsonSuccess(r, w, http.StatusCreated, user)
}

func (u *user) show(w http.ResponseWriter, r *http.Request) {
	const op = "v1.user.show"

	id := r.PathValue("id")

	user, err := u.userSrvc.Show(r.Context(), id)
	if err != nil {
		u.rb.JsonFail(r, w, fmt.Errorf("%s: %w", op, err))
		return
	}

	u.rb.JsonSuccess(r, w, http.StatusOK, user)
}

func (u *user) update(w http.ResponseWriter, r *http.Request) {
	const op = "v1.user.update"

	id := r.PathValue("id")
	var req request.UserUpdate

	err := u.rp.ParseBody(r, &req)
	if err != nil {
		u.rb.JsonFail(r, w, fmt.Errorf("%s: %w", op, err))
		return
	}

	user, err := u.userSrvc.Update(r.Context(), id, req.Name)
	if err != nil {
		u.rb.JsonFail(r, w, fmt.Errorf("%s: %w", op, err))
		return
	}

	u.rb.JsonSuccess(r, w, http.StatusOK, user)
}

func (u *user) delete(w http.ResponseWriter, r *http.Request) {
	const op = "v1.user.delete"

	id := r.PathValue("id")

	err := u.userSrvc.Delete(r.Context(), id)
	if err != nil {
		u.rb.JsonFail(r, w, fmt.Errorf("%s: %w", op, err))
		return
	}

	u.rb.JsonSuccess(r, w, http.StatusNoContent, nil)
}
