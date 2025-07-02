package model

import (
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID          string `json:"id" bun:",pk,type:uuid,default:gen_random_uuid()"`
	FirstName   string `bun:"first_name,notnull"`
	LastName    string `bun:"last_name,notnull"`
	User_Number string `bun:"user_number,notnull,unique"`
	Password    string `bun:"password,notnull"`
	Phone       string `bun:"phone,notnull"`
	Address     string `bun:"address,notnull"`
	Role_ID     string `bun:"role_id,notnull"`

	Role *Role `bun:"rel:belongs-to,join:role_id=id"`

	CreateUpdateUnixTimestamp
	SoftDelete
}
