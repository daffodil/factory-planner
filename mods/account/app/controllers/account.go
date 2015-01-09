package controllers

import "github.com/revel/revel"

type Account struct {
	*revel.Controller
}

func (c Account) JMobileIndex() revel.Result {
	return c.Render()
}

func (c Account) JsonIndex() revel.Result {
	return c.Render()
}
