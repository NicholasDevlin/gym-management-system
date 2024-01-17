package user

import (
	"time"

	"github.com/google/uuid"
)

type UserRes struct {
	Id          uint `json:"id"`
	UUID           uuid.UUID `json:"uuid" form:"uuid"`
	Email          string    `json:"email" form:"email"`
	Token       string    `json:"token" form:"token"`
	PhoneNumber    string    `json:"phoneNumber" form:"phoneNumber"`
	Gender         string    `json:"gender" form:"gender"`
	BirthDate      time.Time `json:"birthDate" form:"birthDate"`

	DisplayName    string  `json:"displayName" form:"displayName"` // to store display name
	GoogleID       *string `json:"googleId" form:"googleId"`      // to store Google ID
	ProfilePicture string  `json:"profilePicture" form:"profilePicture"` // to store profile picture URL
	IsGoogleUser   bool    `json:"isGoogleUser" form:"isGoogleUser"`
}
