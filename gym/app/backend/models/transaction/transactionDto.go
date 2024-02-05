package transaction

import (
	"gym/app/backend/models/user"
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionDto struct {
	Id              uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	UUID            uuid.UUID
	UserId          uint
	TransactionDate time.Time
	TransactionNo   string
	User            user.UserDto
}
