package controllers

import (
	"time"
	"github.com/revel/revel"
)

type NavMobile struct {
	Path string
	Label string
	Title string
	Icon string
}


func StaffMobileNav() []NavMobile {
	nav := make([]NavMobile, 0)
	nav = append(nav, NavMobile{"/index", "Index", "Staff Index", "home"})
	nav = append(nav, NavMobile{"/accounts", "Accounts", "Accounts Index", "info"})
	nav = append(nav, NavMobile{"/parts", "Parts", "Parts Index", "grid"})
	nav = append(nav, NavMobile{"/orders", "Orders", "Orders Index", "grid"})
	return nav
}

type JMobile struct {
	*revel.Controller
}

// Populates Nav, sets current time etc
// TODO - handle pages here
func (c JMobile) Setup(path string) {
	c.RenderArgs["path"] = path
	c.RenderArgs["nav"] = StaffMobileNav()

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

