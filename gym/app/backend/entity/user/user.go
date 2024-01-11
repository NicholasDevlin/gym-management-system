package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string `gorm:"unique;"`
	Password    string
	PhoneNumber string `gorm:"unique"`
	Gender      string
	BirthDate   time.Time

	DisplayName    string  // to store display name
	GoogleID       *string // to store Google ID
	ProfilePicture string  // to store profile picture URL
	IsGoogleUser   bool    `gorm:"default:false"`
}
