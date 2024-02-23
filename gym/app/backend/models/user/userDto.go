package user

import (
	"gym/app/backend/models/role"
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
	BirthDate   *time.Time
	SubscriptionDueDate   *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time

	DisplayName    string 
	GoogleID       *string
	ProfilePicture string 
	IsGoogleUser   bool

	RoleId uint
	Role   role.RoleDto
}
