

package dev

import (
	//"os"
	//"path/filepath"
	//"io/ioutil"
	"fmt"
	//"github.com/revel/revel"
	//"database/sql"
	"github.com/jinzhu/gorm"

	//"github.com/daffodil/factory-planner/mods/fpsys"
	//"github.com/daffodil/factory-planner/mods/accounts"
	//"github.com/daffodil/factory-planner/mods/orders"
	//"github.com/daffodil/factory-planner/mods/parts"
	//"github.com/daffodil/factory-planner/mods/schedule"

)


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
		inner join accounts on accounts.account_id = orders.account_id

		order by order_required asc
	`
}

func DB_CreateViews(db gorm.DB) (interface{}, error) {


	foo := make( map[string]interface{} )

	for ki := range views {
		fmt.Println(ki, views[ki])
		db.Exec(views[ki])
	}



	return foo, nil

}
