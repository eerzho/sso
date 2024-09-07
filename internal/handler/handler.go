package handler

import (
	"net/http"
	"sso/internal/app"
	v1 "sso/internal/handler/v1"
)

func New(
	app *app.App,
	userSrvc v1.UserSrvc,
	authSrvc v1.AuthSrvc,
) http.Handler {
	mux := http.NewServeMux()

	handler := v1.New(mux, app, "/api/v1", userSrvc, authSrvc)

	return handler
}
