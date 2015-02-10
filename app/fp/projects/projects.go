

package projects

import (

	"github.com/jinzhu/gorm"
)

// The Project, eg X2335 used in many parts
type Project struct {
	ProjectId int ` json:"project_id" gorm:"column:project_id; primary_key:yes" `
	AccountId int ` json:"account_id" `
	ProjectRef string   ` json:"project_ref" `
	ProjectDescription string   ` json:"project_description" sql:"type:varchar(255)" `
	OurProjectRef string `json:"our_project_ref"  `

	// Project string   ` json:"project" sql:"type:varchar(100)" `
}

func (me Project) TableName() string {
	return "projects"
}

func DB_IndexProjects(db gorm.DB) {

	cols := []string{"account_id", "project", "project_ref", "our_project_ref"}

	for _, c := range cols {
		db.Model(&Project{}).AddIndex("idx_" + c, c)
	}
}

// Database view extends the Project
type ProjectView struct {
	Project
	Company string ` json:"company" `
	Ticker string ` json:"ticker" `
	ModelsCount int ` json:"models_count" `
}

var PROJECT_VIEW = "v_projects"
var PROJECT_VIEW_COLS string = `
project_id, project_ref, project_description,
account_id, company, ticker,
models_count
`




func GetProjects(db gorm.DB) ([]*ProjectView, error) {

	rows := make([]*ProjectView, 0)
	db.Table(PROJECT_VIEW).Select(PROJECT_VIEW_COLS).Scan(&rows)
	if db.Error != nil {
		return nil, db.Error
	}
	return rows, nil
}


func GetAccountProjects(db gorm.DB, account_id int) ([]*ProjectView, error) {

	rows := make([]*ProjectView, 0)
	db.Table(PROJECT_VIEW).Select(PROJECT_VIEW_COLS).Where("account_id = ?", account_id).Scan(&rows)
	if db.Error != nil {
		return nil, db.Error
	}
	return rows, nil
}

func GetProject(db gorm.DB, project_id int) (*ProjectView, error) {

	row := new(ProjectView)
	db.Table(PROJECT_VIEW).Select(PROJECT_VIEW_COLS).Where("project_id = ?", project_id).Scan(row)
	if db.Error != nil {
		return nil, db.Error
	}
	return row, nil
}
