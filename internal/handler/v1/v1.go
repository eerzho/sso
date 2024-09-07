package v1

import (
	"log/slog"
	"net/http"
	"sso/internal/handler/v1/mwr"
	"sso/internal/handler/v1/request"
	"sso/internal/handler/v1/response"
)

func New(
	mux *http.ServeMux,
	lg *slog.Logger,
	prefix string,
	userSrvc UserSrvc,
	authSrvc AuthSrvc,
) http.Handler {
	rp := request.NewParser()
	rb := response.NewBuilder(lg)

	authMwr := mwr.NewAuth(rb, authSrvc, userSrvc)

	newUser(mux, prefix, rp, rb, authMwr, userSrvc)
	newAuth(mux, prefix, rp, rb, authMwr, authSrvc)

	reqIDMwr := mwr.NewRequestId()
	reqLgMwr := mwr.NewRequestLogger(lg)

	return reqIDMwr.Mwr(reqLgMwr.Mwr(mux))
}
