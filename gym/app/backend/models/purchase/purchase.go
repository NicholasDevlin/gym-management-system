package purchase

import (
	"gym/app/backend/models/user"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	UUID         uuid.UUID
	UserId       uint
	PurchaseDate time.Time
	PurchaseNo   string
	User         user.User `gorm:"foreignKey:UserId"`
}
