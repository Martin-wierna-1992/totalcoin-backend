package usecase

import (
	"errors"
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/repository"
)

type mascotaUsecase struct {
	mascotaRepository repository.MascotaRepository
	dBRepository      repository.DBRepository
}

type Mascota interface {
	List(u []*model.Mascota) ([]*model.Mascota, error)
	ListByClienteID(clienteID int, u []*model.Mascota) ([]*model.Mascota, error)
	Create(u *model.Mascota) (*model.Mascota, error)
}

func NewMascotaUsecase(r repository.MascotaRepository, d repository.DBRepository) Mascota {
	return &mascotaUsecase{r, d}
}

func (uu *mascotaUsecase) List(u []*model.Mascota) ([]*model.Mascota, error) {
	u, err := uu.mascotaRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
func (uu *mascotaUsecase) ListByClienteID(clienteID int, u []*model.Mascota) ([]*model.Mascota, error) {
	u, err := uu.mascotaRepository.FindAllByClienteID(clienteID, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uu *mascotaUsecase) Create(u *model.Mascota) (*model.Mascota, error) {
	data, err := uu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := uu.mascotaRepository.Create(u)

		// do mailing
		// do logging
		// do another process
		return u, err
	})
	mascota, ok := data.(*model.Mascota)

	if !ok {
		return nil, errors.New("cast error")
	}

	if err != nil {
		return nil, err
	}

	return mascota, nil
}
