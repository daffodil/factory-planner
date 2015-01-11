package accounts

import (
	_ "github.com/jinzhu/gorm"
	//"github.com/revel/revel"
)

type Contact struct {


	ContactId int ` json:"contact_id"  gorm:"primary_key:yes" `
	AccountId int ` json:"account_id" `
	AddressId int ` json:"location_id" `


	Contact string 	` json:"contact" sql:"type:varchar(100);" `
	Title string 	` json:"title" sql:"type:varchar(100);default:''" `
	Email string	 ` json:"email" sql:"type:varchar(100);default:''" `

	SecurityId int 	` json:"security_id" sql:"type:int(2);not null" `
	Syslogin string ` json:"syslogin" sql:"type:int(2);not null" `
	Secret string 	` json:"secret" sql:"type:varchar(100);default:''" `
	ConActive bool	` json:"con_active" sql:"type:int(2);not null;default:0" `
	CanLogin bool 	` json:"can_login" sql:"type:int(2);not null;default:0" `
	PassChange bool ` json:"pass_enforced" sql:"type:int(2);not null;default:0" `
	OnlineStatus int ` json:"online_status" sql:"type:int(2);not null";default:0 `

	DirectLine string ` json:"direct_line" sql:"type:varchar(100);default:''" `
	Mobile string 	 ` json:"mobile" sql:"type:varchar(100);default:''" `


	ConNotes string ` json:"con_notes" sql:"type:varchar(100);default:''" `
	WwwPage bool    ` json:"www_page" sql:"type:int(2);not null" `

	ConUid string 	` json:"c_uid" sql:"type:varchar(255);default:''" `
	ConSearch string ` json:"con_search" sql:"type:varchar(255);default:''" `

}
func (me Contact) TableName() string {
	return "contacts"
}
