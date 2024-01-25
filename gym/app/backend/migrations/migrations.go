package migrations

import (
	"gym/app/backend/models/role"
	"gym/app/backend/models/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&role.Role{})
}
