package membershipplan

import uuid "github.com/satori/go.uuid"

type MembershipPlanRes struct {
	Id uint 	`json:"id" form:"id"`
	UUID        uuid.UUID `json:"uuid" form:"uuid"`
	Name string `json:"name" form:"name"`
	Duration int `json:"duration" form:"duration"`
	Description string `json:"description" form:"description"`
	Price int64 `json:"price" form:"price"`
}