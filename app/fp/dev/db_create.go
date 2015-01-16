

package dev

import (
	//"os"
	//"path/filepath"
	//"io/ioutil"
	//"fmt"
	//"github.com/revel/revel"
	//"database/sql"
	"github.com/jinzhu/gorm"

	"github.com/daffodil/factory-planner/app/fp"
	"github.com/daffodil/factory-planner/app/fp/accounts"
	"github.com/daffodil/factory-planner/app/fp/orders"
	"github.com/daffodil/factory-planner/app/fp/parts"
	"github.com/daffodil/factory-planner/app/fp/schedule"
	"github.com/daffodil/factory-planner/app/fp/jobs"

)


func DB_CreateTables(db gorm.DB, drop_first bool) (interface{}, error) {

	foo := make( map[string]interface{} )

	if drop_first {
		db.DropTableIfExists(&fp.Security{})
		db.DropTableIfExists(&fp.Setting{})

		db.DropTableIfExists(&accounts.Account{})
		db.DropTableIfExists(&accounts.Address{})
		db.DropTableIfExists(&accounts.Contact{})

		db.DropTableIfExists(&orders.OrderType{})
		db.DropTableIfExists(&orders.Order{})
		db.DropTableIfExists(&jobs.WorkOrder{})

		db.DropTableIfExists(&schedule.WorkSchedule{})
		db.DropTableIfExists(&parts.Part{})
		db.DropTableIfExists(&parts.Contact2Part{})

	}

	// Settings
	db.AutoMigrate( &fp.Setting{} )
	fp.DB_IndexSetting(db)

	// Security
	db.AutoMigrate( &fp.Security{} )
	fp.DB_IndexSecurity(db)
	fp.DB_CreateDefaultSecurityLevels(db)

	// SysLog
	db.AutoMigrate( &fp.SysLog{} )
	fp.DB_IndexSysLog(db)

	if true == true {
		db.AutoMigrate(&accounts.Account{})
		accounts.DB_IndexAccount(db)
		db.AutoMigrate(&accounts.Address{})
		accounts.DB_IndexAddress(db)
		db.AutoMigrate(&accounts.Contact{})
		accounts.DB_IndexContact(db)


		// orders
		db.AutoMigrate(&orders.Leger{})
		db.AutoMigrate(&orders.OrderType{})
		db.AutoMigrate(&orders.Order{})
		db.AutoMigrate(&orders.OrderItem{})

		orders.DB_IndexLegers(db)
		orders.DB_IndexOrderType(db)
		orders.DB_IndexOrder(db)
		orders.DB_IndexOrderItems(db)

		// jobs
		jobs.DB_IndexWorkOrder(db)
		db.AutoMigrate(&jobs.WorkOrder{})

		orders.DB_CreateDefaultOrderTypes(db)

		db.AutoMigrate(&schedule.WorkSchedule{})

		db.AutoMigrate(&parts.Part{}, &parts.Contact2Part{})
	}
	return foo, nil
}


