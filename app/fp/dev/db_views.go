

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
		select accounts.account_id, accounts.acc_active,
		company, ticker, online, is_supplier, is_client,
		acc_ref,  on_hold,
		root
		from accounts
		order by company asc
	`

	views["v_orders"] = `
		create or replace view v_orders as
		select orders.order_id, orders.account_id, accounts.company, accounts.ticker,
		orders.order_type_id, order_types.order_type, order_types.order_color,
		order_ordered, order_required, client_order_no
		from orders
		inner join order_types on order_types.order_type_id = orders.order_type_id
		sinner join accounts on accounts.account_id = orders.account_id

		order by order_required asc
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