package app

import (
	"log/slog"
	"sso/internal/config"
	"sso/internal/repo/mongo_repo"
	"sso/internal/srvc"
	"sso/internal/util"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	App struct {
		Cfg   *config.Config
		Lg    *slog.Logger
		Mng   *mongo.Database
		Repos *repos
		Srvcs *srvcs
	}

	repos struct {
		User         *mongo_repo.User
		Role         *mongo_repo.Role
		Permission   *mongo_repo.Permission
		RefreshToken *mongo_repo.RefreshToken
	}

	srvcs struct {
		User         *srvc.User
		Role         *srvc.Role
		Permission   *srvc.Permission
		RefreshToken *srvc.RefreshToken
		Auth         *srvc.Auth
	}
)

func MustNew() *App {
	cfg := mustSetupConfig()
	lg := mustSetupLogger(cfg)
	mng := mustSetupMongo(cfg)

	repos := setUpRepositories(mng)
	srvcs := setUpServices(cfg, repos)

	return &App{
		Cfg:   cfg,
		Lg:    lg,
		Mng:   mng,
		Repos: repos,
		Srvcs: srvcs,
	}
}

func setUpRepositories(mng *mongo.Database) *repos {
	user := mongo_repo.NewUser(mng)
	role := mongo_repo.NewRole(mng)
	permission := mongo_repo.NewPermission(mng)
	refreshToken := mongo_repo.NewRefreshToken(mng)

	return &repos{
		User:         user,
		Role:         role,
		Permission:   permission,
		RefreshToken: refreshToken,
	}
}

func setUpServices(cfg *config.Config, repos *repos) *srvcs {
	permission := srvc.NewPermission(repos.Permission)
	role := srvc.NewRole(repos.Role, permission)
	user := srvc.NewUser(repos.User, role)
	refreshToken := srvc.NewRefreshToken(repos.RefreshToken)
	auth := srvc.NewAuth(cfg.JWT.Secret, user, refreshToken)

	return &srvcs{
		User:         user,
		Role:         role,
		Permission:   permission,
		RefreshToken: refreshToken,
		Auth:         auth,
	}
}

func mustSetupConfig() *config.Config {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	return cfg
}

func mustSetupLogger(cfg *config.Config) *slog.Logger {
	lg := util.NewLogger(cfg.Log.Level, cfg.Log.Format)

	return lg
}

func mustSetupMongo(cfg *config.Config) *mongo.Database {
	mng, err := util.NewMongo(cfg.Mongo.DB, cfg.Mongo.URL)
	if err != nil {
		panic(err)
	}

	return mng
}
