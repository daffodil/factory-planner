package controllers

import (
	"github.com/revel/revel"

	//"github.com/daffodil/factory-planner/app"
	"github.com/daffodil/factory-planner/mods/calendar"
)

type Calendar struct {
	*revel.Controller
}
type CalendarPayload struct {
	Success bool `json:"success"`
	View string  `json:"view"`
	Week *calendar.Week  `json:"week"`
	//Accounts []accounts.Account `json:"accounts"`
	Error string ""
}

func (c Calendar) JsonWeek(view string) revel.Result {

	var e error
	payload := new(CalendarPayload)
	payload.Success = true
	payload.View = view
	//payload := make(map[string]interface{})

	payload.Week, e = calendar.WeekFromView(view)
	if e != nil {
		payload.Error = e.Error()
	}
	//payload.Accounts, e = accounts.AccountsIndex(app.Db)

	if e != nil {
		// throw tantrum
	}
	return c.RenderJson(payload)
}
