package role

import "time"

type RoleDto struct {
	Id   uint
	Role string
	CreatedAt time.Time
	UpdatedAt time.Time
}
