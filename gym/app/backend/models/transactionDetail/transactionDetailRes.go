package transactiondetail

import (
	membershipplan "gym/app/backend/models/membershipPlan"
	"gym/app/backend/models/user"

	uuid "github.com/satori/go.uuid"
)

type TransactionDetailRes struct {
	UUID               uuid.UUID                        `json:"uuid" form:"uuid"`
	Quantity           int                              `json:"qty" form:"qty"`
	Subtotal           int64                            `json:"subtotal" form:"subtotal"`
	UserUUID           uuid.UUID                        `json:"userUUID" form:"userUUID"`
	MembershipPlanUUID uuid.UUID                        `json:"membershipPlanUUID" form:"membershipPlanUUID"`
	User               user.UserRes                     `json:"user" form:"user"`
	MembershipPlan     membershipplan.MembershipPlanRes `json:"membershipPlan" form:"membershipPlan"`
}
