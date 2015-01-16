package roles

import (
	"time"
	//"github.com/jinzhu/gorm"
)



type Role struct {

	// table ID
	RoleId int `json:"role_id" gorm:"column:role_id; primary_key:yes"`

	// umm
	Role string `json:"role" `
	// lookup key = "welder"
	RoleKey int `json:"role_key" sql:"type:varchar(50);not null;" `

	// What interface shows
	RoleLabel string `json:"role_label" sql:"type:varchar(50);not null;" `
	RoleDescription *string `json:"role_label" sql:"type:varchar(255);not null;" `

}
