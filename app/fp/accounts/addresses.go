package accounts

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Address struct {

	AddressId int `json:"address_id" gorm:"column:address_id; primary_key:yes"`
	AccountId int `json:"account_id" `

	// whether this address is active or legacy for archive
	AddrActive bool `json:"addr_active" sql:"type:int(2);not null;default:0" `

	// brief descrition eg Watford HQ
	Location string `json:"location" sql:"type:varchar(100);not null;default:''" `

	// address excluding postocde
	Address string `json:"address" sql:"type:varchar(100);not null;default:''" `

	// postcode is seperate for geolocation etc
	Postcode string `json:"postcode" sql:"type:varchar(100);not null;default:''" `

	Tel string `json:"tel" sql:"type:varchar(100);not null;default:''" `
	Fax string `json:"fax" sql:"type:varchar(100);not null;default:''" `

	// flag to indicate HQ of organisation
	IsHq bool `json:"is_hq" sql:"type:int(2);not null;default:0" `

	// flag to indicate and actual physical address..
	IsAddress bool `json:"is_address" sql:"type:int(2);not null;default:0" `

	// adds address to billing options, need to be is_address = true
	IsBilling bool `json:"is_billing" sql:"type:int(2);not null;default:0" `

	// search field
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

var ADDRESS_VIEW = "v_addresses"
var ADDRESS_VIEW_COLS =  `
address_id, account_id, addr_active, location, address, postcode,
tel, fax, is_hq,  is_billing
`


func GetAccountAddresses(db gorm.DB, account_id int)([]Address, error) {


	rows := make([]Address, 0)

	//where := search_vars.GetSQL("company", "acc_active")
	//fmt.Println("where=", where)
	res := db.Table(ADDRESS_VIEW).Select(ADDRESS_VIEW_COLS).Where("account_id=?", account_id).Scan(&rows)
	fmt.Println(res)

	return rows, nil

}
