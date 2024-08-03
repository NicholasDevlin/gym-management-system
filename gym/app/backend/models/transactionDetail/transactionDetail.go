package transactiondetail

import (
	membershipplan "gym/app/backend/models/membershipPlan"
	transactionmemberdetail "gym/app/backend/models/transactionMemberDetail"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type TransactionDetail struct {
	gorm.Model
	UUID                    uuid.UUID
	TransactionId           uint
	MembershipPlanId        uint
	Quantity                int
	Price                   int
	MembershipPlan          membershipplan.MembershipPlan                     `gorm:"foreignKey:MembershipPlanId"`
	TransactionMemberDetail []transactionmemberdetail.TransactionMemberDetail `gorm:"foreignKey:TransactionDetailId"`
	IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
}
