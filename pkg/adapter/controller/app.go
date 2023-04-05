package controller

type AppController struct {
	Usuario  interface{ Usuario }
	Cliente  interface{ Cliente }
	Vendedor interface{ Vendedor }
	Mascota  interface{ Mascota }
	Combo    interface{ Combo }
}
