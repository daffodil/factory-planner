package controllers

import (
	"github.com/revel/revel"

	"github.com/daffodil/factory-planner/app"
	"github.com/daffodil/factory-planner/app/fp"
	"github.com/daffodil/factory-planner/app/fp/files"
	//"github.com/daffodil/factory-planner/app/fp/orders"
)

type Files struct {
	*revel.Controller
}




// handles /ajax/files
func (c Files) FilesIndexJson() revel.Result {

	var e error
	payload := MakePayload()


	search_vars := fp.GetSearchVars(c.Params.Query)

	payload["files"], e = files.GetFilesIndex(app.Db, search_vars)
	if e != nil {
		payload["error"] = e.Error()
	}

	return c.RenderJson(payload)
}


// handles file by id /ajax/file/is/;file_id
func (c Files) FileJson(file_id int) revel.Result {

	var e error
	payload := MakePayload()

	payload["file"], e = files.GetFile(app.Db, file_id)
	if e != nil {
		payload["error"] = e
		return c.RenderJson(payload)
	}

	return c.RenderJson(payload)
}




// Render extjs panel
/*
func (c Files) StaffAccountsPage() revel.Result {

	c.RenderArgs["CurrPath"] = "/staff/accounts"
	c.RenderArgs["MainNav"] = StaffNav()
	return c.RenderTemplate("staff/accounts.html")
}
*/
