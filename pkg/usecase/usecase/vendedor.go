package usecase

import (
	"errors"
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/repository"
)

type vendedorUsecase struct {
	vendedorRepository repository.VendedorRepository
	dBRepository       repository.DBRepository
}

type Vendedor interface {
	List(u []*model.Vendedor) ([]*model.Vendedor, error)
	Create(u *model.Vendedor) (*model.Vendedor, error)
}

func NewVendedorUsecase(r repository.VendedorRepository, d repository.DBRepository) Vendedor {
	return &vendedorUsecase{r, d}
}

func (uu *vendedorUsecase) List(u []*model.Vendedor) ([]*model.Vendedor, error) {
	u, err := uu.vendedorRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uu *vendedorUsecase) Create(u *model.Vendedor) (*model.Vendedor, error) {
	data, err := uu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := uu.vendedorRepository.Create(u)

		// do mailing
		// do logging
		// do another process
		return u, err
	})
	vendedor, ok := data.(*model.Vendedor)

	if !ok {
		return nil, errors.New("cast error")
	}

	if err != nil {
		return nil, err
	}

	return vendedor, nil
}
