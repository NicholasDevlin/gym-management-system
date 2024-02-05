package migrations

import (
	membershipplan "gym/app/backend/models/membershipPlan"
	"gym/app/backend/models/role"
	"gym/app/backend/models/transaction"
	"gym/app/backend/models/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&role.Role{})
	db.AutoMigrate(&membershipplan.MembershipPlan{})
	db.AutoMigrate(&transaction.Transaction{})
}
