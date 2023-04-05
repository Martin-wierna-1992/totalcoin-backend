package registry

import (
	"golang-clean-architecture/pkg/adapter/controller"
	"golang-clean-architecture/pkg/adapter/repository"
	"golang-clean-architecture/pkg/usecase/usecase"
)

func (r *registry) NewMascotaController() controller.Mascota {
	u := usecase.NewMascotaUsecase(
		repository.NewMascotaRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewMascotaController(u)
}
