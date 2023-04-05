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

type mascotaController struct {
	mascotaUsecase usecase.Mascota
}

type Mascota interface {
	GetMascotas(c Context) error
	GetMascotasByClienteID(c echo.Context) error
	CreateMascotas(c Context) error
}

func NewMascotaController(us usecase.Mascota) Mascota {
	return &mascotaController{us}
}

func (uc *mascotaController) GetMascotas(ctx Context) error {
	var u []*model.Mascota

	u, err := uc.mascotaUsecase.List(u)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, u)
}
func (uc *mascotaController) GetMascotasByClienteID(ctx echo.Context) error {
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

	var u []*model.Mascota
	u, errx := uc.mascotaUsecase.ListByClienteID(int(claims["entityID"].(float64)), u)
	if errx != nil {
		return errx
	}

	return ctx.JSON(http.StatusOK, u)
}

func (uc *mascotaController) CreateMascotas(ctx Context) error {
	var params model.Mascota

	if err := ctx.Bind(&params); err != nil {
		return err
	}
	params.Created_at = time.Now()
	params.Enabled = true
	u, err := uc.mascotaUsecase.Create(&params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, u)
}
