package repository

import (
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/repository"

	"github.com/jinzhu/gorm"
)

type comboRepository struct {
	db *gorm.DB
}

func NewComboRepository(db *gorm.DB) repository.ComboRepository {
	return &comboRepository{db}
}

func (ur *comboRepository) FindAll(u []*model.Combo) ([]*model.Combo, error) {
	err := ur.db.Preload("Vendedor").Preload("Mascota").Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}
func (ur *comboRepository) FindAllByVendedor(vendedorID int, u []*model.Combo) ([]*model.Combo, error) {
	err := ur.db.Preload("Vendedor").Preload("Mascota").Where("vendedor_id=?", vendedorID).Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *comboRepository) Create(u *model.Combo) (*model.Combo, error) {
	if err := ur.db.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
