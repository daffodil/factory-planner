package controllers

import (
	"time"
	"github.com/revel/revel"

	"github.com/daffodil/factory-planner/app"
	"github.com/daffodil/factory-planner/app/fp/accounts"
	"github.com/daffodil/factory-planner/app/fp/jobs"
	"github.com/daffodil/factory-planner/app/fp/projects"
)

func init(){
	revel.InterceptFunc(checkUser, revel.BEFORE, &JMobile{})
}

type NavMobile struct {
	Path string
	Label string
	Title string
	Icon string
}


func StaffMobileNav() []NavMobile {
	nav := make([]NavMobile, 0)
	nav = append(nav, NavMobile{"index", "Index", "Staff Index", "home"})
	nav = append(nav, NavMobile{"accounts", "Accounts", "Accounts Index", "info"})
	nav = append(nav, NavMobile{"projects", "Projects", "Projects Index", "grid"})
	nav = append(nav, NavMobile{"orders", "Orders", "Orders Index", "grid"})
	return nav
}

type JMobile struct {
	*revel.Controller
}

// Populates Nav, sets current time etc
// TODO - handle pages here
func (c JMobile) Setup(tab string) {
	c.RenderArgs["tab"] = tab
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

func (c JMobile) Accounts() revel.Result {

	var e error
	c.Setup("accounts")

	c.RenderArgs["accounts"], e = accounts.GetAccounts(app.Db)
	if e != nil {

	}

	return c.Render()
}


func (c JMobile) Account(account_id int) revel.Result {

	c.Setup("accounts")

	var e error
	c.RenderArgs["account"], e = accounts.GetAccount(app.Db, account_id)
	if e != nil {

	}
	// Addresses
	c.RenderArgs["addresses"], e = accounts.GetAccountAddresses(app.Db, account_id)
	if e != nil {

	}


	// Contacts
	c.RenderArgs["contacts"], e = accounts.GetAccountContacts(app.Db, account_id)
	if e != nil {

	}

	// Jobs
	c.RenderArgs["jobs"], e = jobs.GetAccountJobs(app.Db, account_id)
	if e != nil {

	}

	// Models
	c.RenderArgs["models"], e = projects.GetAccountModelsNested(app.Db, account_id)
	if e != nil {

	}

	// Projects
	c.RenderArgs["projects"], e = projects.GetAccountProjects(app.Db, account_id)
	if e != nil {

	}



	return c.Render()
}
