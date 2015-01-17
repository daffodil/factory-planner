package controllers

import (
	//"time"
	"github.com/revel/revel"

	"github.com/daffodil/factory-planner/app"
	"github.com/daffodil/factory-planner/app/fp/orders"
	"github.com/daffodil/factory-planner/app/fp/jobs"
)


type Orders struct {
	*revel.Controller
}

type OrdersPayload struct {
	Success bool `json:"success"`
	Error string  `json:"error"`
	Orders []orders.OrderView `json:"orders"`
}

func (c Orders) OrdersJson() revel.Result {

	var e error
	payload := MakePayload()

	payload["orders"], e = orders.GetOrders(app.Db)
	if e != nil {
		payload["error"] = e.Error()
	}

	return c.RenderJson(payload)
}


func (c Orders) AccountOrdersJson(account_id int) revel.Result {

	var e error
	payload := new(OrdersPayload)
	payload.Success = true

	payload.Orders, e = orders.GetAccountOrders(app.Db, account_id)
	if e != nil {
		payload.Error = e.Error()
	}

	return c.RenderJson(payload)
}

type WorkOrdersPayload struct {
	Success bool `json:"success"`
	Error string  `json:"error"`
	WorkOrders []jobs.WorkOrder `json:"work_orders"`
}
func (c Orders) AccountWorkOrdersJson(account_id int) revel.Result {

	var e error
	payload := new(WorkOrdersPayload)
	payload.Success = true

	payload.WorkOrders, e = jobs.GetAccountWorkOrders(app.Db, account_id)
	if e != nil {
		payload.Error = e.Error()
	}

	return c.RenderJson(payload)
}

// Render extjs panel
func (c Orders) StaffOrdersPage() revel.Result {

	c.RenderArgs["CurrPath"] = "/staff/orders"
	c.RenderArgs["MainNav"] = StaffNav()
	return c.RenderTemplate("staff/orders.html")
}
