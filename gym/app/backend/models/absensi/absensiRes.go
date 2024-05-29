package absensi

import (
	"gym/app/backend/models/user"
	"time"

	uuid "github.com/satori/go.uuid"
)

type AbsensiRes struct {
	UUID   uuid.UUID
	UserUUID uuid.UUID
	Date   time.Time
	User   user.UserRes
}
