package models


import (
	"gorm.io/gorm"
)


type Products struct {
	gorm.Model
	Nombre	string
	Descripcion *string
	Sku		string
	Cantidad	uint
	Precio_compra 	uint
	Precio_venta	uint
}