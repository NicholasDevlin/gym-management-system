package transactiondetail

import (
	membershipplan "gym/app/backend/models/membershipPlan"

	uuid "github.com/satori/go.uuid"
)

type TransactionDetailReq struct {
	Id               uint                             `json:"id" form:"id"`
	UUID             uuid.UUID                        `json:"uuid" form:"uuid"`
	TransactionId    uint                             `json:"transactionId" form:"transactionId"`
	MembershipPlanUUID uuid.UUID                              `json:"membershipPlanUUID" form:"membershipPlanUUID"`
	Quantity         int                              `json:"qty" form:"qty"`
	MembershipPlan   membershipplan.MembershipPlanReq `json:"membershipPlan" form:"membershipPlan"`
}
