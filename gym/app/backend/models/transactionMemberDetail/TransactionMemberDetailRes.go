package transactionmemberdetail

import (
	"gym/app/backend/models/user"

	uuid "github.com/satori/go.uuid"
)

type TransactionMemberDetailRes struct {
	UUID                  uuid.UUID    `json:"uuid" form:"uuid"`
	TransactionDetailUUID uuid.UUID    `json:"transactionDetailUUID" form:"transactionDetailUUID"`
	UserUUID              uuid.UUID    `json:"userUUID" form:"userUUID"`
	User                  user.UserRes `json:"user" form:"user"`
}
