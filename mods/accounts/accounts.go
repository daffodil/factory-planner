package accounts

import (
	"github.com/jinzhu/gorm"
	//"github.com/revel/revel"
)


type Account struct {

	// The primary key
	AccId int ` deaddb:"account_id" json:"account_id" gorm:"column:account_id; primary_key:yes"`

	AccActive bool   ` json:"acc_active" sssgorm:"column:is_client" sql:"type:int(2);default:0" `
	Root bool   ` ssjson:"acc_active" sssgorm:"column:is_client" sql:"type:int(2);default:0" `

	// Given name of the company eg Tesla Mirror Inc
	Company string ` json:"company" `

	// The nickname for use inside buisness eg TesMir
	Ticker string `  json:"ticker" sql:"type:varchar(25);default:''" `

	// The account reference and probably also accounts key eg sage Ref
	AccRef string `  json:"acc_ref" sql:"type:varchar(25);default:''" `

	// a list of toop level domains for this client.. ie
	// we dont want to send outside these domain
	// and on input we can sniff out emails from these domains
	// and match to contracts
	// TODO
	////Domains []string

	// Flag to indicate account is active
	//AccStatus string

	// Flag to indicate account is on hold
	// need this as an alert system
	OnHold int  ` json:"on_hold" sql:"type:int(2);default:0"`

	// An account has flags for the "type"
	IsClient bool   ` json:"is_client" gorm:"column:is_client" sql:"type:int(2);default:0" `
	IsSupplier bool  ` json:"is_supplier" gorm:"column:is_supplier" sql:"type:int(2);default:0"`
	//IsSubContracter bool

	// Client can login at website
	Online bool  ` json:"online" ssgorm:"column:is_supplier" sql:"type:int(2);default:0"`

	// Latest list of notes on this account
	Notice string  ` json:"notice" sql:"default:''" `
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
