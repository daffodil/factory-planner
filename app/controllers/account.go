package controllers

import (
	"github.com/revel/revel"

	"github.com/daffodil/factory-planner/mods/accounts"
)

type Account struct {
	*revel.Controller
}

func (c Account) JMobileIndex() revel.Result {
	return c.Render()
}

func (c Account) JsonIndex() revel.Result {

	payload := make(map[string]interface{})
	payload["accounts"] = accounts.AccountsIndex()

	return c.RenderJson()
}

