package absensi

import (
	"gym/app/backend/models/user"
	"time"

	uuid "github.com/satori/go.uuid"
)

type AbsensiRes struct {
	UUID     uuid.UUID    `json:"uuid" form:"uuid"`
	UserUUID uuid.UUID    `json:"userUUID" form:"userUUID"`
	Date     time.Time    `json:"date" form:"date"`
	User     user.UserRes `json:"user" form:"user"`
}
