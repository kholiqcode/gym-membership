package class_booking

import (
	"github.com/google/wire"
	_rClass "gym/cmd/domain/class/repository"
	_rClassBooking "gym/cmd/domain/class_booking/repository"
	_sClassBooking "gym/cmd/domain/class_booking/service"
	_repoMember "gym/cmd/domain/member/repository"
	"gym/infrastructure/database"
	"sync"
)

var (
	hdl     *ClassBookingHandlerImpl
	hdlOnce sync.Once

	svc     *_sClassBooking.ClassBookingServiceImpl
	svcOnce sync.Once

	repo     *_rClassBooking.ClassBookingRepositoryImpl
	repoOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHandler,
		ProvideService,
		ProvideRepository,

		// bind each one of the interfaces
		wire.Bind(new(ClassBookingHandler), new(*ClassBookingHandlerImpl)),
		wire.Bind(new(_sClassBooking.ClassBookingService), new(*_sClassBooking.ClassBookingServiceImpl)),
		wire.Bind(new(_rClassBooking.ClassBookingRepository), new(*_rClassBooking.ClassBookingRepositoryImpl)),
	)
)

func ProvideHandler(svc _sClassBooking.ClassBookingService) (*ClassBookingHandlerImpl, error) {
	hdlOnce.Do(func() {
		hdl = &ClassBookingHandlerImpl{
			Svc: svc,
		}
	})

	return hdl, nil
}

func ProvideService(repo _rClassBooking.ClassBookingRepository, repoClass _rClass.ClassRepository, repoMember _repoMember.MemberRepository) (*_sClassBooking.ClassBookingServiceImpl, error) {

	svcOnce.Do(func() {
		svc = &_sClassBooking.ClassBookingServiceImpl{
			RepoClassBooking: repo,
			RepoClass:        repoClass,
			RepoMember:       repoMember,
		}
	})

	return svc, nil
}

func ProvideRepository(db *database.DatabaseImpl) (*_rClassBooking.ClassBookingRepositoryImpl, error) {

	repoOnce.Do(func() {
		repo = &_rClassBooking.ClassBookingRepositoryImpl{
			Db: db.DB,
		}
	})

	return repo, nil
}
