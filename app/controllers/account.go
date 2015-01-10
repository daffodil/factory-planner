package controllers

import (
	"github.com/revel/revel"

	//"github.com/daffodil/factory-planner/app"
	//"github.com/daffodil/factory-planner/mods/accounts"
)

type Account struct {
	*revel.Controller
}

func (c Account) DEADJMobileIndex() revel.Result {
	return c.Render()
}

type AccountsPayload struct {
	Success bool `json:"success"`
	//Accounts []accounts.Account `json:"accounts"`
}

func (c Account) JsonIndex() revel.Result {

	var e error
	payload := new(AccountsPayload)
	payload.Success = true
	//payload := make(map[string]interface{})

	//payload.Accounts, e = accounts.AccountsIndex(app.Db)

	if e != nil {
		// throw tantrum
	}
	return c.RenderJson(payload)
}

