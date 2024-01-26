package membershipplan

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type MembershipPlan struct {
	gorm.Model
	UUID        uuid.UUID
	Name string
	Duration int
	Description string
}