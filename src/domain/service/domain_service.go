package service

import (
	"github.com/google/wire"
	"lexington/src/domain"

	pay "lexington/src/usecases"
)

var _ pay.DomainService = (*DomainServiceImpl)(nil)

var DomainServiceSet = wire.NewSet(NewDomainServiceImpl)

type DomainServiceImpl struct {
	repo domain.DomainRepo
}

func NewDomainServiceImpl(repo domain.DomainRepo) pay.DomainService {
	return &DomainServiceImpl{
		repo: repo,
	}
}
