

package projects

import (
	"github.com/jinzhu/gorm"
)

// The Project, eg X2335 used in many parts
type Project struct {
	ProjectId int ` json:"project_id" gorm:"column:project_id; primary_key:yes" `
	AccountId int ` json:"account_id" `
	ProjectNo string   ` json:"project_no" `
	Project string   ` json:"project" `
	OurProjectNo string `json:"our_project_no" sql:"type:varchar(100)" `
}

func (me Project) TableName() string {
	return "projects"
}

func DB_IndexProjects(db gorm.DB) {

	cols := []string{"account_id", "project"}

	for _, c := range cols {
		db.Model(&Project{}).AddIndex("idx_" + c, c)
	}
}




