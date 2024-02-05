package transaction

import (
	transactiondetail "gym/app/backend/models/transactionDetail"
	"gym/app/backend/models/user"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UUID              uuid.UUID
	UserId            uint
	TransactionDate   time.Time
	TransactionNo     string
	Status            string
	User              user.User                             `gorm:"foreignKey:UserId"`
	TransactionDetail []transactiondetail.TransactionDetail `gorm:"foreignKey:TransactionId"`
}
