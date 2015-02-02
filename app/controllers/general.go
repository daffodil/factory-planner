package controllers

import (
	"github.com/revel/revel"

	"github.com/daffodil/factory-planner/app"
	"github.com/daffodil/factory-planner/app/fp"

)

type General struct {
	*revel.Controller
}



// handles file by id /ajax/file/is/;file_id
func (c General) Alive() revel.Result {

	//var e error
	payload := MakePayload()

	payload["sys_info"] = fp.GetSysInfo(app.Db)

	return c.RenderJson(payload)
}


