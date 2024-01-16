package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email       string    `gorm:"unique;"`
	Password    string
	PhoneNumber string `gorm:"unique"`
	Gender      string
	BirthDate   time.Time

	DisplayName    string  // to store display name
	GoogleID       *string // to store Google ID
	ProfilePicture string  // to store profile picture URL
	IsGoogleUser   bool    `gorm:"default:false"`
}
