

package projects

import (
	"github.com/jinzhu/gorm"
)

// The Project, eg X2335 used in many parts
type ProjectLinks struct {
	ProjectId int ` json:"project_id" `
	ModelId int ` json:"model_id" `
}

func (me ProjectLinks) TableName() string {
	return "project_links"
}

func DB_IndexProjectLinks(db gorm.DB) {

	cols := []string{"project_id", "model_id"}

	for _, c := range cols {
		db.Model(&ProjectLinks{}).AddIndex("idx_" + c, c)
	}
}

