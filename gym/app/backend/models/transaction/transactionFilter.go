package transaction

import (
	customtimeformat "gym/app/backend/utils/customTimeFormat"

	uuid "github.com/satori/go.uuid"
)

type TransactionFilter struct {
	TransactionDateFrom customtimeformat.CustomTime `json:"transactionDateFrom" form:"transactionDateFrom" query:"transactionDateFrom"`
	TransactionDateTo   customtimeformat.CustomTime `json:"transactionDateTo" form:"transactionDateTo" query:"transactionDateTo"`
	MemberUUID          uuid.UUID                   `json:"memberUUID" form:"memberUUID" query:"memberUUID"`
	TransactionUUID     uuid.UUID                   `json:"transactionUUID" form:"transactionUUID" query:"transactionUUID"`
	TransactionId       uint                        `json:"transactionId" form:"transactionId" query:"transactionId"`
	UserUUID            uuid.UUID                   `json:"userUUID" form:"userUUID" query:"userUUID"`
	Status              string                      `json:"status" form:"status" query:"status"`
	IsComplete          bool                        `json:"isComplete" form:"isComplete" query:"isComplete"`
	TransactionNo       string                      `json:"transactionNo" form:"transactionNo" query:"transactionNo"`
	UUID                uuid.UUID                   `json:"transactionUUID" form:"transactionUUID" query:"transactionUUID"`
	MembershipUUID      uuid.UUID                   `json:"MmembershipUUID" form:"MmembershipUUID" query:"MmembershipUUID"`
}
