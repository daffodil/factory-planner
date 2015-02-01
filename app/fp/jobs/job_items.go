package jobs

import (
	//"time"
	"github.com/jinzhu/gorm"
)



type JobItem struct {

	JobItemId int `json:"job_item_id" gorm:"column:job_item_id; primary_key:yes"`
	//JobNo int `json:"job_no" sql:"type:varchar(200);not null;" `
	JobId int `json:"job_id"  `
	ProjectId int `json:"project_id"  `
	ItemDescription string `json:"item_description"  `
	JobItemImport string `json:"job_item_import"  `
	Qty int ` json:"item_qty" `
	//OrderId int `json:"order_id"  `
	//JobImport string `json:"job_import" sql:"type:varchar(255);not null;default:''" `

	//WorkOrderRequired time.Time `json:"work_order_required" deadsql:"type:varchar(10);not null;default:''" `
}

func (me JobItem) TableName() string {
	return "job_items"
}
func DB_IndexJobItems(db gorm.DB) {

	cols := []string{"job_id", "project_id"  }

	for _, c := range cols {
		db.Model(&JobItem{}).AddIndex("idx_" + c, c)
	}
}


func GetJobItems(db gorm.DB, job_id int) ([]JobItem, error) {

	recs := make([]JobItem, 0)
	//db.Find(&worders, WorkOrder{AccountId: account_id})

	return recs, nil
}
