package v1

import (
	"fmt"
	"net/http"
	"sso/internal/handler/v1/request"
	"sso/internal/handler/v1/response"
)

type auth struct {
	rp       *request.Parser
	rb       *response.Builder
	authSrvc AuthSrvc
}

func newAuth(
	mux *http.ServeMux,
	prefix string,
	rp *request.Parser,
	rb *response.Builder,
	authSrvc AuthSrvc,
) {
	prefix += "/auth"
	a := auth{
		rp:       rp,
		rb:       rb,
		authSrvc: authSrvc,
	}

	mux.HandleFunc("POST "+prefix, a.login)
}

func (a *auth) login(w http.ResponseWriter, r *http.Request) {
	const op = "v1.auth.auth"

	var req request.Login
	err := a.rp.ParseBody(r, &req)
	if err != nil {
		a.rb.JsonFail(w, r, fmt.Errorf("%s: %w", op, err))
		return
	}

	token, err := a.authSrvc.Login(r.Context(), req.Email, req.Password, a.rp.GetHeaderIP(r))
	if err != nil {
		a.rb.JsonFail(w, r, fmt.Errorf("%s: %w", op, err))
		return
	}

	a.rb.JsonSuccess(w, r, http.StatusOK, token)
}
