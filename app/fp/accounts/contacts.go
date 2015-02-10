package accounts

import (

	//"fmt"

	"github.com/jinzhu/gorm"

	"github.com/daffodil/factory-planner/app/fp"
)

type Contact struct {

	ContactId int ` json:"contact_id"  gorm:"primary_key:yes" `
	AccountId int ` json:"account_id" `
	AddressId int ` json:"address_id" `

	Contact string 	` json:"contact" sql:"type:varchar(100);" `
	Title string 	` json:"title" sql:"type:varchar(100);default:''" `
	Email string	 ` json:"email" sql:"type:varchar(100);default:''" `

	SecurityId int 	` json:"security_id" sql:"type:int(2);not null;" `
	Syslogin *string ` json:"syslogin" sql:"type:varchar(20);" `
	Secret *string 	` json:"secret" sql:"type:varchar(100)" `

	ConActive *bool	` json:"con_active" sql:"type:int(2);" `
	CanLogin *bool 	` json:"can_login" sql:"type:int(2);" `
	PassChange *bool ` json:"pass_change" sql:"type:int(2);" `
	OnlineStatus *int ` json:"online_status" sql:"type:int(2);" `

	DirectLine *string ` json:"direct_line" sql:"type:varchar(100);not null;default:''" `
	Mobile *string 	 ` json:"mobile" sql:"type:varchar(100);not null;default:''" `


	ConNotes *string ` json:"con_notes" sql:"type:varchar(100);not null;default:''" `
	WwwPage *bool    ` json:"www_page" sql:"type:int(2)" `

	ConUid *string 	` json:"con_uid" sql:"type:varchar(255);default:''" `
	ConSearch *string ` json:"con_search" sql:"type:varchar(255);default:''" `
}

func (me Contact) TableName() string {
	return "contacts"
}


func DB_IndexContact(db gorm.DB) {

	cols := []string{
		"account_id", "address_id", "contact", "title",
		"email"	, "security_id",
		"syslogin", "con_active", "can_login", "con_search", "www_page"}

	for _, c := range cols {
		db.Model(&Contact{}).AddIndex("idx_" + c, c)
	}
}


// Database view extends the contact with other stuff
type ContactView struct {
	Contact
	Company string ` json:"company" `
	Ticker string ` json:"ticker" `
	AccActive string ` json:"acc_active" `
	Root *bool ` json:"root" `
	Online *bool ` json:"online" `
}

var CONTACT_VIEW = "v_contacts"
var CONTACT_VIEW_COLS string = `
account_id, company, ticker, acc_ref, root, acc_active, online, address_id,
contact_id, contact, mobile, email, direct_line, title, con_active,
can_login, security_id, security, syslogin, www_page
`


func SearchContacts(db gorm.DB, search_vars fp.SearchVars) ([]ContactView, error){

	rows := make([]ContactView, 0)

	where := search_vars.GetSQL("contact", "acc_active")

	res := db.Table(CONTACT_VIEW).Select(CONTACT_VIEW_COLS).Where(where).Scan(&rows)
	if res.Error != nil {
		return rows, res.Error
	}
	return rows, nil

}


func GetAccountContacts(db gorm.DB, account_id int) ([]ContactView, error){

	rows := make([]ContactView, 0)

	res := db.Table(CONTACT_VIEW).Select(CONTACT_VIEW_COLS).Where("account_id=?", account_id).Scan(&rows)
	if res.Error != nil {
		return rows, res.Error
	}
	return rows, nil

}

func GetStaff(db gorm.DB)([]ContactView, error){


	return GetAccountContacts(db, rootAccount.AccountId)
}

func GetContact(db gorm.DB, contact_id int) (ContactView, error){

	var row ContactView

	res := db.Table(CONTACT_VIEW).Select(CONTACT_VIEW_COLS).Where("contact_id=?", contact_id).Scan(&row)
	if res.Error != nil {
		return row, res.Error
	}

	return row, nil

}
