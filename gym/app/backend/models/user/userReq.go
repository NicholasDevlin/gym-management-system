package user

import (
	"gym/app/backend/models/role"
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserReq struct {
	Id                  uint       `json:"id" form:"id"`
	UUID                uuid.UUID  `json:"uuid" form:"uuid"`
	Email               string     `json:"email" form:"email"`
	Password            string     `json:"password" form:"password"`
	PhoneNumber         string     `json:"phoneNumber" form:"phoneNumber"`
	Gender              string     `json:"gender" form:"gender"`
	BirthDate           *time.Time `json:"birthDate" form:"birthDate"`
	SubscriptionExpirationDate *time.Time `json:"subscriptionDueDate" form:"subscriptionDueDate"`

	DisplayName    string  `json:"name" form:"name"`                 
	GoogleID       *string `json:"googleId" form:"googleId"`         
	ProfilePicture string  `json:"profilePicture" form:"profilePicture"` 
	IsGoogleUser   bool    `json:"isGoogleUser" form:"isGoogleUser"`

	RoleId uint         `json:"roleId" form:"roleId"`
	Role   role.RoleReq `json:"role" form:"role"`
}
