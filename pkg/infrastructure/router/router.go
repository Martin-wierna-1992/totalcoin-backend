package router

import (
	"golang-clean-architecture/pkg/adapter/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/api/usuarios/authenticate", func(context echo.Context) error { return c.Usuario.LoginUsuario(context) })

	e.GET("/api/usuarios", func(context echo.Context) error { return c.Usuario.GetUsuarios(context) })
	e.POST("/api/usuarios", func(context echo.Context) error { return c.Usuario.CreateUsuarios(context) })
	e.PUT("/api/usuarios", func(context echo.Context) error { return c.Usuario.UpdateUsuarios(context) })
	e.DELETE("/api/usuarios", func(context echo.Context) error { return c.Usuario.DeleteUsuarios(context) })

	e.GET("/api/clientes", func(context echo.Context) error { return c.Cliente.GetClientes(context) })
	e.POST("/api/clientes", func(context echo.Context) error { return c.Cliente.CreateClientes(context) })

	e.GET("/api/vendedores", func(context echo.Context) error { return c.Vendedor.GetVendedores(context) })
	e.POST("/api/vendedores", func(context echo.Context) error { return c.Vendedor.CreateVendedores(context) })

	e.GET("/api/mascotas", func(context echo.Context) error { return c.Mascota.GetMascotas(context) })
	e.GET("/api/mascotas/by-client", func(context echo.Context) error { return c.Mascota.GetMascotasByClienteID(context) })
	e.POST("/api/mascotas", func(context echo.Context) error { return c.Mascota.CreateMascotas(context) })

	e.GET("/api/combos", func(context echo.Context) error { return c.Combo.GetCombos(context) })
	e.GET("/api/combos/by-vendedor", func(context echo.Context) error { return c.Combo.GetCombosByVendedorID(context) })
	e.POST("/api/combos", func(context echo.Context) error { return c.Combo.CreateCombos(context) })

	return e
}
