package controllers

import (
	"time"
	"github.com/revel/revel"
)

type JMobile struct {
	*revel.Controller
}

// Populates Nav, sets current time etc
func (c JMobile) Setup() {

	c.RenderArgs["now"] = time.Now()

}


func (c JMobile) Index() revel.Result {

	c.Setup()

	return c.Render()
}

func (c JMobile) Pages() revel.Result {

	c.Setup()

	return c.Render()
}

