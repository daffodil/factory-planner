
package parts

import (


)


type Part struct {

	PartId int ` json:"part_id" gorm:"column:part_id; primary_key:yes"`
	AccountId int ` json:"account_id" sql:"not null" `
	ClientRef string `json:"client_ref" sql:"type:varchar(100);not null;default:''" `
	OurRef string `json:"our_ref" sql:"type:varchar(100);not null;default:''" `
	Model  string `json:"model" sql:"type:varchar(25);not null;default:''" `
	Description string `json:"description" sql:"type:varchar(255);not null;default:''"`

}

func (me Part) TableName() string {
	return "parts"
}
