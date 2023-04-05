package repository

import (
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/repository"

	"github.com/jinzhu/gorm"
)

type vendedorRepository struct {
	db *gorm.DB
}

func NewVendedorRepository(db *gorm.DB) repository.VendedorRepository {
	return &vendedorRepository{db}
}

func (ur *vendedorRepository) FindAll(u []*model.Vendedor) ([]*model.Vendedor, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *vendedorRepository) Create(u *model.Vendedor) (*model.Vendedor, error) {
	if err := ur.db.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
