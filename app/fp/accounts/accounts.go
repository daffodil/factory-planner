package accounts

import (

	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/daffodil/factory-planner/app/fp"
)


type Account struct {

	// Primary key
	AccountId int ` json:"account_id" gorm:"column:account_id; primary_key:yes"`

	// Account is active
	AccActive *bool   ` json:"acc_active"  sql:"type:int(2)" `

	// Flag to indicate the global root account
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

	// This accounts users can login at website
	Online *bool  ` json:"online" sql:"type:int(2)"`

	// Important notice on this account
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





// Database view extends the Account with other stuff
type AccountView struct {
	Account
	OrdersDue int ` json:"orders_due" `
}

// Columns for select, messy I know if anyones got brighter ideas
var view_cols string = `
account_id, company, ticker, acc_ref, root, acc_active,
on_hold, is_client, is_supplier, orders_due
`

// returns search and view results
func GetAccountsIndex(db gorm.DB, search_vars fp.SearchVars) ([]AccountView, error) {

	var rows []AccountView
	fmt.Println("getttttts=", search_vars)

	where := search_vars.GetSQL("company", "acc_active")
	fmt.Println("where=", where)
	db.Table("v_accounts").Select(view_cols).Where(where).Scan(&rows)

	return rows, nil

}

// Return account by ID
func GetAccount(db gorm.DB, account_id int)(*AccountView, error) {

	fmt.Println("account_id=", account_id)
	var row *AccountView = new(AccountView)
	db.Table("v_accounts").Select(view_cols).Where("account_id = ?", account_id).Scan(row)

	return row, nil
}

// Global Company Account
var rootAccount *AccountView

// Initialise rootAccount, called on startup
func InitRoot(db gorm.DB) {
	GetRootAccount(db)
}

// Returns/Loads Global Root account
func GetRootAccount(db gorm.DB)(*AccountView, error) {
	if rootAccount == nil {
		rootAccount = new(AccountView)
		db.Table("v_accounts").Select(view_cols).Where("root = 1").Scan(rootAccount)
	}
	return rootAccount, nil
}









