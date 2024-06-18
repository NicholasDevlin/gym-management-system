package user

import (
	customtimeformat "gym/app/backend/utils/customTimeFormat"

	uuid "github.com/satori/go.uuid"
)

type UserFilter struct {
	Id                         uint                        `json:"id" form:"id" query:"id"`
	UUID                       uuid.UUID                   `json:"uuid" form:"uuid" query:"uuid"`
	Email                      string                      `json:"email" form:"email" query:"email"`
	PhoneNumber                string                      `json:"phoneNumber" form:"phoneNumber" query:"phoneNumber"`
	Gender                     string                      `json:"gender" form:"gender" query:"gender"`
	SubscriptionExpirationDate customtimeformat.CustomTime `json:"subscriptionDueDate" form:"subscriptionDueDate" query:"subscriptionDueDate"`
	Active                     *bool                        `json:"active" form:"active" query:"active"`
	DisplayName                string                      `json:"name" form:"name" query:"name"`
	RoleUUID                   uuid.UUID                   `json:"roleUUID" form:"roleUUID" query:"roleUUID"`
	Role                       string                      `json:"role" form:"role" query:"role"`
	LastDayActive                     bool                        `json:"lastDayActive" form:"lastDayActive" query:"lastDayActive"`  
}
