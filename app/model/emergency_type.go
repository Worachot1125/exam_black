package model

import (
	"github.com/uptrace/bun"
)

type Emergency_Type struct {
	bun.BaseModel `bun:"table:emergency_types"`

	ID   string `json:"id" bun:",pk,type:uuid,default:gen_random_uuid()"`
	Name string `json:"name" bun:"name,notnull"`
}
