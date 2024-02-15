package transactiondetail

import (
	membershipplan "gym/app/backend/models/membershipPlan"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type TransactionDetail struct {
	gorm.Model
	UUID             uuid.UUID
	TransactionId    uint
	MembershipPlanId uint
	Quantity         int
	MembershipPlan   membershipplan.MembershipPlan `gorm:"foreignKey:MembershipPlanId"`
}
