package controllers

import (
	//"strings"
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

	if e != nil {
		//payload.Error = e.Error()
	}

	pay["sys_info"] = fp.GetSysInfo(app.Db)

	return c.RenderJson(pay)
}

