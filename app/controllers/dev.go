package controllers

import (
	"github.com/revel/revel"

	"github.com/daffodil/factory-planner/app"
	"github.com/daffodil/factory-planner/mods/dev"

)

type Dev struct {
	*revel.Controller
}


func (c Dev) DevPage() revel.Result {
	return c.RenderTemplate("jpanel/dev_panel.html")
}



// ====================================================================
type RoutesInfoPayload struct {
	Success bool `json:"success"`
	Routes []RouteInfo `json:"routes"`
}
type RouteInfo struct {
	Url string `json:"url"`
	Controller string `json:"controller"`
	Action string `json:"action"`
}

// Handle the routes view - only return /json/* and /xml/* adn /kml/*
func (c Dev) JsonRoutes(table string) revel.Result {

	data := new(RoutesInfoPayload)
	data.Success = true

	data.Routes = make([]RouteInfo,0)
	for _, r := range revel.MainRouter.Routes {
		data.Routes = append(data.Routes, RouteInfo{Url:r.Path, Controller: r.ControllerName, Action: r.Action} )
	}

	return c.RenderJson(data)
}


// Returns all db info
func (c Dev) DB_InfoJson() revel.Result {

	data, err := dev.DB_GetTablesAndViewsPayload(app.Db.DB())
	if err != nil {
		revel.ERROR.Println(err)
	}
	return c.RenderJson(data)
}

// Returns views
func (c Dev) DB_ViewsJson() revel.Result {

	data, err := dev.DB_GetViewsPayload(app.Db.DB())
	if err != nil {
		revel.ERROR.Println(err)
	}
	return c.RenderJson(data)
}

// Returns tables
func (c Dev) DB_TablesJson() revel.Result {
	data, err := dev.DB_GetTablesPayload(app.Db.DB())
	if err != nil {
		revel.ERROR.Println(err)
	}
	return c.RenderJson(data)
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


// return data on a view
func (c Dev) DB_ViewJson(table string) revel.Result {

	data, err := dev.DB_GetViewPayload(app.Db.DB(), table)
	if err != nil {
		revel.ERROR.Println(err)
	}
	return c.RenderJson(data)
}
