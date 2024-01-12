package migrations

import (
	"gym/app/backend/entity/role"
	"gym/app/backend/entity/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&role.Role{})
}
