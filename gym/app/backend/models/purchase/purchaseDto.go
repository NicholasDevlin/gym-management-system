package purchase

import (
	"gym/app/backend/models/user"
	"time"

	uuid "github.com/satori/go.uuid"
)

type PurchaseDto struct {
	Id uint
	CreatedAt time.Time
	UpdatedAt time.Time
	UUID         uuid.UUID
	UserId       uint
	PurchaseDate time.Time
	PurchaseNo   string
	User user.UserDto
}