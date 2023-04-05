package repository

import (
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/repository"

	"github.com/jinzhu/gorm"
)

type clienteRepository struct {
	db *gorm.DB
}

func NewClienteRepository(db *gorm.DB) repository.ClienteRepository {
	return &clienteRepository{db}
}

func (ur *clienteRepository) FindAll(u []*model.Cliente) ([]*model.Cliente, error) {
	err := ur.db.Model(&u).Preload("Mascota").Find(&u).Error
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *clienteRepository) Create(u *model.Cliente) (*model.Cliente, error) {
	if err := ur.db.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
