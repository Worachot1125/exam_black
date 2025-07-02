package model

import (
	"github.com/uptrace/bun"
)

type User_Role struct {
	bun.BaseModel `bun:"table:user_roles"`

	ID          string `json:"id" bun:",pk,type:uuid,default:gen_random_uuid()"`
	User_ID	string `json:"user_id" bun:"user_id,notnull"`
	Role_ID	string `json:"role_id" bun:"role_id,notnull"`

	User *User `bun:"rel:belongs-to,join:user_id=id"`
	Role *Role `bun:"rel:belongs-to,join:role_id=id"`
}
