package transactiondetail

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionDetailDto struct {
	Id uint
	CreatedAt time.Time
	UpdatedAt time.Time
	UUID uuid.UUID
	TransactionId uint
	MembershipPlanId uint
	Quantity int
}