package model

import (
	"time"
)

type Cliente struct {
	ID         int       `gorm:"primary_key" json:"id"`
	Nombre     string    `json:"nombre"`
	Direccion  string    `json:"direccion"`
	Telefono   string    `json:"telefono"`
	Usuario_id int       `json:"usuario_id"`
	Created_at time.Time `json:"created_at"`
	Enabled    bool      `json:"enabled"`

	Mascota []Mascota `gorm:"foreignKey:Cliente_id;references:ID"`
}

func (Cliente) TableName() string { return "clientes" }
