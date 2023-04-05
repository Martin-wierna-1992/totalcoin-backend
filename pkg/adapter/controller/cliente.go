package controller

import (
	"errors"
	"fmt"
	"golang-clean-architecture/pkg/usecase/usecase"
	"net/http"

	"golang-clean-architecture/pkg/domain/model"

	"github.com/golang-jwt/jwt/v4"

	"github.com/labstack/echo"
)

type clienteController struct {
	clienteUsecase usecase.Cliente
}

type Cliente interface {
	GetClientes(c echo.Context) error
	CreateClientes(c Context) error
}

func NewClienteController(us usecase.Cliente) Cliente {
	return &clienteController{us}
}

func (uc *clienteController) GetClientes(ctx echo.Context) error {

	//// TOKEN CLAIM //// -- MOVER ESTO A UN MIDDLEWARE
	token, err := jwt.Parse(ctx.Request().Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return "TotalCoinSecret", nil
	})
	if token == nil {
		fmt.Println("invalid token")
		return errors.New("Token error")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("couldn't parse claims")
		return errors.New("Token error")
	}
	fmt.Println("ID de CLIENTE", claims["entityID"])
	/////////////////////

	var u []*model.Cliente

	u, errx := uc.clienteUsecase.List(u)
	if errx != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, u)
}

func (uc *clienteController) CreateClientes(ctx Context) error {
	var params model.Cliente

	if err := ctx.Bind(&params); err != nil {
		return err
	}

	u, err := uc.clienteUsecase.Create(&params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, u)
}
