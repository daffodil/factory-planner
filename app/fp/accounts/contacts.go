package accounts

import (
	"github.com/jinzhu/gorm"

)

type Contact struct {


	ContactId int ` json:"contact_id"  gorm:"primary_key:yes" `
	AccountId int ` json:"account_id" `
	AddressId int ` json:"address_id" `


	Contact string 	` json:"contact" sql:"type:varchar(100);" `
	Title string 	` json:"title" sql:"type:varchar(100);default:''" `
	Email string	 ` json:"email" sql:"type:varchar(100);default:''" `

	SecurityId int 	` json:"security_id" sql:"type:int(2);not null" `
	Syslogin string ` json:"syslogin" sql:"type:varchar(20);" `
	Secret string 	` json:"secret" sql:"type:varchar(100)" `
	ConActive bool	` json:"con_active" sql:"type:int(2)" `
	CanLogin bool 	` json:"can_login" sql:"type:int(2)" `
	PassChange bool ` json:"pass_change" sql:"type:int(2)" `
	OnlineStatus int ` json:"online_status" sql:"type:int(2)" `

	DirectLine string ` json:"direct_line" sql:"type:varchar(100);default:''" `
	Mobile string 	 ` json:"mobile" sql:"type:varchar(100);default:''" `


	ConNotes string ` json:"con_notes" sql:"type:varchar(100);default:''" `
	WwwPage bool    ` json:"www_page" sql:"type:int(2)" `

	ConUid string 	` json:"con_uid" sql:"type:varchar(255);default:''" `
	ConSearch string ` json:"con_search" sql:"type:varchar(255);default:''" `

}
func (me Contact) TableName() string {
	return "contacts"
}

func DB_IndexContact(db gorm.DB) {

	cols := []string{
		"account_id", "address_id", "contact", "title",
		"email"	, "security_id", "syslogin", "con_active", "can_login", "con_search", "www_page"}

	for _, c := range cols {
		db.Model(&Contact{}).AddIndex("idx_" + c, c)
	}
}