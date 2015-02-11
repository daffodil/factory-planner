package controllers

import (
	"github.com/revel/revel"

	"github.com/daffodil/factory-planner/app"
	//"github.com/daffodil/factory-planner/app/fp"
	"github.com/daffodil/factory-planner/app/fp/jobs"
	"github.com/daffodil/factory-planner/app/fp/projects"
	"github.com/daffodil/factory-planner/app/fp/schedules"
)

type Schedules struct {
	*revel.Controller
}



// /work_schedules
func (c Schedules) WorkSchedules() revel.Result {

	var e error
	pay := MakePayload()

	pay["work_schedules"], e = schedules.GetWorkSchedules(app.Db)
	if e != nil {
		pay["error"] = e.Error()
		return c.RenderJson(pay)
	}
	pay["project_2_models"], e = projects.GetProject2Models(app.Db)
	if e != nil {
		pay["error"] = e.Error()
		return c.RenderJson(pay)
	}

	return c.RenderJson(pay)
}

// /work_schedules
func (c Schedules) WorkSchedulesTree() revel.Result {

	var e error
	pay := MakePayload()

	pay["job_types"], e = jobs.GetJobTypes(app.Db)
	if e != nil {
		pay["error"] = e.Error()
		return c.RenderJson(pay)
	}

	pay["schedule"], e = schedules.GetWorkSchedulesTree(app.Db)
	if e != nil {
		pay["error"] = e.Error()
		return c.RenderJson(pay)
	}

	return c.RenderJson(pay)
}
