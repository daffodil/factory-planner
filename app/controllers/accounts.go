package controllers

import (
	"github.com/revel/revel"

	"github.com/daffodil/factory-planner/app"
	"github.com/daffodil/factory-planner/mods/accounts"
)

type Accounts struct {
	*revel.Controller
}


type AccountsPayload struct {
	Success bool `json:"success"`
	Error string  `json:"error"`
	Accounts []accounts.Account `json:"accounts"`
}

func (c Accounts) JsonAccounts() revel.Result {

	var e error
	payload := new(AccountsPayload)
	payload.Success = true
	//payload := make(map[string]interface{})

	payload.Accounts, e = accounts.AccountsIndex(app.Db)

	if e != nil {
		payload.Error = e.Error()
	}
	return c.RenderJson(payload)
}

func (c Accounts) JsonAccount() revel.Result {

	//var e error
	payload := new(AccountsPayload)
	payload.Success = true
	//payload := make(map[string]interface{})

	//payload.Accounts, e = accounts.AccountsIndex(app.Db)

	//if e != nil {
	// throw tantrum
	//}
	return c.RenderJson(payload)
}

// Render extjs panel
func (c Accounts) StaffAccountsPage() revel.Result {
	c.RenderArgs["CurrPath"] = "/staff/accounts"
	c.RenderArgs["MainNav"] = StaffNav()
	return c.RenderTemplate("staff/accounts.html")
}
