package handler

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	_ "sso/docs"
	"sso/internal/app"
	v1 "sso/internal/handler/v1"
	"sso/internal/repo/mongo_repo"
	"sso/internal/srvc"
)

// @Title sso http api
// @Version 1.0
// @BasePath /api
// @SecurityDefinitions.apikey BearerAuth
// @In header
// @Name Authorization
func New(app *app.App) http.Handler {
	mux := http.NewServeMux()

	// repo
	userRepo := mongo_repo.NewUser(app.Mng)
	refreshTokenRepo := mongo_repo.NewRefreshToken(app.Mng)

	// srvc
	userSrvc := srvc.NewUser(userRepo)
	refreshTokenSrvc := srvc.NewRefreshToken(refreshTokenRepo)
	authSrvc := srvc.NewAuth(app.Cfg.JWT.Secret, userSrvc, refreshTokenSrvc)

	// handler
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	handler := v1.New(mux, app, "/api/v1", userSrvc, authSrvc)

	return handler
}
