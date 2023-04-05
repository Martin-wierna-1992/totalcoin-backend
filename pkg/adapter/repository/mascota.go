package repository

import (
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/repository"

	"github.com/jinzhu/gorm"
)

type mascotaRepository struct {
	db *gorm.DB
}

func NewMascotaRepository(db *gorm.DB) repository.MascotaRepository {
	return &mascotaRepository{db}
}

func (ur *mascotaRepository) FindAll(u []*model.Mascota) ([]*model.Mascota, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}
func (ur *mascotaRepository) FindAllByClienteID(clienteID int, u []*model.Mascota) ([]*model.Mascota, error) {
	err := ur.db.Where("cliente_id=?", clienteID).Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *mascotaRepository) Create(u *model.Mascota) (*model.Mascota, error) {
	if err := ur.db.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
