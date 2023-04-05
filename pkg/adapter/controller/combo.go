package controller

import (
	"errors"
	"fmt"
	"golang-clean-architecture/pkg/usecase/usecase"
	"net/http"
	"time"

	"golang-clean-architecture/pkg/domain/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
)

type comboController struct {
	comboUsecase usecase.Combo
}

type Combo interface {
	GetCombos(c Context) error
	GetCombosByVendedorID(c echo.Context) error
	CreateCombos(c Context) error
}

func NewComboController(us usecase.Combo) Combo {
	return &comboController{us}
}

func (uc *comboController) GetCombos(ctx Context) error {
	var u []*model.Combo

	u, err := uc.comboUsecase.List(u)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, u)
}
func (uc *comboController) GetCombosByVendedorID(ctx echo.Context) error {
	//// TOKEN CLAIM ////
	token, _ := jwt.Parse(ctx.Request().Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
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

	var u []*model.Combo

	u, err := uc.comboUsecase.ListByVendedorID(int(claims["entityID"].(float64)), u)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, u)
}

func (uc *comboController) CreateCombos(ctx Context) error {
	var params model.Combo

	if err := ctx.Bind(&params); err != nil {
		return err
	}
	params.Created_at = time.Now()
	u, err := uc.comboUsecase.Create(&params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, u)
}
