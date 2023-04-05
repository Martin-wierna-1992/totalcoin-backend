package controller

import (
	"fmt"
	"golang-clean-architecture/pkg/usecase/usecase"
	"net/http"
	"time"

	"golang-clean-architecture/pkg/domain/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type usuarioController struct {
	usuarioUsecase usecase.Usuario
}

type Usuario interface {
	LoginUsuario(c echo.Context) error
	GetUsuarios(c Context) error
	CreateUsuarios(c Context) error
	UpdateUsuarios(c Context) error
	DeleteUsuarios(c Context) error
}

func NewUsuarioController(us usecase.Usuario) Usuario {
	return &usuarioController{us}
}

func (uc *usuarioController) LoginUsuario(ctx echo.Context) error {

	var params model.Usuario

	if err := ctx.Bind(&params); err != nil {
		return err
	}
	var emailParam = params.Email
	var passwordParam = params.Password
	uc.usuarioUsecase.FindByEmail(emailParam, &params)
	fmt.Println("login cliente ++ ", params.Cliente.ID)
	fmt.Println("login vend ++ ", params.Vendedor.ID)

	var userType = ""
	var entityId = 0

	if params.Cliente.ID == 0 {
		userType = "VENDEDOR"
	} else {
		entityId = params.Cliente.ID
	}

	if params.Vendedor.ID == 0 {
		userType = "CLIENTE"
	} else {
		entityId = params.Vendedor.ID
	}

	if CheckPasswordHash(passwordParam, params.Password) == true {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)

		claims["id"] = params.ID
		claims["userType"] = userType
		claims["entityID"] = entityId
		claims["email"] = params.Email
		claims["enabled"] = params.Enabled
		claims["role"] = params.Role
		claims["exp"] = time.Now().Add(time.Minute * 120).Unix()
		var sampleSecretKey = []byte("TotalCoinSecret")
		tokenString, err := token.SignedString(sampleSecretKey)

		if err != nil {
			fmt.Errorf("Something Went Wrong: %s", err.Error())
		}

		type Respuesta struct {
			Jwt string
		}

		u := &Respuesta{Jwt: tokenString}

		fmt.Println(u)

		return ctx.JSON(http.StatusOK, u)
	}

	return ctx.JSON(401, nil)
}
func (uc *usuarioController) GetUsuarios(ctx Context) error {
	var u []*model.Usuario
	u, err := uc.usuarioUsecase.List(u)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, u)
}

func (uc *usuarioController) CreateUsuarios(ctx Context) error {
	var params model.Usuario

	if err := ctx.Bind(&params); err != nil {
		return err
	}
	params.Password, _ = HashPassword(params.Password)
	u, err := uc.usuarioUsecase.Create(&params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, u)
}

func (uc *usuarioController) UpdateUsuarios(ctx Context) error {
	var params model.Usuario

	if err := ctx.Bind(&params); err != nil {
		return err
	}
	u, err := uc.usuarioUsecase.Update(&params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, u)
}

func (uc *usuarioController) DeleteUsuarios(ctx Context) error {
	var params model.Usuario

	if err := ctx.Bind(&params); err != nil {
		return err
	}
	u, err := uc.usuarioUsecase.Delete(&params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, u)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
