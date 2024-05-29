package absensi

import (
	"gym/app/backend/models/user"
	"time"

	uuid "github.com/satori/go.uuid"
)

type AbsensiDto struct {
	Id        uint
	UUID      uuid.UUID
	UserUUID  uuid.UUID
	UserId    uint
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.UserDto
}
