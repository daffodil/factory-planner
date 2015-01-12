package controllers

import (
	//"time"
	"github.com/revel/revel"

	"github.com/daffodil/factory-planner/app"
	"github.com/daffodil/factory-planner/mods/orders"
)


type Orders struct {
	*revel.Controller
}

type OrdersPayload struct {
	Success bool `json:"success"`
	Error string  `json:"error"`
	Orders []orders.OrderView `json:"orders"`
}

func (c Orders) JsonAccountOrders(account_id int) revel.Result {

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
	WorkOrders []orders.WorkOrder `json:"work_orders"`
}
func (c Orders) JsonAccountWorkOrders(account_id int) revel.Result {

	var e error
	payload := new(WorkOrdersPayload)
	payload.Success = true

	payload.WorkOrders, e = orders.GetAccountWorkOrders(app.Db, account_id)
	if e != nil {
		payload.Error = e.Error()
	}

	return c.RenderJson(payload)
}
