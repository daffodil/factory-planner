

package projects

// The Project, eg X2335 used in many parts

type Project struct {
	ProjectId int ` json:"project_id" gorm:"column:project_id; primary_key:yes" `
	AccountId int ` json:"account_id" `
	Project string   ` json:"project" `
}

func (me Project) TableName() string {
	return "brands"
}

func DB_IndexProjects(db gorm.DB) {

	cols := []string{"account_id", "project"}

	for _, c := range cols {
		db.Model(&Project{}).AddIndex("idx_" + c, c)
	}
}


