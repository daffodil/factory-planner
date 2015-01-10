package controllers

import (
	"github.com/revel/revel"

	//"github.com/daffodil/factory-planner/app"
	//"github.com/daffodil/factory-planner/mods/accounts"
)

type Account struct {
	*revel.Controller
}


type AccountsPayload struct {
	Success bool `json:"success"`
	//Accounts []accounts.Account `json:"accounts"`
}

func (c Account) JsonAccountsIndex() revel.Result {

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

func (c Account) JsonAccount() revel.Result {

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
