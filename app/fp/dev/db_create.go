

package dev

import (
	//"os"
	//"path/filepath"
	//"io/ioutil"
	"fmt"
	//"github.com/revel/revel"
	//"database/sql"
	"github.com/jinzhu/gorm"

	"github.com/daffodil/factory-planner/app/fp"
	"github.com/daffodil/factory-planner/app/fp/accounts"
	"github.com/daffodil/factory-planner/app/fp/files"
	"github.com/daffodil/factory-planner/app/fp/jobs"
	"github.com/daffodil/factory-planner/app/fp/orders"
	"github.com/daffodil/factory-planner/app/fp/parts"
	"github.com/daffodil/factory-planner/app/fp/projects"
	"github.com/daffodil/factory-planner/app/fp/schedules"


)


func DB_CreateTables(db gorm.DB, drop_first bool) (interface{}, error) {

	foo := make( map[string]interface{} )

	db.LogMode(false)

	if drop_first {
		db.DropTableIfExists(&fp.Security{})
		db.DropTableIfExists(&fp.Setting{})

		db.DropTableIfExists(&accounts.Account{})
		db.DropTableIfExists(&accounts.Address{})
		db.DropTableIfExists(&accounts.Contact{})

		db.DropTableIfExists(&projects.ProjectType{})
		db.DropTableIfExists(&orders.Order{})
		db.DropTableIfExists(&jobs.Job{})
		//db.DropTableIfExists(&jobs.JobType{})

		db.DropTableIfExists(&schedules.WorkSchedule{})
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

		// Files
		db.AutoMigrate(&files.File{})
		files.DB_IndexFiles(db)

		// Projects
		db.AutoMigrate(&projects.Brand{})
		projects.DB_IndexBrands(db)
		db.AutoMigrate(&projects.Model{})
		projects.DB_IndexModels(db)

		db.AutoMigrate(&projects.ProjectType{})
		projects.DB_CreateDefaultProjectTypes(db)
		projects.DB_CreateDefaultProjectTypes(db)

		db.AutoMigrate(&projects.Project{})
		projects.DB_IndexProjects(db)

		db.AutoMigrate(&projects.ProjectModelLink{})
		projects.DB_IndexProjectModelLinks(db)

		db.AutoMigrate(&projects.ProjectContactLink{})
		projects.DB_IndexProjectContactLinks(db)

		// Orders
		db.AutoMigrate(&orders.Leger{})
		//db.AutoMigrate(&orders.OrderType{})
		db.AutoMigrate(&orders.Order{})
		db.AutoMigrate(&orders.OrderItem{})

		orders.DB_IndexLegers(db)
		//orders.DB_IndexOrderType(db)
		orders.DB_IndexOrder(db)
		orders.DB_IndexOrderItems(db)




		// jobs
		//db.AutoMigrate(&jobs.JobType{})
		//jobs.DB_IndexJobType(db)


		db.AutoMigrate(&jobs.Job{})
		jobs.DB_IndexJob(db)
		db.AutoMigrate(&jobs.JobItem{})
		jobs.DB_IndexJobItems(db)




		db.AutoMigrate(&schedules.WorkSchedule{})

		db.AutoMigrate(&parts.Part{}, &parts.Contact2Part{})
	}
	db.LogMode(true)
	fmt.Println("LLLLLLLLLLLLLLLLLLLLEEEEEEEEEEEEEG")
	orders.DB_CreateDefaultLegers(db)
	db.LogMode(false)
	return foo, nil
}


