package transactiondetail

import (
	membershipplan "gym/app/backend/models/membershipPlan"

	uuid "github.com/satori/go.uuid"
)

type TransactionDetailRes struct {
	UUID           uuid.UUID                        `json:"uuid" form:"uuid"`
	Quantity       int                              `json:"qty" form:"qty"`
	MembershipPlan membershipplan.MembershipPlanReq `json:"membershipPlan" form:"membershipPlan"`
}
