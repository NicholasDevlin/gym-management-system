package transactiondetail

import (
	membershipplan "gym/app/backend/models/membershipPlan"

	uuid "github.com/satori/go.uuid"
)

type TransactionDetailRes struct {
	UUID           uuid.UUID                        `json:"uuid" form:"uuid"`
	Quantity       int                              `json:"qty" form:"qty"`
	Subtotal       int64                              `json:"subtotal" form:"subtotal"`
	MembershipPlanUUID uuid.UUID                              `json:"membershipPlanUUID" form:"membershipPlanUUID"`
	MembershipPlan membershipplan.MembershipPlanRes `json:"membershipPlan" form:"membershipPlan"`
}
