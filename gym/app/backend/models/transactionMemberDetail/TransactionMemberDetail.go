package transactionmemberdetail

import (
	transactiondetail "gym/app/backend/models/transactionDetail"
	"gym/app/backend/models/user"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type TransactionMemberDetail struct {
	gorm.Model
	UUID                uuid.UUID
	TransactionDetailId uint
	UserId              uint
	User                user.User
	TransactionDetail   transactiondetail.TransactionDetail
}
