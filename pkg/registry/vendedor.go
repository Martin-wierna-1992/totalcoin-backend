package registry

import (
	"golang-clean-architecture/pkg/adapter/controller"
	"golang-clean-architecture/pkg/adapter/repository"
	"golang-clean-architecture/pkg/usecase/usecase"
)

func (r *registry) NewVendedorController() controller.Vendedor {
	u := usecase.NewVendedorUsecase(
		repository.NewVendedorRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewVendedorController(u)
}
