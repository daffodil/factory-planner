

package dev

import (
	//"os"
	//"path/filepath"
	//"io/ioutil"
	//"fmt"
	//"github.com/revel/revel"
	//"database/sql"
	"github.com/jinzhu/gorm"

	"github.com/daffodil/factory-planner/app/fp/fpsys"
	"github.com/daffodil/factory-planner/app/fp/accounts"
	"github.com/daffodil/factory-planner/app/fp/orders"
	"github.com/daffodil/factory-planner/app/fp/parts"
	"github.com/daffodil/factory-planner/app/fp/schedule"

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
		db.DropTableIfExists(&parts.Contact2Part{})

	}


	db.AutoMigrate( &fpsys.Setting{} )
	fpsys.DB_IndexSetting(db)
	db.AutoMigrate( &fpsys.Security{} )
	fpsys.DB_IndexSecurity(db)
	db.AutoMigrate( &fpsys.SysLog{} )
	fpsys.DB_IndexSysLog(db)
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

		db.AutoMigrate(&parts.Part{}, &parts.Contact2Part{})
	}
	return foo, nil
}


