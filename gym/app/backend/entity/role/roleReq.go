package role

type RoleReq struct {
	Id   uint   `json:"id" form:"id"`
	Role string `json:"role" form:"role"`
}
