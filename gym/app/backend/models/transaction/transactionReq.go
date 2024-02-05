package transaction

import (
	"gym/app/backend/models/user"
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionReq struct {
	Id              uint         `json:"id" form:"id"`
	UUID            uuid.UUID    `json:"uuid" form:"uuid"`
	UserId          uint         `json:"userId" form:"userId"`
	TransactionDate time.Time    `json:"transactionDate" form:"transactionDate"`
	TransactionNo   string       `json:"transactionNo" form:"transactionNo"`
	User            user.UserReq `json:"user" form:"user"`
}
