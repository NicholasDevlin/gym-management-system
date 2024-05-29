package transactionmemberdetail

import (
	"gym/app/backend/models/user"
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionMemberDetailDto struct {
	Id                    uint
	CreatedAt             time.Time
	UpdatedAt             time.Time
	UUID                  uuid.UUID
	TransactionDetailId   uint
	TransactionDetailUUID uuid.UUID
	UserId                uint
	UserUUID              uuid.UUID
	User                  user.UserDto
}
