

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

	views["v_addresses"] = `
		create or replace view v_addresses as
		select
		addresses.address_id, addresses.location, addresses.address,  addresses.postcode,
		addresses.tel, addresses.fax, addresses.is_hq, addresses.is_billing, addresses.addr_active,
		addresses.account_id, accounts.company
		from addresses
		inner join accounts on addresses.account_id = accounts.account_id
		order by location asc
	`


	views["v_brands"] = `
		create or replace view v_brands as
		select brands.brand_id, brand,
		brands.account_id, accounts.company, accounts.ticker, accounts.acc_ref
		from brands
		inner join accounts on brands.account_id = accounts.account_id
		order by brand asc
	`

	views["v_contacts"] = `
		create or replace view v_contacts as
		select
		contacts.contact_id, contacts.contact, contacts.title,  contacts.con_active,
		contacts.email, contacts.mobile, contacts.direct_line,
		contacts.account_id, accounts.company, accounts.ticker, accounts.acc_ref, accounts.root, accounts.acc_active, accounts.online,
		contacts.can_login,  contacts.syslogin, contacts.pass_change, contacts.security_id, security.security,
		contacts.www_page, contacts.address_id
		from contacts
		inner join security on contacts.security_id = security.security_id
		inner join accounts on contacts.account_id = accounts.account_id

		order by contact asc
	`


	views["v_files"] = `
		create or replace view v_files as
		select files.file_id,
		file_name, file_description, file_date, file_checksum, file_uid,
		mime_type, revision,
		contacts.account_id, accounts.company, accounts.ticker,
		files.contact_id, contacts.contact
		from files
		inner join contacts on contacts.contact_id = files.contact_id
		inner join accounts on accounts.account_id = contacts.account_id

		order by file_date desc
	`

	views["v_jobs"] = `
		create or replace view v_jobs as
		select
		jobs.job_id, jobs.order_id,
		orders.purchase_order, orders.client_extra_ref, order_ordered, order_required,
		orders.account_id, accounts.company, accounts.ticker,
		(select count(*) from job_items
			where jobs.job_id = job_items.job_id
		) as job_items_count
		from jobs
		inner join orders on jobs.order_id = orders.order_id
		inner join accounts on orders.account_id = accounts.account_id
		order by job_id desc
	`

	views["v_job_items"] = `
		create or replace view v_job_items as
		select
		job_items.job_item_id, job_items.job_id, job_items.project_id,
		job_items.item_description, job_items.item_qty
		from job_items
		inner join jobs on jobs.job_id = job_items.job_id
		order by job_item_id asc
	`

	views["v_models"] = `
		create or replace view v_models as
		select models.model_id, models.model,
		models.brand_id, brands.brand,
		brands.account_id, accounts.company, accounts.ticker,accounts.acc_ref,
		(select count(project_id) from project_model_links
			where models.model_id = project_model_links.project_id
		) as projects_count
		from models
		inner join brands on brands.brand_id = models.brand_id
		inner join accounts on brands.account_id = accounts.account_id
		order by model, brand
	`

	views["v_orders"] = `
		create or replace view v_orders as
		select orders.order_id, orders.purchase_order,
		orders.account_id, accounts.company, accounts.ticker,
		order_ordered, order_required, client_extra_ref
		from orders
		inner join accounts on accounts.account_id = orders.account_id

		order by order_required asc
	`

	views["v_projects"] = `
		create or replace view v_projects as
		select
		projects.project_id, projects.project_ref, projects.project_description,
		projects.account_id, accounts.ticker, accounts.company,
		(select count(model_id) from project_model_links
			where project_model_links.project_id = projects.project_id
		) as models_count
		from projects
		inner join accounts on accounts.account_id = projects.account_id
		order by project_ref asc
	`
	views["v_project_models"] = `
		create or replace view v_project_models as
		select
		project_model_links.project_id, projects.project_ref, projects.project_description,
		accounts.account_id, accounts.ticker, accounts.company,
		models.model_id, models.model, brands.brand_id, brands.brand
		from project_model_links
		inner join projects on project_model_links.project_id = projects.project_id
		inner join models on project_model_links.model_id = models.model_id
		inner join brands on models.brand_id = brands.brand_id
		inner join accounts on brands.account_id = accounts.account_id
	`

	views["v_work_schedules"] = `
		create or replace view v_work_schedules as
		select
		work_schedules.work_sched_id, work_schedules.job_item_id,
		job_items.job_type_id, job_types.job_type, job_types.job_type_color,
		work_schedules.work_sched_required,
		YEAR( work_schedules.work_sched_required ) as work_sched_year,
		WEEKOFYEAR( work_schedules.work_sched_required ) as work_sched_week,
		work_schedules.x_work_sched_year, work_schedules.x_work_sched_week,
		work_sched_qty_required,
		job_items.job_id, jobs.order_id,  orders.purchase_order,
		accounts.account_id, accounts.ticker, accounts.company,
		job_items.project_id, projects.project_ref
		from work_schedules
		inner join job_items on job_items.job_item_id = work_schedules.job_item_id
		inner join projects on job_items.project_id = projects.project_id
		inner join job_types on job_items.job_type_id = job_types.job_type_id
		inner join jobs on jobs.job_id = job_items.job_id
		inner join orders on jobs.order_id = orders.order_id
		inner join accounts on accounts.account_id = orders.account_id
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
