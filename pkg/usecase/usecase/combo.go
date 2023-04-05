package usecase

import (
	"errors"
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/repository"
)

type comboUsecase struct {
	comboRepository repository.ComboRepository
	dBRepository    repository.DBRepository
}

type Combo interface {
	List(u []*model.Combo) ([]*model.Combo, error)
	ListByVendedorID(vendedorID int, u []*model.Combo) ([]*model.Combo, error)
	Create(u *model.Combo) (*model.Combo, error)
}

func NewComboUsecase(r repository.ComboRepository, d repository.DBRepository) Combo {
	return &comboUsecase{r, d}
}

func (uu *comboUsecase) List(u []*model.Combo) ([]*model.Combo, error) {
	u, err := uu.comboRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
func (uu *comboUsecase) ListByVendedorID(vendedorID int, u []*model.Combo) ([]*model.Combo, error) {
	u, err := uu.comboRepository.FindAllByVendedor(vendedorID, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uu *comboUsecase) Create(u *model.Combo) (*model.Combo, error) {
	data, err := uu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := uu.comboRepository.Create(u)

		// do mailing
		// do logging
		// do another process
		return u, err
	})
	combo, ok := data.(*model.Combo)

	if !ok {
		return nil, errors.New("cast error")
	}

	if err != nil {
		return nil, err
	}

	return combo, nil
}
