

package projects

import (
	//"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
)

// The Project, eg X2335 used in many parts
type ProjectModelLink struct {
	ProjectLinkId int ` json:"project_model_link_id" gorm:"column:project_model_link_id; primary_key:yes" `
	ProjectId int ` json:"project_id" `
	ModelId int ` json:"model_id" `
}

func (me ProjectModelLink) TableName() string {
	return "project_model_links"
}

func DB_IndexProjectModelLinks(db gorm.DB) {

	cols := []string{"project_id", "model_id"}

	for _, c := range cols {
		db.Model(&ProjectModelLink{}).AddIndex("idx_" + c, c)
	}
}

// Database view extends the Account with other stuff
type ProjectModelView struct {
	Model
	Brand string ` json:"brand" `

	AccountId int ` json:"account_id" `
	Ticker string ` json:"ticker" `
	Company string ` json:"company" `

	ProjectId int ` json:"project_id" `
	ProjectRef string ` json:"project_ref" `
}

var PROJECT_MODEL_VIEW string = "v_project_models"
var PROJECT_MODEL_VIEW_COLS string = `
model_id, model, brand_id, brand, account_id, company, ticker,
project_id, project_ref
`


// Brands for the account
func GetProjectModels(db gorm.DB, project_id int) ([]*ProjectModelView, error) {
	var rows []*ProjectModelView
	db.Table(PROJECT_MODEL_VIEW).Select(PROJECT_MODEL_VIEW_COLS).Where("project_id = ?", project_id).Scan(&rows)
	return rows, nil
}

func GetProject2ModelsLookup(db gorm.DB) (map[string][]ProjectModelView, error) {
	var rows []ProjectModelView
	db.Table(PROJECT_MODEL_VIEW).Select(PROJECT_MODEL_VIEW_COLS).Scan(&rows)

	lookup := make(map[string][]ProjectModelView)
	for _, r := range rows {
		pid := strconv.Itoa(r.ProjectId)
		_, exist := lookup[pid]
		if exist == false {
			lookup[pid] = make([]ProjectModelView, 0)
		}
		lookup[pid] = append(lookup[pid], r)

	}
	//fmt.Println( lookup)
	return lookup, nil
}


func GetProject2Models(db gorm.DB) ([]ProjectModelView, error) {
	var rows []ProjectModelView
	db.Table(PROJECT_MODEL_VIEW).Select(PROJECT_MODEL_VIEW_COLS).Scan(&rows)
	return rows, nil
}
func GetProject2ModelLinks(db gorm.DB) ([]ProjectModelLink, error) {
	var rows []ProjectModelLink
	db.Table(PROJECT_MODEL_VIEW).Select(PROJECT_MODEL_VIEW_COLS).Scan(&rows)
	return rows, nil
}
