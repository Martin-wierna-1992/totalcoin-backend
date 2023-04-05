package registry

import (
	"golang-clean-architecture/pkg/adapter/controller"

	"github.com/jinzhu/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Usuario:  r.NewUsuarioController(),
		Cliente:  r.NewClienteController(),
		Vendedor: r.NewVendedorController(),
		Mascota:  r.NewMascotaController(),
		Combo:    r.NewComboController(),
	}
}