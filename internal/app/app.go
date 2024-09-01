package app

import (
	"log"
	"log/slog"
	"sso/internal/config"
	"sso/internal/util"

	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Cfg *config.Config
	Lg  *slog.Logger
	Mng *mongo.Database
}

func MustNew() *App {
	cfg := mustSetupConfig()
	lg := mustSetupLogger(cfg)
	mng := mustSetupMongo(cfg)

	return &App{
		Cfg: cfg,
		Lg:  lg,
		Mng: mng,
	}
}

func mustSetupConfig() *config.Config {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	return mng
}
