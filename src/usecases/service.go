package usecases

import (
	"github.com/Bifang-Bird/simbapkg/pkg"
	"github.com/google/wire"
	"lexington/src/domain"
)

var _ UseCase = (*service)(nil)

var UseCaseSet = wire.NewSet(NewUseCase)

type service struct {
	repo          domain.DomainRepo
	mqPub         pkg.MqPublisher
	domainService DomainService
}

func NewUseCase(
	repo domain.DomainRepo,
	domainService DomainService,
	eventPub pkg.MqPublisher,
) UseCase {
	return &service{
		repo:          repo,
		domainService: domainService,
		mqPub:         eventPub,
	}
}
