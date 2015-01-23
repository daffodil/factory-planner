

package projects

import (
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

