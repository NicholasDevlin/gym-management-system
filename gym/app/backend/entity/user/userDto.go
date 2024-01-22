package user

import (
	"gym/app/backend/entity/role"
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserDto struct {
	Id          uint
	UUID        uuid.UUID
	Email       string
	Password    string
	PhoneNumber string
	Gender      string
	BirthDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time

	DisplayName    string  // to store display name
	GoogleID       *string // to store Google ID
	ProfilePicture string  // to store profile picture URL
	IsGoogleUser   bool

	RoleId uint
	Role   role.RoleDto
}
