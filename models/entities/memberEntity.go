package entities

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	Id_member uint              `gorm:"primaryKey; autoIncrement" json:"id_member"`
	Username  string            `gorm:"type:varchar(255); uniqueIndex; not null" json:"username"`
	Gender    string            `gorm:"type:varchar(255)" json:"gender"`
	Skintype  string            `gorm:"type:varchar(255)" json:"skintype"`
	Skincolor string            `gorm:"type:varchar(255)" json:"skincolor"`
	Reviews   []*Review_product `gorm:"many2many:like_review" json:"reviews"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
