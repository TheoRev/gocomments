package migration

import (
	"github.com/TheoRev/gocomments/configuration"
	"github.com/TheoRev/gocomments/models"
)

// Migrate crea las tablas del modelo de db
func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()

	db.CreateTable(&models.User{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Vote{})
	db.Model(&models.Vote{}).AddUniqueIndex("commentId_userId_unique", "commentId", "userId")
}
