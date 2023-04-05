package controller

import (
	"golang-clean-architecture/pkg/usecase/usecase"
	"net/http"

	"golang-clean-architecture/pkg/domain/model"
)

type vendedorController struct {
	vendedorUsecase usecase.Vendedor
}

type Vendedor interface {
	GetVendedores(c Context) error
	CreateVendedores(c Context) error
}

func NewVendedorController(us usecase.Vendedor) Vendedor {
	return &vendedorController{us}
}

func (uc *vendedorController) GetVendedores(ctx Context) error {
	var u []*model.Vendedor

	u, err := uc.vendedorUsecase.List(u)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, u)
}

func (uc *vendedorController) CreateVendedores(ctx Context) error {
	var params model.Vendedor

	if err := ctx.Bind(&params); err != nil {
		return err
	}

	u, err := uc.vendedorUsecase.Create(&params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, u)
}
