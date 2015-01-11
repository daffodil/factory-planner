

package dev

import (
	//"os"
	//"path/filepath"
	//"io/ioutil"
	//"fmt"
	//"github.com/revel/revel"
	//"database/sql"
	"github.com/jinzhu/gorm"

	"github.com/daffodil/factory-planner/mods/accounts"
	"github.com/daffodil/factory-planner/mods/orders"
	"github.com/daffodil/factory-planner/mods/parts"
	"github.com/daffodil/factory-planner/mods/schedule"
)


func DB_CreateTables(db gorm.DB) (interface{}, error) {

	foo := make( map[string]interface{} )



	db.AutoMigrate( &accounts.Account{}, &accounts.Address{}, &accounts.Contact{} )

	db.AutoMigrate( &orders.OrderType{}, &orders.Order{}, &orders.WorkOrder{})

	db.AutoMigrate( &schedule.WorkSchedule{})

	db.AutoMigrate( &parts.Part{}, &parts.Contact2Part{} )
	return foo, nil
}
