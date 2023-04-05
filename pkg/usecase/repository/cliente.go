package repository

import "golang-clean-architecture/pkg/domain/model"

type ClienteRepository interface {
	FindAll(u []*model.Cliente) ([]*model.Cliente, error)
	Create(u *model.Cliente) (*model.Cliente, error)
}
