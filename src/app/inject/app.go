package inject

import (
	"merchant/pkg"
	"merchant/src/app/config"
)

type App struct {
	Cfg *config.Config
	DB  pkg.DB
}

func New(
	cfg *config.Config,
	db pkg.DB,

) *App {
	return &App{
		Cfg: cfg,
		DB:  db,
	}
}
