package model

import "time"

type Vendedor struct {
	ID         int       `gorm:"primary_key" json:"id"`
	Nombre     string    `json:"nombre"`
	Direccion  string    `json:"direccion"`
	Telefono   string    `json:"telefono"`
	Usuario_id int       `json:"usuario_id"`
	Created_at time.Time `json:"created_at"`
	Enabled    bool      `json:"enabled"`
	Combo      []Combo   `gorm:"foreignKey:vendedor_id;references:ID"`
}

func (Vendedor) TableName() string { return "vendedores" }
