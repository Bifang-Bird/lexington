//go:build wireinject
// +build wireinject

package inject

import (
	"merchant/pkg"
	"merchant/pkg/dbFactory"
	"merchant/src/app/config"

	"github.com/google/wire"
)

func InitApp(
	cfg *config.Config,
	ds *config.DataSource,
) (*App, func(), error) {
	panic(wire.Build(
		New,
		dbEngineFunc,
	))
}
func dbEngineFunc(ds *config.DataSource) (pkg.DB, func(), error) {
	db, err := dbFactory.GetDb(*ds)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, nil
}
