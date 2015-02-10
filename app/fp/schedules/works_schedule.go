
package schedules

import (
	"fmt"
	"time"
	"sort"
	"github.com/jinzhu/gorm"

	"github.com/daffodil/factory-planner/app/fp/projects"
)


type WorkSchedule struct {

	WorkSchedId int ` json:"work_sched_id"  gorm:"column:work_sched_id; primary_key:yes" `
	JobItemId int ` json:"job_item_id" `
	//WsTypeId int ` json:"ws_type_id" `
	WorkSchedRequired time.Time ` json:"work_sched_required"  `

	XWorkSchedYear int  ` json:"x_work_sched_year"  `
	XWorkSchedWeek int ` json:"x_work_sched_week"  `



	WorkSchedQtyRequired int ` json:"work_sched_qty_required"  `
	WorkSchedQtyCompleted int ` json:"work_sched_qty_completed"  `

}

type WorkScheduleView struct {
	WorkSchedule
	WorkSchedYear int  ` json:"work_sched_year"  `
	WorkSchedWeek int ` json:"work_sched_week"  `


	JobId int ` json:"job_id" `
	OrderId int ` json:"order_id" `
	PurchaseOrder string ` json:"purchase_order" `

	OrderTypeId int ` json:"order_type_id" `
	OrderType string ` json:"order_type" `

	ProjectId int ` json:"project_id" `
	ProjectRef string ` json:"project_ref" `

	AccountId int ` json:"account_id" `
	Company string ` json:"company" `
	Ticker string ` json:"ticker" `
	//ModelsCount int ` json:"models_count" `
}

var WORK_SCHEDULE_VIEW = "v_work_schedules"
var WORK_SCHEDULE_VIEW_COLS string = `
work_sched_id, job_item_id, job_id, order_id,
order_type_id, order_type, purchase_order,
project_id, project_ref,
work_sched_year, work_sched_week, work_sched_required,
work_sched_qty_required,
account_id, company, ticker
`


func (me WorkSchedule) TableName() string {
	return "work_schedules"
}


func GetWorkSchedules(db gorm.DB) ([]*WorkScheduleView, error) {

	rows := make([]*WorkScheduleView,0)
	db.Table(WORK_SCHEDULE_VIEW).Select(WORK_SCHEDULE_VIEW_COLS).Scan(&rows)

	return rows, nil
}

func GetWorkSchedulesByDateRange(db gorm.DB, date_start, date_end string) ([]*WorkScheduleView, error) {

	rows := make([]*WorkScheduleView,0)
	view_cols := []string{
		"work_sched_id", "work_order_id", "work_sched_year", "work_sched_week", "work_sched_required"}
	//"on_hold"	, "is_client", "is_supplier", "online"}

	where := "work_sched_required >= ? and work_sched_required <= ?"
	//fmt.Println("where=", where)
	db.Table("v_work_schedules").Select(view_cols).Where(where, date_start, date_end).Scan(&rows)

	return rows, nil

}

type T_WorkSchedTree struct {
	ModelProject []T_ModelProject
}
type T_ModelProject struct {
	Models []T_Model
	Projects []T_Project
}

type T_Model struct {
	ModelId int
	Model string
	BrandId int
	Brand string

}

type T_Project struct {
	ProjectId int
	ProjectRef string
	Jobs []T_Order
}

type T_Order struct {
	ProjectTypeId int
	JobId int
	OrderId int
	PurchaseOrder string
	WorkSchedule []T_WorkSched
}

type T_WorkSched struct {
	Year int
	Week int
	Date time.Time
}

func GetWorkSchedulesTree(db gorm.DB) (*T_WorkSchedTree, error) {

	tree := new(T_WorkSchedTree)
	tree.ModelProject = make([]T_ModelProject, 0)

	model_rows, errm := projects.GetModels(db)
	if errm != nil {

	}
	var models_map map[int]*projects.ModelView = make(map[int]*projects.ModelView)
	for _, m := range model_rows {
		_, ok := models_map[m.ModelId]
		if ok == false {
			models_map[m.ModelId] = m
		}
	}
	fmt.Println(models_map)



	project_2_models_lookup, err_pro := projects.GetProject2ModelsLookup(db)
	if err_pro != nil {

	}

	model_heads := make(map[string][]int)
	for pid, mm := range project_2_models_lookup {
		fmt.Println("mm=", pid, mm)
		mods := make([]int, 0)
		for _, mo := range mm {
			mods = append(mods, mo.ModelId)
			fmt.Println("      mo=", mo, model_heads)
		}
		fmt.Println(" mode=", mods)
		sort.Ints(mods)
		//sort.Reverse(mods)
		fmt.Println(" mode=", mods)
		//_, found := model_heads
	}

	//var projects_map map[int][]int = make(map[int][]int)
	//for _, r := range project_2_models_rows {
	//	_, ok := projects_2_models_map[r.ProjectId]
	//	if ok == false {
	//		projects_2_models_map[r.ProjectId] = make([]int, 0)
	//	}
	//	projects_2_models_map[r.ProjectId] = append(projects_2_models_map[r.ProjectId], r.ModelId)
	//}

	//fmt.Println(projects_2_models_map)

	//rows := make([]*WorkScheduleView,0)
	//db.Table(WORK_SCHEDULE_VIEW).Select(WORK_SCHEDULE_VIEW_COLS).Scan(&rows)

	// Get Projects
	project_rows, errp := projects.GetProjects(db)
	if errp != nil {

	}

	// Make Projects Map
	var projects_map map[int]T_Project = make(map[int]T_Project)
	for _, p := range project_rows {
		projects_map[p.ProjectId] = T_Project{ProjectId: p.ProjectId, ProjectRef: p.ProjectRef}
	}
	//fmt.Println(projects_map)


	for _, r := range project_2_models_lookup {
		mm := T_ModelProject{}
		mm.Models = make([]T_Model, 0)
		mm.Projects = make([]T_Project, 0)
		if 1 == 0 {
			fmt.Println(r)
		}
		tree.ModelProject = append(tree.ModelProject, mm)
	}

	return tree, nil
}
