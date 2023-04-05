package registry

import (
	"golang-clean-architecture/pkg/adapter/controller"
	"golang-clean-architecture/pkg/adapter/repository"
	"golang-clean-architecture/pkg/usecase/usecase"
)

func (r *registry) NewUsuarioController() controller.Usuario {
	u := usecase.NewUsuarioUsecase(
		repository.NewUsuarioRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewUsuarioController(u)
}
