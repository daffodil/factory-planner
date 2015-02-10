package jobs

import (
	//"time"
	"github.com/jinzhu/gorm"
)



type JobItem struct {

	JobItemId int `json:"job_item_id" gorm:"column:job_item_id; primary_key:yes"`
	//JobNo int `json:"job_no" sql:"type:varchar(200);not null;" `
	JobId int `json:"job_id"  `
	OrderTypeId int `json:"order_type_id"  `
	ProjectId int `json:"project_id"  `
	ItemDescription string `json:"item_description"  `
	JobItemImport string `json:"job_item_import"  `
	ItemQty int ` json:"item_qty" `
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

var JOB_ITEM_VIEW string = "v_job_items"
var JOB_ITEM_VIEW_COLS =  `
job_item_id, job_id, project_id, item_description, item_qty
`

func GetJobItems(db gorm.DB, job_id int) ([]JobItem, error) {

	rows := make([]JobItem, 0)
	res := db.Table(JOB_ITEM_VIEW).Select(JOB_ITEM_VIEW_COLS).Where("job_id=?", job_id).Scan(&rows)
	if res.Error != nil {
		return rows, res.Error
	}

	return rows, nil
}
