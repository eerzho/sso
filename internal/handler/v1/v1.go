package v1

import (
	"net/http"
	"sso/internal/app"
)

func New(mux *http.ServeMux, app *app.App, prefix string) {
	newUser(mux, app, prefix)
}
