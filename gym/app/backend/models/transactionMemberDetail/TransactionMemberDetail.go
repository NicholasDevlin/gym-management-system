package transactionmemberdetail

import (
	"gym/app/backend/models/user"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
 "gorm.io/plugin/soft_delete"

)

type TransactionMemberDetail struct {
	gorm.Model
	UUID                uuid.UUID
	TransactionDetailId uint
	UserId              uint
	User                user.User `gorm:"foreignKey:UserId"`
	AdditionalPrice int64
	IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
}
