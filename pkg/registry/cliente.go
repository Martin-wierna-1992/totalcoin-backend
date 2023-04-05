package registry

import (
	"golang-clean-architecture/pkg/adapter/controller"
	"golang-clean-architecture/pkg/adapter/repository"
	"golang-clean-architecture/pkg/usecase/usecase"
)

func (r *registry) NewClienteController() controller.Cliente {
	u := usecase.NewClienteUsecase(
		repository.NewClienteRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewClienteController(u)
}
