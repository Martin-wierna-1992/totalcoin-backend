package model

import (
	"time"
)

type Mascota struct {
	ID               int       `gorm:"primary_key" json:"id"`
	Cliente_id       int       `json:"cliente_id"`
	Nombre           string    `json:"nombre"`
	Peso             float32   `json:"peso"`
	Fecha_nacimiento time.Time `json:"fecha_nacimiento"`
	Castrado         bool      `json:"castrado"`
	Created_at       time.Time `json:"created_at"`
	Tipo             string    `json:"tipo"`
	Enabled          bool      `json:"enabled"`
	Combo            []Combo   `gorm:"foreignKey:mascota_id;references:ID"`
}

func (Mascota) TableName() string { return "mascotas" }
