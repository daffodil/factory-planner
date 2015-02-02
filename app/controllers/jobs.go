package controllers

import (
	"github.com/revel/revel"

	"github.com/daffodil/factory-planner/app"
	"github.com/daffodil/factory-planner/app/fp"
	"github.com/daffodil/factory-planner/app/fp/jobs"
	//"github.com/daffodil/factory-planner/app/fp/orders"
)

type Jobs struct {
	*revel.Controller
}




// handles /ajax/jobs
func (c Jobs) JobsIndex() revel.Result {

	var e error
	payload := MakePayload()


	search_vars := fp.GetSearchVars(c.Params.Query)
	payload["_"] = search_vars
	payload["jobs"], e = jobs.GetJobs(app.Db, "latest")
	if e != nil {
		payload["error"] = e.Error()
	}

	return c.RenderJson(payload)
}


// handles file by id /ajax/file/is/;file_id
func (c Jobs) AccountJobs(account_id int) revel.Result {

	var e error
	payload := MakePayload()

	payload["file"], e = jobs.GetAccountJobs(app.Db, account_id)
	if e != nil {
		payload["error"] = e
		return c.RenderJson(payload)
	}

	return c.RenderJson(payload)
}


