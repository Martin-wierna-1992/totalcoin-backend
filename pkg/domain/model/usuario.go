package model

import "time"

type Usuario struct {
	ID         int       `gorm:"primary_key" json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	Created_at time.Time `json:"created_at"`
	Enabled    bool      `json:"enabled"`
	Cliente    Cliente   `gorm:"foreignKey:usuario_id;references:ID"`
	Vendedor   Vendedor  `gorm:"foreignKey:usuario_id;references:ID"`
}

func (Usuario) TableName() string { return "usuarios" }
