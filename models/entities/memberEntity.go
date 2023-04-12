package entities

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	Id_member uint   `gorm:"primaryKey; autoIncrement" json:"id_member"`
	Username  string `gorm:"uniqueIndex" json:"username"`
	Gender    string `json:"gender"`
	Skintype  string `json:"skintype"`
	Skincolor string `json:"skincolor"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
