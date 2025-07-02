package model

import (
	"github.com/uptrace/bun"
)

type Role struct {
	bun.BaseModel `bun:"table:roles"`

	ID   string `json:"id" bun:",pk,type:uuid,default:gen_random_uuid()"`
	Name string `bun:"name,notnull"`
}
