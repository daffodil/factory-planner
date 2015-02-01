package controllers

import (
	//"fmt"
	//"math/rand"
	//"time"
	//"net/http"
	"github.com/revel/revel"

	"github.com/daffodil/factory-planner/app"
	"github.com/daffodil/factory-planner/app/fp"
	"github.com/daffodil/factory-planner/app/fp/accounts"
	"github.com/daffodil/factory-planner/app/fp/orders"
	"github.com/daffodil/factory-planner/app/fp/projects"
)

type Accounts struct {
	*revel.Controller
}


type AccountsPayload struct {
	Success bool `json:"success"`
	Error string  `json:"error"`
	Accounts []accounts.AccountView `json:"accounts"`
}



// handles /ajax/accounts
func (c Accounts) AccountsJson() revel.Result {

	var e error
	payload := new(AccountsPayload)
	payload.Success = true
	//payload := make(map[string]interface{})

	//search := GetSearch( c )
	search_vars := fp.GetSearchVars(c.Params.Query)

	payload.Accounts, e = accounts.GetAccountsIndex(app.Db, search_vars)

	if e != nil {
		payload.Error = e.Error()
	}
	return c.RenderJson(payload)
}

// handles /ajax/accounts/all
func (c Accounts) AccountsAllJson() revel.Result {

	var e error
	payload := new(AccountsPayload)
	payload.Success = true
	//payload := make(map[string]interface{})

	//search := GetSearch( c )
	search_vars := fp.GetSearchVars(c.Params.Query)

	payload.Accounts, e = accounts.GetAccountsIndex(app.Db, search_vars)

	if e != nil {
		payload.Error = e.Error()
	}
	return c.RenderJson(payload)
}


// handles account by id /ajax/account/;account_id
func (c Accounts) AccountJson(account_id int) revel.Result {

	var e error
	payload := MakePayload()

	// account
	payload["account"], e = accounts.GetAccount(app.Db, account_id)
	if e != nil {
		payload["error"] = e
		return c.RenderJson(payload)
	}

	// orders
	payload["orders"], e = orders.GetAccountOrders(app.Db, account_id)
	if e != nil {
		payload["error"] = e
		return c.RenderJson(payload)
	}

	// models
	payload["models"], e = projects.GetAccountModels(app.Db, account_id)
	if e != nil {
		payload["error"] = e
		return c.RenderJson(payload)
	}


	return c.RenderJson(payload)
}


// handles account by id /ajax/account/;account_id/contacts
func (c Accounts) AccountContactsJson(account_id int) revel.Result {

	var e error
	payload := MakePayload()

	// contacts
	payload["contacts"], e = accounts.GetAccountContacts(app.Db, account_id)
	if e != nil {
		payload["error"] = e
		return c.RenderJson(payload)
	}

	return c.RenderJson(payload)
}

// Render extjs panel
func (c Accounts) StaffAccountsPage() revel.Result {

	c.RenderArgs["CurrPath"] = "/staff/accounts"
	c.RenderArgs["MainNav"] = StaffNav()
	return c.RenderTemplate("staff/accounts.html")
}




func (c Accounts) RootAccountJson() revel.Result {

	var e error
	payload := MakePayload()
	payload["account"], e = accounts.GetRootAccount(app.Db)
	if e != nil {
		payload["error"] = e.Error()
	}


	return c.RenderJson(payload)
}


func (c Accounts) RootAccountStaffJson() revel.Result {
	var e error
	payload := MakePayload()
	payload["account"], e = accounts.GetRootAccount(app.Db)
	if e != nil {
		payload["error"] = e.Error()
	}

	payload["staff"], e = accounts.GetStaff(app.Db)
	if e == nil {

	} else {
		payload["error"] = e.Error()
	}
	payload["sys_info"] = fp.GetSysInfo(app.Db)

	/*
	// Random erros and latency
	sleepy := rand.Intn(7)
	erri := rand.Intn(5)
	fmt.Println("sleep=", sleepy, "erri=", erri)
	time.Sleep(time.Duration(sleepy) * time.Second)


	if erri == 1 {
		return c.RenderJson("fooovar")

	} else if erri == 2 {
		c.Response.Status = http.StatusNotFound
		return c.RenderText("404")

	} else if erri == 3 {
		c.Response.Status = http.StatusInternalServerError
		return c.RenderText("500")
	} else if erri == 4 {
		return c.RenderText("das dsa {}")
	}
	*/
	return c.RenderJson(payload)


}


// handles /ajax/contacts
func (c Accounts) ContactsJson() revel.Result {

	var e error
	pay := MakePayload()

	//payload := make(map[string]interface{})

	//search := GetSearch( c )
	search_vars := fp.GetSearchVars(c.Params.Query)

	pay["contacts"], e = accounts.SearchContacts(app.Db, search_vars)

	if e != nil {
		pay["error"] = e.Error()
	}
	return c.RenderJson(pay)
}

// handles /ajax/contact/<id>
func (c Accounts) ContactJson(contact_id int) revel.Result {

	var e error
	pay := MakePayload()

	pay["contact"], e = accounts.GetContact(app.Db, contact_id)

	if e != nil {
		pay["error"] = e.Error()
	}
	return c.RenderJson(pay)
}
