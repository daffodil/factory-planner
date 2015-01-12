

package dev

import (
	//"os"
	//"path/filepath"
	//"io/ioutil"
	//"fmt"
	//"github.com/revel/revel"
	//"database/sql"
	"github.com/jinzhu/gorm"

	"github.com/daffodil/factory-planner/mods/fpsys"
	"github.com/daffodil/factory-planner/mods/accounts"
	"github.com/daffodil/factory-planner/mods/orders"
	"github.com/daffodil/factory-planner/mods/parts"
	"github.com/daffodil/factory-planner/mods/schedule"

)


func DB_CreateTables(db gorm.DB, drop_first bool) (interface{}, error) {

	foo := make( map[string]interface{} )

	if drop_first {
		db.DropTableIfExists(&fpsys.Security{})
		db.DropTableIfExists(&fpsys.Setting{})

		db.DropTableIfExists(&accounts.Account{})
		db.DropTableIfExists(&accounts.Address{})
		db.DropTableIfExists(&accounts.Contact{})

		db.DropTableIfExists(&orders.OrderType{})
		db.DropTableIfExists(&orders.Order{})
		db.DropTableIfExists(&orders.WorkOrder{})

		db.DropTableIfExists(&schedule.WorkSchedule{})
		db.DropTableIfExists(&parts.Part{})

	}


	db.AutoMigrate( &fpsys.Setting{} )
	db.AutoMigrate( &fpsys.Security{} )
	if true == true {
		db.AutoMigrate(&accounts.Account{})
		accounts.DB_IndexAccount(db)
		db.AutoMigrate(&accounts.Address{})
		accounts.DB_IndexAddress(db)
		db.AutoMigrate(&accounts.Contact{})
		accounts.DB_IndexContact(db)


		db.AutoMigrate(&orders.OrderType{})
		db.AutoMigrate(&orders.Order{})
		db.AutoMigrate(&orders.WorkOrder{})
		orders.DB_IndexOrderType(db)
		orders.DB_IndexOrder(db)
		orders.DB_IndexWorkOrder(db)

		db.AutoMigrate(&schedule.WorkSchedule{})

		db.AutoMigrate(&parts.Part{})
	}
	return foo, nil
}
