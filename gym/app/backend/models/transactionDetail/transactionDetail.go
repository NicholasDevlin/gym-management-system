package transactiondetail

import (
	membershipplan "gym/app/backend/models/membershipPlan"
	"gym/app/backend/models/user"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type TransactionDetail struct {
	gorm.Model
	UUID             uuid.UUID
	TransactionId    uint
	MembershipPlanId uint
	UserId           uint
	Quantity         int
	User             user.User
	MembershipPlan   membershipplan.MembershipPlan `gorm:"foreignKey:MembershipPlanId"`
}
