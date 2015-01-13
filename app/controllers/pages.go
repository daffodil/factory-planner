package controllers

import "github.com/revel/revel"




type Pages struct {
	*revel.Controller
}

func (c Pages) Index() revel.Result {
	return c.Render()
}
