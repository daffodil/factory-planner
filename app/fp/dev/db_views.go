

package dev

import (

	"fmt"

	"github.com/revel/revel"

	"github.com/jinzhu/gorm"
	"database/sql"

)

// The sql  of views to create
var views map[string]string

func init() {
	views = make(map[string]string)
   
	views["v_accounts"] = `
		create or replace view v_accounts as
		select accounts.account_id, accounts.root, accounts.acc_active,
		company, ticker, online, is_supplier, is_client,
		acc_ref,  on_hold,
		(select count(*) from orders
					where orders.account_id = accounts.account_id
					) as orders_due
		from accounts

		order by company asc
	`

	views["v_brands"] = `
		create or replace view v_brands as
		select brands.brand_id, brand,
		brands.account_id, accounts.company, accounts.ticker, accounts.acc_ref
		from brands
		inner join accounts on brands.account_id = accounts.account_id
		order by brand asc
	`

	views["v_files"] = `
		create or replace view v_files as
		select files.file_id,
		file_name, file_description, file_date, file_checksum, file_uid,
		mime_type, revision,
		files.account_id, accounts.company, accounts.ticker,
		files.contact_id, contacts.contact
		from files
		inner join contacts on contacts.contact_id = files.contact_id
		inner join accounts on accounts.account_id = contacts.account_id

		order by file_date desc
	`

	views["v_models"] = `
		create or replace view v_models as
		select models.model_id, models.model,
		models.brand_id, brands.brand,
		brands.account_id, accounts.company, accounts.ticker,accounts.acc_ref
		from models
		inner join brands on brands.brand_id = models.brand_id
		inner join accounts on brands.account_id = accounts.account_id
		order by model, brand
	`

	views["v_orders"] = `
		create or replace view v_orders as
		select orders.order_id, orders.account_id, accounts.company, accounts.ticker,
		orders.order_type_id, order_types.order_type, order_types.order_color,
		order_ordered, order_required, client_order_no
		from orders
		left join order_types on order_types.order_type_id = orders.order_type_id
		inner join accounts on accounts.account_id = orders.account_id

		order by order_required asc
	`


	views["v_work_schedules"] = `
		create or replace view v_work_schedules as
		select
		work_schedules.work_sched_id, work_schedules.work_order_id,
		work_schedules.work_sched_required,
		YEAR( work_schedules.work_sched_required ) as work_sched_year,
		WEEKOFYEAR( work_schedules.work_sched_required ) as work_sched_week,
		work_schedules.x_work_sched_year, work_schedules.x_work_sched_week,
		work_orders.work_order_no
		from work_schedules
		inner join work_orders on work_orders.work_order_id = work_schedules.work_order_id
		order by work_schedules.work_sched_required asc
	`
}

type DB_View struct {
	Name string ` json:"name" `
	SQL string ` json:"sql;omitempty" `
	Error string ` json:"error" `
}

func DB_CreateViews(db gorm.DB) (map[string]DB_View, error) {


	reply := make( map[string]DB_View )

	for ki := range views {
		view := DB_View{Name: ki}
		fmt.Println("===", ki) //, views[ki])
		foo := db.Exec(views[ki])
		if foo.Error != nil {
			view.Error = foo.Error.Error()
		}
		//fmt.Println("foo=", foo.Error)
		reply[ki] = view
	}


	return reply, nil

}


func DB_GetViews(DB *sql.DB) ([]DB_Table,  error) {

	lst :=  make( []DB_Table, 0)

	sql := "select table_name  "
	sql += " from INFORMATION_SCHEMA.views WHERE table_schema = ?"
	revel.INFO.Println( GetDatabaseName(), sql)
	rows, err := DB.Query(sql, GetDatabaseName() )
	if err != nil {
		revel.ERROR.Println(err)
		return nil,  err
	}
	defer rows.Close()

	for rows.Next(){
		t := DB_Table{}
		err := rows.Scan( &t.Name )
		if err != nil {
			revel.ERROR.Println(err)
		} else {
			t.IsView = true
			lst = append(lst, t)
		}
	}
	return lst, nil
}


func DB_GetView(db gorm.DB, view string) (DB_View, error) {

	viewObj := DB_View{}

	view_def, found := views[view]
	fmt.Println(found, view)

	viewObj.Name = view
	viewObj.SQL = view_def

	return viewObj, nil

}
