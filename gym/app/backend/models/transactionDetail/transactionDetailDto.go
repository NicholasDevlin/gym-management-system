package transactiondetail

import (
	membershipplan "gym/app/backend/models/membershipPlan"
	"gym/app/backend/models/user"
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionDetailDto struct {
	Id                 uint
	CreatedAt          time.Time
	UpdatedAt          time.Time
	UUID               uuid.UUID
	UserUUID           uuid.UUID
	UserId             uint
	TransactionId      uint
	MembershipPlanId   uint
	MembershipPlanUUID uuid.UUID
	Quantity           int
	User               user.UserDto
	MembershipPlan     membershipplan.MembershipPlanDto
}
