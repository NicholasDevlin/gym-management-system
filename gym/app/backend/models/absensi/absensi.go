package absensi

import (
	"time"

	"gym/app/backend/models/user"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Absensi struct {
	gorm.Model
	UUID     uuid.UUID
	UserId   uint
	Date     time.Time
	CheckOut *time.Time
	User     user.User `gorm:"foreignKey:UserId"`
}
