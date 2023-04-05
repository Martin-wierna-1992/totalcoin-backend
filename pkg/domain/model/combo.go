package model

import "time"

type Combo struct {
	ID           int       `gorm:"primary_key" json:"id"`
	Peso         float32   `json:"peso"`
	Complementos int       `json:"complementos"`
	Mascota_id   int       `json:"mascota_id"`
	Vendedor_id  int       `json:"vendedor_id"`
	Created_at   time.Time `json:"created_at"`
	Enabled      bool      `json:"enabled"`
	Mascota      Mascota
	Vendedor     Vendedor
}

func (Combo) TableName() string { return "combos" }
