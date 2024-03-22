package transactiondetail

import (
	membershipplan "gym/app/backend/models/membershipPlan"
	transactionmemberdetail "gym/app/backend/models/transactionMemberDetail"

	uuid "github.com/satori/go.uuid"
)

type TransactionDetailReq struct {
	UUID                       uuid.UUID                                          `json:"uuid" form:"uuid"`
	TransactionId              uint                                               `json:"transactionId" form:"transactionId"`
	MembershipPlanUUID         uuid.UUID                                          `json:"membershipPlanUUID" form:"membershipPlanUUID"`
	Quantity                   int                                                `json:"qty" form:"qty"`
	MembershipPlan             membershipplan.MembershipPlanReq                   `json:"membershipPlan" form:"membershipPlan"`
	TransactionMemberDetail []transactionmemberdetail.TransactionMemberDetailReq `json:"transactionMemberDetail" form:"transactionMemberDetail"`
}
