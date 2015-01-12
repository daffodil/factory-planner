package orders

import (
	"time"
	"github.com/jinzhu/gorm"
)



type WorkOrder struct {

	WorkOrderId int `json:"work_order_id" gorm:"column:work_order_id; primary_key:yes"`
	OrderId int `json:"order_id" sql:"type:int(11);not null;" `


	WorkOrderNo string `json:"work_order_no" sql:"type:varchar(100);not null;default:''" `
	WorkOrderNotes string `json:"work_order_notes" sql:"type:varchar(255);not null;default:''" `

	WorkOrderRequired time.Time `json:"work_order_required" deadsql:"type:varchar(10);not null;default:''" `
}

func (me WorkOrder) TableName() string {
	return "works_orders"
}
func DB_IndexWorkOrder(db gorm.DB) {

	cols := []string{
		"order_id", "work_order_no", "work_order_required" }

	for _, c := range cols {
		db.Model(&WorkOrder{}).AddIndex("idx_" + c, c)
	}
}
