package user

import (
	"time"
)

type UserRes struct {
	Id          string
	Email       string
	Token       string
	PhoneNumber string
	Gender      string
	BirthDate   time.Time

	DisplayName    string // to store display name
	GoogleID       *string // to store Google ID
	ProfilePicture string // to store profile picture URL
	IsGoogleUser   bool
}
