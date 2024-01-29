package membershipplan

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type MembershipPlanDto struct {
	Id uint
	CreatedAt time.Time
	UpdatedAt time.Time
	UUID        uuid.UUID
	Name string
	Duration int
	Description string
	Price int64
}