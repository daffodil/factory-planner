package controllers

import (
	"time"
	"github.com/revel/revel"
)

type Nav struct {
	Path string
	Label string
	Title string
	Icon string
}


func StaffNav() []Nav {
	nav := make([]Nav, 0)
	nav = append(nav, Nav{"/index", "Index", "Staff Index", "home"})
	nav = append(nav, Nav{"/accounts", "Accounts", "Accounts Index", "info"})
	nav = append(nav, Nav{"/parts", "Parts", "Parts Index", "grid"})
	nav = append(nav, Nav{"/orders", "Orders", "Orders Index", "grid"})
	return nav
}

type JMobile struct {
	*revel.Controller
}

// Populates Nav, sets current time etc
func (c JMobile) Setup(path string) {
	c.RenderArgs["path"] = path
	c.RenderArgs["nav"] = StaffNav()

	c.RenderArgs["now"] = time.Now()

}


func (c JMobile) Index() revel.Result {

	c.Setup("index")

	return c.Render()
}

func (c JMobile) Pages() revel.Result {

	c.Setup("foo")

	return c.Render()
}

