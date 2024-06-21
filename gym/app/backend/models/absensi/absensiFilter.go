package absensi

import (
	customtimeformat "gym/app/backend/utils/customTimeFormat"

	uuid "github.com/satori/go.uuid"
)

type AbsensiFilter struct {
	UUID     uuid.UUID `json:"uuid" form:"uuid" query:"uuid"`
	UserUUID uuid.UUID `json:"userUUID" form:"userUUID" query:"userUUID"`
	DateFrom customtimeformat.CustomTime `json:"dateFrom" form:"dateFrom" query:"dateFrom"`
	DateTo   customtimeformat.CustomTime `json:"dateTo" form:"dateTo" query:"dateTo"`
}
