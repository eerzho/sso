package handler

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	_ "sso/docs"
	"sso/internal/app"
	v1 "sso/internal/handler/v1"
)

// @Title sso http api
// @Version 1.0
// @BasePath /api
// @SecurityDefinitions.apikey BearerAuth
// @In header
// @Name Authorization
func New(app *app.App) http.Handler {
	mux := http.NewServeMux()

	// handler
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	handler := v1.New(mux, app, "/api/v1")

	return handler
}
