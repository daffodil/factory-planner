package accounts

import (
	"github.com/jinzhu/gorm"
)

type Address struct {

	AddressId int `json:"address_id" gorm:"column:address_id; primary_key:yes"`
	AccountId int `json:"account_id" `
	AddrActive bool `json:"addr_active" sql:"type:int(2);not null;default:0" `

	Location string `json:"location" sql:"type:varchar(100);not null;default:''" `
	Address string `json:"address" sql:"type:varchar(100);not null;default:''" `
	Postcode string `json:"postcode" sql:"type:varchar(100);not null;default:''" `

	Tel string `json:"tel" sql:"type:varchar(100);not null;default:''" `
	Fax string `json:"fax" sql:"type:varchar(100);not null;default:''" `

	IsHq bool `json:"is_hq" sql:"type:int(2);not null;default:0" `
	IsAddress bool `json:"is_address" sql:"type:int(2);not null;default:0" `
	IsBilling bool `json:"is_billing" sql:"type:int(2);not null;default:0" `

	AddrSearch string `json:"addr_search" sql:"type:varchar(100);not null;default:''" `
}

func (me Address) TableName() string {
	return "addresses"
}

func DB_IndexAddress(db gorm.DB) {

	cols := []string{
		"account_id", "location", "address", "postcode",
		"addr_search"	, "addr_active", "is_hq", "is_address", "is_billing" }

	for _, c := range cols {
		db.Model(&Address{}).AddIndex("idx_" + c, c)
	}
}
