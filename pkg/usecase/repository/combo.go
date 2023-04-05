package repository

import "golang-clean-architecture/pkg/domain/model"

type ComboRepository interface {
	FindAll(u []*model.Combo) ([]*model.Combo, error)
	FindAllByVendedor(vendedorID int, u []*model.Combo) ([]*model.Combo, error)
	Create(u *model.Combo) (*model.Combo, error)
}
