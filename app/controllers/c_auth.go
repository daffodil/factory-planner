package controllers

import (
	//"strings"
	"fmt"
	"strconv"
	"github.com/revel/revel"

	"github.com/daffodil/factory-planner/app"
	"github.com/daffodil/factory-planner/app/fp"
	"github.com/daffodil/factory-planner/app/fp/accounts"
)

type Auth struct {
	*revel.Controller
}


// handles /ajax/staff
func (c Auth) StaffLogin() revel.Result {

	var e error
	pay := MakePayload()

	user, login_error, err := accounts.AuthenticateUser(app.Db, c.Params.Get("syslogin"),  c.Params.Get("secret"))
	if err != nil {

	}
	if login_error != nil {
		pay["login_error"] = login_error.Error()
	}
	pay["user"] = user
	c.Session["user"] = strconv.Itoa(user.ContactId)

	if e != nil {
		//payload.Error = e.Error()
	}

	pay["sys_info"] = fp.GetSysInfo(app.Db)

	return c.RenderJson(pay)
}


// /jmobile/login
func (c Auth) Mobile() revel.Result {

	var e error
	if e != nil {

	}

	return c.RenderTemplate("jmobile/login.html")
}
// /jmobile/login
func (c Auth) MobileLogin() revel.Result {

	//c.RenderArgs["login_error"] = ""

	user, login_error, err := accounts.AuthenticateUser(app.Db, c.Params.Get("syslogin"),  c.Params.Get("secret"))
	if err != nil {
		fmt.Println(err)
	}
	if login_error != nil {
		c.RenderArgs["login_error"] = login_error.Error()
		fmt.Println(login_error.Error())
	} else {
		c.Session["user"] = strconv.Itoa(user.ContactId)
		return c.Redirect("/jmobile")
	}

	return c.RenderTemplate("jmobile/login.html")
}
