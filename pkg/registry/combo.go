package registry

import (
	"golang-clean-architecture/pkg/adapter/controller"
	"golang-clean-architecture/pkg/adapter/repository"
	"golang-clean-architecture/pkg/usecase/usecase"
)

func (r *registry) NewComboController() controller.Combo {
	u := usecase.NewComboUsecase(
		repository.NewComboRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewComboController(u)
}
