package jobs

import (
	//"time"
	"github.com/jinzhu/gorm"
)



type Job struct {

	JobId int `json:"job_id" gorm:"column:job_id; primary_key:yes"`
	//JobNo int `json:"job_no" sql:"type:varchar(200);not null;" `


	OrderId int `json:"order_id"  `
	JobImport string `json:"job_import" sql:"type:varchar(255);not null;default:''" `

	//WorkOrderRequired time.Time `json:"work_order_required" deadsql:"type:varchar(10);not null;default:''" `
}

func (me Job) TableName() string {
	return "jobs"
}
func DB_IndexJob(db gorm.DB) {

	cols := []string{"order_id" }

	for _, c := range cols {
		db.Model(&Job{}).AddIndex("idx_" + c, c)
	}
}


func GetAccountJobs(db gorm.DB, account_id int) ([]Job, error) {

	recs := make([]Job, 0)
	//db.Find(&worders, WorkOrder{AccountId: account_id})

	return recs, nil
}
