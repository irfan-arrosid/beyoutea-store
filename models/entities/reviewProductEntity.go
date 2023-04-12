package entities

import (
	"time"

	"gorm.io/gorm"
)

type Review_product struct {
	Id_review   uint `gorm:"primaryKey;autoIncrement" json:"id_review"`
	ProductID   uint `gorm:"not null" json:"id_product"`
	MemberID    uint `gorm:"not null" json:"id_member"`
	Desc_review string
	Members     []*Member `gorm:"many2many:like_review" json:"members"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
