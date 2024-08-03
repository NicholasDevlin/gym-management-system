package absensi

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type AbsensiReq struct {
	UUID     uuid.UUID `json:"uuid" form:"uuid"`
	UserUUID uuid.UUID `json:"userUUID" form:"userUUID"`
	Date     time.Time `json:"date" form:"date"`
	CheckOut *time.Time `json:"checkOut" form:"checkOut"`
}
