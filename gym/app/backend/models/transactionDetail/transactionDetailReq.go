package transactiondetail

import (
	membershipplan "gym/app/backend/models/membershipPlan"
	"gym/app/backend/models/user"

	uuid "github.com/satori/go.uuid"
)

type TransactionDetailReq struct {
	Id                 uint                             `json:"id" form:"id"`
	UUID               uuid.UUID                        `json:"uuid" form:"uuid"`
	TransactionId      uint                             `json:"transactionId" form:"transactionId"`
	UserUUID           uuid.UUID                        `json:"userUUID" form:"userUUID"`
	MembershipPlanUUID uuid.UUID                        `json:"membershipPlanUUID" form:"membershipPlanUUID"`
	Quantity           int                              `json:"qty" form:"qty"`
	User               user.UserReq                     `json:"user" form:"user"`
	MembershipPlan     membershipplan.MembershipPlanReq `json:"membershipPlan" form:"membershipPlan"`
}
