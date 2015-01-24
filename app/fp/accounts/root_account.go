package accounts

import (

	//"fmt"
	"github.com/jinzhu/gorm"

	//"github.com/daffodil/factory-planner/app/fp"
)


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
		db.Table(ACCOUNT_VIEW).Select(ACCOUNT_VIEW_COLS).Where("root = 1").Scan(rootAccount)
	}
	return rootAccount, nil
}
