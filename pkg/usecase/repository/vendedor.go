package repository

import "golang-clean-architecture/pkg/domain/model"

type VendedorRepository interface {
	FindAll(u []*model.Vendedor) ([]*model.Vendedor, error)
	Create(u *model.Vendedor) (*model.Vendedor, error)
}
