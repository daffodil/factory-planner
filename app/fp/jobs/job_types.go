package jobs


import (
	"github.com/jinzhu/gorm"
)


type JobType struct {
	JobTypeId int ` json:"job_type_id" gorm:"column:job_type_id; primary_key:yes"`
	JobType string ` json:"job_type"  sql:"type:varchar(100);not null;default:''" `
	JobTypeColor string ` json:"job_type_color" sql:"type:varchar(20)" `

}


func (me JobType) TableName() string {
	return "job_types"
}

func DB_IndexJobType(db gorm.DB) {

	cols := []string{
		"job_type"}

	for _, c := range cols {
		db.Model(&JobType{}).AddIndex("idx_" + c, c)
	}
}
func DB_CreateDefaultJobTypes(db gorm.DB) error {

	defaults := []JobType {
		JobType{JobTypeId: 100,  JobType: "Not Specified", JobTypeColor: "#aaaaaa"},
		JobType{JobTypeId: 200,  JobType: "Concept", JobTypeColor: "#FFFF99"},
		JobType{JobTypeId: 300,  JobType: "Prototype", JobTypeColor: "#CCFFCC"},
		JobType{JobTypeId: 400,  JobType: "Pre Volume", JobTypeColor: "#CCFFFF"},
		JobType{JobTypeId: 500,  JobType: "Production", JobTypeColor: "#FF99CC"},
	}

	var count int
	for _, rec := range defaults {
		db.Model(JobType{}).Where("order_type_id = ?", rec.JobTypeId).Count(&count)
		if count == 0 {
			db.Create(rec)
		}
	}
	return nil
}


func GetJobTypes(db gorm.DB) ([]JobType, error) {

	var rows  []JobType
	db.Order("job_type_id asc").Find(&rows)
	if db.Error != nil {
		return nil, db.Error
	}
	return rows, nil
}
