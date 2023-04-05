package usecase

import (
	"errors"
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/repository"
)

type clienteUsecase struct {
	clienteRepository repository.ClienteRepository
	dBRepository      repository.DBRepository
}

type Cliente interface {
	List(u []*model.Cliente) ([]*model.Cliente, error)
	Create(u *model.Cliente) (*model.Cliente, error)
}

func NewClienteUsecase(r repository.ClienteRepository, d repository.DBRepository) Cliente {
	return &clienteUsecase{r, d}
}

func (uu *clienteUsecase) List(u []*model.Cliente) ([]*model.Cliente, error) {
	u, err := uu.clienteRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uu *clienteUsecase) Create(u *model.Cliente) (*model.Cliente, error) {
	data, err := uu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := uu.clienteRepository.Create(u)

		// do mailing
		// do logging
		// do another process
		return u, err
	})
	cliente, ok := data.(*model.Cliente)

	if !ok {
		return nil, errors.New("cast error")
	}

	if err != nil {
		return nil, err
	}

	return cliente, nil
}
