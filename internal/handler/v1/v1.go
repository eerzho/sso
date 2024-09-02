package v1

import (
	"net/http"
	"sso/internal/app"
)

func New(
	mux *http.ServeMux,
	app *app.App,
	prefix string,
	userSrvc UserSrvc,
) {
	newUser(mux, app, prefix, userSrvc)
}
