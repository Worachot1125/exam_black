package model

import (
	"github.com/uptrace/bun"
)

type Emergency_report struct {
	bun.BaseModel `bun:"table:emergency_reports"`

	ID                string `json:"id" bun:",pk,type:uuid,default:gen_random_uuid()"`
	User_ID           string `json:"user_id" bun:"user_id,notnull"`
	Emergency_Type_ID string `json:"emergency_type_id" bun:"emergency_type_id,notnull"`
	Description       string `json:"description" bun:"description,notnull"`
	Image_URL         string `json:"image_url" bun:"image_url,notnull"`
	Status            string `json:"status" bun:"status,notnull,default:'Pending'"`

	User           *User           `bun:"rel:belongs-to,join:user_id=id"`
	Emergency_Type *Emergency_Type `bun:"rel:belongs-to,join:emergency_type_id=id"`

	CreateUpdateUnixTimestamp
	SoftDelete
}
