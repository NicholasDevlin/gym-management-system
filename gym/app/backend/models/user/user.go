package user

import (
	"gym/app/backend/models/role"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID        uuid.UUID
	Email       string `gorm:"unique"`
	Password    string
	PhoneNumber string `gorm:"unique"`
	Gender      string
	BirthDate   *time.Time
	SubscriptionExpirationDate *time.Time

	DisplayName    string  
	GoogleID       *string
	ProfilePicture string  
	IsGoogleUser   bool    `gorm:"default:false"`

	RoleId uint
	Role   role.Role `gorm:"foreignKey:RoleId"`
}
