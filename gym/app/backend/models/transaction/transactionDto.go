package transaction

import (
	transactiondetail "gym/app/backend/models/transactionDetail"
	"gym/app/backend/models/user"
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionDto struct {
	Id                uint
	CreatedAt         time.Time
	UpdatedAt         time.Time
	UUID              uuid.UUID
	UserId            uint
	TransactionDate   time.Time
	TransactionNo     string
	Status            string
	User              user.UserDto
	TransactionDetail []transactiondetail.TransactionDetailDto
}
