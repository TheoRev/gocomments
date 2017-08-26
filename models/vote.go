package models

import (
	"github.com/jinzhu/gorm"
)

// Vote estructura  de votos
type Vote struct {
	gorm.Model
	CommentID uint `json:"commentId" gorm:"not null"`
	UserID    uint `json:"userId" gorm:"not null"`
	Value     bool `json:"value" gorm:"not null"`
}
