package accounts

import (
	"github.com/jinzhu/gorm"
	//"github.com/revel/revel"
)


type Account struct {

	// The primary key
	AccId int ` deaddb:"account_id" json:"account_id" gorm:"column:account_id; primary_key:yes"`

	// Given name of the company eg Tesla Mirror Inc
	Company string ` deaddb:"company" json:"company" gorm:"column:company"`

	// The nickname for use inside buisness eg TesMir
	Ticker string ` deaddb:"company" json:"ticker" `

	// The account reference and probably also accounts key eg sage Ref
	AccRef string ` deaddb:"acc_ref" json:"acc_ref"`

	// a list of toop level domains for this client.. ie
	// we dont want to send outside these domain
	// and on input we can sniff out emails from these domains
	// and match to contracts
	// TODO
	Domains []string

	// Flag to indicate account is active
	AccStatus string

	// Flag to indicate account is on hold
	// need this as an alert system
	OnHold int  ` json:"on_hold" `

	// An account has flags for the "type"
	IsClient bool   ` json:"is_client" gorm:"column:client"`
	IsSupplier bool  ` json:"is_supplier" gorm:"column:supplier" `
	//IsSubContracter bool

	// Latest list of notes on this account
	Notes []string
}

func (me Account) TableName() string {
	return "accounts"
}



func AccountsIndex(db gorm.DB) ([]Account, error) {

	var accs []Account

	var q *gorm.DB
	q = db.Where("acc_active = ?", 1)
	q.Order("company asc").Find(&accs)


	return accs, nil

}
