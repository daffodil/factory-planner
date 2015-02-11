package controllers

import (
	"github.com/revel/revel"


)


type Pages struct {
	*revel.Controller
}

func (c Pages) Index() revel.Result {

	return c.Render()
}


func (c Pages) Login() revel.Result {
	return c.Render()
}

func (c Pages) DesktopInstall() revel.Result {
	return c.Render()
}


//===============================


func init(){
	revel.InterceptFunc(setContext, revel.BEFORE, &Pages{})
}
