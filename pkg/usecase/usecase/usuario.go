package usecase

import (
	"errors"
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/repository"
)

type usuarioUsecase struct {
	usuarioRepository repository.UsuarioRepository
	dBRepository      repository.DBRepository
}

type Usuario interface {
	FindByEmail(email string, u *model.Usuario) (*model.Usuario, error)
	List(u []*model.Usuario) ([]*model.Usuario, error)
	Create(u *model.Usuario) (*model.Usuario, error)
	Update(u *model.Usuario) (*model.Usuario, error)
	Delete(u *model.Usuario) (*model.Usuario, error)
}

func NewUsuarioUsecase(r repository.UsuarioRepository, d repository.DBRepository) Usuario {
	return &usuarioUsecase{r, d}
}

func (uu *usuarioUsecase) FindByEmail(email string, u *model.Usuario) (*model.Usuario, error) {
	u, err := uu.usuarioRepository.FindByEmail(email, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
func (uu *usuarioUsecase) List(u []*model.Usuario) ([]*model.Usuario, error) {
	u, err := uu.usuarioRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uu *usuarioUsecase) Create(u *model.Usuario) (*model.Usuario, error) {
	data, err := uu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := uu.usuarioRepository.Create(u)

		// do mailing
		// do logging
		// do another process
		return u, err
	})
	usuario, ok := data.(*model.Usuario)

	if !ok {
		return nil, errors.New("cast error")
	}

	if err != nil {
		return nil, err
	}

	return usuario, nil
}
func (uu *usuarioUsecase) Update(u *model.Usuario) (*model.Usuario, error) {
	data, err := uu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := uu.usuarioRepository.Update(u)

		// do mailing
		// do logging
		// do another process
		return u, err
	})
	usuario, ok := data.(*model.Usuario)

	if !ok {
		return nil, errors.New("cast error")
	}

	if err != nil {
		return nil, err
	}

	return usuario, nil
}
func (uu *usuarioUsecase) Delete(u *model.Usuario) (*model.Usuario, error) {
	data, err := uu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := uu.usuarioRepository.Delete(u)

		// do mailing
		// do logging
		// do another process
		return u, err
	})
	usuario, ok := data.(*model.Usuario)

	if !ok {
		return nil, errors.New("cast error")
	}

	if err != nil {
		return nil, err
	}

	return usuario, nil
}
