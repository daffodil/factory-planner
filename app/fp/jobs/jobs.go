package jobs

import (

	"time"
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

// Database view extends the `Job` with stuff
type JobView struct {
	Job

	AccountId int ` json:"account_id" `
	Company string ` json:"company" `
	Ticker string ` json:"ticker" `

	PurchaseOrder string  ` json:"purchase_order" `
	OrderOrdered time.Time ` json:"order_ordered" `
	OrderRequired *time.Time ` json:"order_required" `

	JobItemsCount int ` json:"job_items_count"  gorm:"column:job_items_count;" `
}


var JOB_VIEW = "v_jobs"
var JOB_VIEW_COLS =  `
job_id, order_id, purchase_order, account_id, company, ticker,
order_ordered, order_required, job_items_count
`


func GetJobs(db gorm.DB, view string) ([]JobView, error) {

	rows := make([]JobView, 0)
	res := db.Table(JOB_VIEW).Select(JOB_VIEW_COLS).Scan(&rows)
	if res.Error != nil {
		return rows, res.Error
	}
	return rows, nil
}


func GetAccountJobs(db gorm.DB, account_id int) ([]Job, error) {

	recs := make([]Job, 0)
	//db.Find(&worders, WorkOrder{AccountId: account_id})

	return recs, nil
}


func GetJob(db gorm.DB, job_id int) (*JobView, error) {

	row := new(JobView)
	res := db.Table(JOB_VIEW).Select(JOB_VIEW_COLS).Scan(row)
	if res.Error != nil {
		return nil, res.Error
	}
	return row, nil
}


