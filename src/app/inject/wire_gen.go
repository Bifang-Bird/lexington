// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package inject

import (
	"merchant/pkg"
	"merchant/pkg/dbFactory"
	"merchant/src/app/config"
)

// Injectors from wire.go:

func InitApp(cfg *config.Config, ds *config.DataSource) (*App, func(), error) {
	db, cleanup, err := dbEngineFunc(ds)
	if err != nil {
		return nil, nil, err
	}
	app := New(cfg, db)
	return app, func() {
		cleanup()
	}, nil
}

// wire.go:

func dbEngineFunc(ds *config.DataSource) (pkg.DB, func(), error) {
	db, err := dbFactory.GetDb(*ds)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, nil
}
