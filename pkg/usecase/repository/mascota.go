package repository

import "golang-clean-architecture/pkg/domain/model"

type MascotaRepository interface {
	FindAll(u []*model.Mascota) ([]*model.Mascota, error)
	FindAllByClienteID(clienteID int, u []*model.Mascota) ([]*model.Mascota, error)
	Create(u *model.Mascota) (*model.Mascota, error)
}
