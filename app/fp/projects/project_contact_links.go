

package projects

import (
	"github.com/jinzhu/gorm"
)

// links of a contact to a project
type ProjectContactLink struct {
	ProjectContactLinkId int ` json:"project_contact_link_id" gorm:"column:project_contact_link_id; primary_key:yes" `
	ProjectId int ` json:"project_id" `
	ContactId int ` json:"contact_id" `
}

func (me ProjectContactLink) TableName() string {
	return "project_contact_links"
}

func DB_IndexProjectContactLinks(db gorm.DB) {

	cols := []string{"project_id", "contact_id"}

	for _, c := range cols {
		db.Model(&ProjectContactLink{}).AddIndex("idx_" + c, c)
	}
}

