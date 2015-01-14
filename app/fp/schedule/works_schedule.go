
package schedule

import (
	"time"

	"github.com/jinzhu/gorm"
)


type WorkSchedule struct {

	WorkSchedId int ` json:"work_sched_id"  gorm:"column:work_sched_id; primary_key:yes" `
	WorkOrderId int ` json:"work_order_id" `

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
}

func (me WorkSchedule) TableName() string {
	return "work_schedules"
}


func GetWorksSchedule(db gorm.DB, date_start, date_end string) ([]*WorkScheduleView, error) {

	rows := make([]*WorkScheduleView,0)
	view_cols := []string{
		"work_sched_id", "work_order_id", "work_sched_year", "work_sched_week", "work_sched_required"}
		//"on_hold"	, "is_client", "is_supplier", "online"}

	where := "work_sched_required >= ? and work_sched_required <= ?"
	//fmt.Println("where=", where)
	db.Table("v_work_schedules").Select(view_cols).Where(where, date_start, date_end).Scan(&rows)

	return rows, nil

}
