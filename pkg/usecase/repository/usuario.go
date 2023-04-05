package repository

import "golang-clean-architecture/pkg/domain/model"

type UsuarioRepository interface {
	FindByEmail(email string, u *model.Usuario) (*model.Usuario, error)
	FindAll(u []*model.Usuario) ([]*model.Usuario, error)
	Create(u *model.Usuario) (*model.Usuario, error)
	Update(u *model.Usuario) (*model.Usuario, error)
	Delete(u *model.Usuario) (*model.Usuario, error)
}
