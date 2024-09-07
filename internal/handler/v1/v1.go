package v1

import (
	"net/http"
	"sso/internal/app"
	"sso/internal/handler/v1/mwr"
	"sso/internal/handler/v1/request"
	"sso/internal/handler/v1/response"
)

func New(
	mux *http.ServeMux,
	app *app.App,
	prefix string,
	userSrvc UserSrvc,
	authSrvc AuthSrvc,
) http.Handler {
	rp := request.NewParser()
	rb := response.NewBuilder(app.Cfg.IsDebug, app.Lg)

	authMwr := mwr.NewAuth(rb, authSrvc, userSrvc)

	newUser(mux, prefix, rp, rb, authMwr, userSrvc)
	newAuth(mux, prefix, rp, rb, authMwr, authSrvc)

	reqIDMwr := mwr.NewRequestId()
	reqLgMwr := mwr.NewRequestLogger(app.Lg)

	return reqIDMwr.Mwr(reqLgMwr.Mwr(mux))
}
