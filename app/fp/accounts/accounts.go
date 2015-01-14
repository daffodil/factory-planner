package accounts

import (

	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/daffodil/factory-planner/app/fp"
)


type Account struct {

	// The primary key
	AccountId int ` json:"account_id" gorm:"column:account_id; primary_key:yes"`

	AccActive *bool   ` json:"acc_active"  sql:"type:int(2)" `
	Root *bool   ` json:"root"  sql:"type:int(2)" `

	// Given name of the company eg Tesla Mirror Inc
	Company string ` json:"company" `

	// The nickname for use inside buisness eg TesMir
	Ticker string `  json:"ticker" sql:"type:varchar(25);default:''" `

	// The account reference and probably also accounts key eg sage Ref
	AccRef string `  json:"acc_ref" sql:"type:varchar(25);default:''" `



	// Flag to indicate account is on hold
	// need this as an alert system
	OnHold *bool  ` json:"on_hold" sql:"type:int(2);"`

	// An account has flags for the "type"
	IsClient *bool   ` json:"is_client" gorm:"column:is_client" sql:"type:int(2)" `
	IsSupplier *bool  ` json:"is_supplier" gorm:"column:is_supplier" sql:"type:int(2)"`
	//IsSubContracter bool

	// Client can login at website
	Online *bool  ` json:"online" ssgorm:"column:is_supplier" sql:"type:int(2)"`

	// Latest list of notes on this account
	Notice string  ` json:"notice" sql:"default:''" `
}

func (me Account) TableName() string {
	return "accounts"
}

// Adds indexes to accounts table
func DB_IndexAccount(db gorm.DB) {

	cols := []string{
		"acc_active", "company", "ticker", "acc_ref",
		"on_hold"	, "is_client", "is_supplier", "online"}

	for _, c := range cols {
		db.Model(&Account{}).AddIndex("idx_" + c, c)
	}
}


type AccountView struct {
	Account
}



func GetAccountsIndex(db gorm.DB, search_vars fp.SearchVars) ([]AccountView, error) {

	var accounts []AccountView
	fmt.Println("getttttts=", search_vars)
	//db.Find(&orders, OrderView{AccountId: account_id})
	cols := "account_id, company, ticker, acc_ref, root, acc_active, "
	cols += " on_hold, is_client, is_supplier "

	where := search_vars.GetSQL("company", "acc_active")
	fmt.Println("where=", where)
	db.Table("v_accounts").Select(cols).Where(where).Scan(&accounts)

	return accounts, nil

}
