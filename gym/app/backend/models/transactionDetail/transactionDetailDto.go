package transactiondetail

import (
	membershipplan "gym/app/backend/models/membershipPlan"
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
	Subtotal           int64
	MembershipPlan membershipplan.MembershipPlanDto
}
