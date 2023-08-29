package repo

import (
	"github.com/Bifang-Bird/simbapkg/pkg"
	"github.com/google/wire"
	"lexington/src/domain"
)

var _ domain.DomainRepo = (*DomainRepoImpl)(nil)

var DomainRepositorySet = wire.NewSet(NewDomainRepoImpl)

type DomainRepoImpl struct {
	pg pkg.DB
}

func NewDomainRepoImpl(pg pkg.DB) domain.DomainRepo {
	return &DomainRepoImpl{
		pg: pg,
	}
}
