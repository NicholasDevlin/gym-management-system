package absensi

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type AbsensiReq struct {
	UUID     uuid.UUID
	UserUUID uuid.UUID
	Date     time.Time
}
