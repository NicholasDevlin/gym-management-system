package transaction

import (
	transactiondetail "gym/app/backend/models/transactionDetail"
	"gym/app/backend/models/user"
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionRes struct {
	UUID              uuid.UUID                                `json:"uuid" form:"uuid"`
	TransactionDate   time.Time                                `json:"transactionDate" form:"transactionDate"`
	TransactionNo     string                                   `json:"transactionNo" form:"transactionNo"`
	Status            string                                   `json:"status" form:"status"`
	UserUUID          uuid.UUID                                `json:"userUUID" form:"userUUID"`
	User              user.UserRes                             `json:"user" form:"user"`
	Total             int64                                      `json:"total" form:"total"`
	TransactionDetail []transactiondetail.TransactionDetailRes `json:"transactionDetail" form:"transactionDetail"`
}
