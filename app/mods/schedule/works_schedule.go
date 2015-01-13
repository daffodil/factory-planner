
package schedule


type WorkSchedule struct {

	WorkSchedId int ` json:"work_sched_id"  gorm:"column:work_sched_id; primary_key:yes" `
	WorkOrderId int ` json:"work_order_id" `
	WorkSchedYear int  ` json:"work_sched_year"  `
	WorkSchedWeek int ` json:"work_sched_week"  `
	WorkSchedQtyRequired int ` json:"work_sched_qty_required"  `
	WorkSchedQtyCompleted int ` json:"work_sched_qty_completed"  `

}

func (me WorkSchedule) TableName() string {
	return "works_schedule"
}
