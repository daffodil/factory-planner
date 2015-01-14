package controllers

import (
	"fmt"
	"github.com/revel/revel"

	"github.com/daffodil/factory-planner/app"
	"github.com/daffodil/factory-planner/app/fp/dev"

)

type Dev struct {
	*revel.Controller
}
// Show the extjs panel
// TODO make dev  panel in it/ dept instead of staff
func (c Dev) DevPage() revel.Result {
	c.RenderArgs["CurrPath"] = "/staff/dev"
	c.RenderArgs["MainNav"] = StaffNav()

	return c.RenderTemplate("staff/dev_panel.html")
}




// Json struct for revel routes
type RoutesInfoPayload struct {
	Success bool `json:"success"`
	Routes []RouteInfo `json:"routes"`
}
// revel route info
type RouteInfo struct {
	Url string `json:"url"`
	Controller string `json:"controller"`
	Action string `json:"action"`
}

// /ajax/dev/routes
// routes view - get the loaded routes and encodes for json
func (c Dev) RoutesJson(table string) revel.Result {

	data := new(RoutesInfoPayload)
	data.Success = true

	data.Routes = make([]RouteInfo,0)
	for _, r := range revel.MainRouter.Routes {
		data.Routes = append(data.Routes, RouteInfo{Url:r.Path, Controller: r.ControllerName, Action: r.Action} )
	}

	return c.RenderJson(data)
}


// DEADReturns all db info
/*
func (c Dev) DB_InfoJson() revel.Result {

	data, err := dev.DB_GetTablesAndViewsPayload(app.Db.DB())
	if err != nil {
		revel.ERROR.Println(err)
	}
	return c.RenderJson(data)
}
*/


// Returns tables
func (c Dev) DB_TablesJson() revel.Result {
	data, err := dev.DB_GetTablesPayload(app.Db.DB())
	if err != nil {
		revel.ERROR.Println(err)
	}
	return c.RenderJson(data)
}

//=============================================================================
// Create tables
func (c Dev) DB_TablesCreateJson() revel.Result {

	drop := c.Params.Get("drop") == "1"

	data, err := dev.DB_CreateTables(app.Db, drop)
	if err != nil {
		revel.ERROR.Println(err)
	}
	return c.RenderJson(data)
}
//=============================================================================
// Views index
func (c Dev) DB_ViewsJson() revel.Result {

	var payload map[string]interface{} = make(map[string]interface{})
	payload["success"] = true
	var err error
	payload["views"], err = dev.DB_GetViews(app.Db.DB())
	if err != nil {
		revel.ERROR.Println(err)
	}

	return c.RenderJson(payload)
}


//=============================================================================
// Create the database Views
func (c Dev) DB_ViewsCreateJson() revel.Result {

	var payload map[string]interface{} = make(map[string]interface{})
	payload["success"] = true
	var err error
	payload["views"], err = dev.DB_CreateViews(app.Db)
	if err != nil {
		revel.ERROR.Println(err)
	}

	return c.RenderJson(payload)
}
//=============================================================================
// Show a database view
func (c Dev) DB_ViewJson(view string) revel.Result {

	var payload map[string]interface{} = make(map[string]interface{})
	payload["success"] = true
	var err error
	fmt.Println("GOT+", view)
	payload["view"], err = dev.DB_GetView(app.Db, view)
	if err != nil {
		revel.ERROR.Println(err)
	}

	return c.RenderJson(payload)
}

// ====================================================================
// Returns data on a table
func (c Dev) DB_TableJson(table string) revel.Result {

	data, err := dev.DB_GetTablePayload(app.Db.DB(), table)
	if err != nil {
		revel.ERROR.Println(err)
	}
	return c.RenderJson(data)
}


