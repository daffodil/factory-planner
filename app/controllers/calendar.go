package controllers

import (
	"github.com/revel/revel"

	//"github.com/daffodil/factory-planner/app"
	"github.com/daffodil/factory-planner/mods/calendar"
)

type Calendar struct {
	*revel.Controller
}


type WeekPayload struct {
	Success bool `json:"success"`
	View string  `json:"view"`
	Week *calendar.Week  `json:"week"`
	Error string ""
}

// /ajax/week/view/;view where view is next_week, this_week, etc
func (c Calendar) JsonWeek(view string) revel.Result {

	var e error
	payload := new(WeekPayload)
	payload.Success = true
	payload.View = view

	payload.Week, e = calendar.WeekFromView(view)
	if e != nil {
		payload.Error = e.Error()
	}

	return c.RenderJson(payload)
}





type WeeksPayload struct {
	Success bool `json:"success"`
	View string  `json:"view"`
	Weeks []*calendar.Week  `json:"weeks"`
	Error string ""
}

// /ajax/weeks where view is next_week, this_week, etc
func (c Calendar) JsonWeeks() revel.Result {

	var e error
	payload := new(WeeksPayload)
	payload.Success = true

	payload.Weeks, e = calendar.WeeksView(2015, 2, -1, 5)
	if e != nil {
		payload.Error = e.Error()
	}

	return c.RenderJson(payload)
}

// render planner
func (c Calendar) StaffPlannerPage() revel.Result {
	c.RenderArgs["CurrPath"] = "/staff/planner"
	c.RenderArgs["MainNav"] = StaffNav()
	return c.RenderTemplate("staff/planner.html")
}

