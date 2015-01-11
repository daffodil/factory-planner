

package dev

import (
	//"os"
	//"path/filepath"
	//"io/ioutil"
	"fmt"
	//"github.com/revel/revel"
	//"database/sql"
	"github.com/jinzhu/gorm"

	"github.com/daffodil/factory-planner/mods/accounts"
	"github.com/daffodil/factory-planner/mods/parts"
)


func DB_CreateTables(db gorm.DB) (interface{}, error) {

	foo := make( map[string]interface{} )

	s := new(accounts.Contact)
	fmt.Println(s)
	db.AutoMigrate( &accounts.Account{}, &accounts.Contact{} )

	db.AutoMigrate( &parts.Part{} )
	return foo, nil
}
