package transactiondetail

import (
	membershipplan "gym/app/backend/models/membershipPlan"
	transactionmemberdetail "gym/app/backend/models/transactionMemberDetail"
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionDetailDto struct {
	Id                 uint
	CreatedAt          time.Time
	UpdatedAt          time.Time
	UUID               uuid.UUID
	TransactionId      uint
	MembershipPlanId   uint
	MembershipPlanUUID uuid.UUID
	Quantity           int
	MembershipPlan     membershipplan.MembershipPlanDto
	TransactionMemberDetail []transactionmemberdetail.TransactionMemberDetailDto
}
