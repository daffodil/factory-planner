package controllers

import (
	"github.com/revel/revel"

	"github.com/daffodil/factory-planner/app"
	"github.com/daffodil/factory-planner/app/fp"
	"github.com/daffodil/factory-planner/app/fp/accounts"
	"github.com/daffodil/factory-planner/app/fp/orders"
)

type Accounts struct {
	*revel.Controller
}


type AccountsPayload struct {
	Success bool `json:"success"`
	Error string  `json:"error"`
	Accounts []accounts.AccountView `json:"accounts"`
}



// handles /ajax/accounts
func (c Accounts) AccountsJson() revel.Result {

	var e error
	payload := new(AccountsPayload)
	payload.Success = true
	//payload := make(map[string]interface{})

	//search := GetSearch( c )
	search_vars := fp.GetSearchVars(c.Params.Query)

	payload.Accounts, e = accounts.GetAccountsIndex(app.Db, search_vars)

	if e != nil {
		payload.Error = e.Error()
	}
	return c.RenderJson(payload)
}


// handles account by id /ajax/account/;account_id
func (c Accounts) AccountJson(account_id int) revel.Result {

	var e error
	payload := MakePayload()
	payload["account"], e = accounts.GetAccount(app.Db, account_id)
	if e != nil {
		payload["error"] = e
		return c.RenderJson(payload)
	}

	payload["orders"], e = orders.GetAccountOrders(app.Db, account_id)
	if e != nil {
		payload["error"] = e
		return c.RenderJson(payload)
	}
	return c.RenderJson(payload)
}

func (c Accounts) RootAccountJson() revel.Result {

	var e error
	payload := MakePayload()
	payload["account"], e = accounts.GetRootAccount(app.Db)
	if e != nil {
		payload["error"] = e.Error()
	}


	return c.RenderJson(payload)
}


// Render extjs panel
func (c Accounts) StaffAccountsPage() revel.Result {

	c.RenderArgs["CurrPath"] = "/staff/accounts"
	c.RenderArgs["MainNav"] = StaffNav()
	return c.RenderTemplate("staff/accounts.html")
}
