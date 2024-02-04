package purchase

import (
	"gym/app/backend/models/user"
	"time"

	uuid "github.com/satori/go.uuid"
)

type PurchaseReq struct {
	Id uint `json:"id" form:"id"`
	UUID         uuid.UUID `json:"uuid" form:"uuid"`
	UserId       uint `json:"userId" form:"userId"`
	PurchaseDate time.Time  `json:"purchaseDate" form:"purchaseDate"`
	PurchaseNo   string `json:"purchaseNo" form:"purchaseNo"`
	User user.UserReq `json:"user" form:"user"`
}