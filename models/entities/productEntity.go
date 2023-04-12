package entities

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id_product   uint   `gorm:"primaryKey;autoIncrement" json:"id_product"`
	Name_product string `json:"name_product"`
	Price        int    `json:"price"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
