package repository

import (
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/repository"

	"github.com/jinzhu/gorm"
)

type usuarioRepository struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) repository.UsuarioRepository {
	return &usuarioRepository{db}
}

func (ur *usuarioRepository) FindByEmail(email string, u *model.Usuario) (*model.Usuario, error) {
	err := ur.db.Model(&u).Preload("Vendedor").Preload("Cliente").Where("email = ?", email).First(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *usuarioRepository) FindAll(u []*model.Usuario) ([]*model.Usuario, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *usuarioRepository) Create(u *model.Usuario) (*model.Usuario, error) {
	if err := ur.db.Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *usuarioRepository) Update(u *model.Usuario) (*model.Usuario, error) {
	if err := ur.db.Model(&u).Update(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *usuarioRepository) Delete(u *model.Usuario) (*model.Usuario, error) {
	if err := ur.db.Model(&u).Delete(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
