package projects

import (
	"github.com/jinzhu/gorm"
)


type ProjectType struct {
	ProjectTypeId int ` json:"project_type_id" gorm:"column:project_type_id; primary_key:yes"`
	ProjectType string ` json:"project_type"  sql:"type:varchar(100);not null;default:''" `
	ProjectColor string ` json:"project_color sql:"type:varchar(20)" `

}


func (me ProjectType) TableName() string {
	return "project_types"
}

func DB_IndexProjectType(db gorm.DB) {

	cols := []string{
		"project_type"}

	for _, c := range cols {
		db.Model(&ProjectType{}).AddIndex("idx_" + c, c)
	}
}
func DB_CreateDefaultProjectTypes(db gorm.DB) error {

	defaults := []ProjectType {
		ProjectType{ProjectTypeId: 100,  ProjectType: "Not Specified", ProjectColor: "#aaaaaa"},
		ProjectType{ProjectTypeId: 200,  ProjectType: "Concept", ProjectColor: "#FFFF99"},
		ProjectType{ProjectTypeId: 300,  ProjectType: "Prototype", ProjectColor: "#CCFFCC"},
		ProjectType{ProjectTypeId: 400,  ProjectType: "Pre Volume", ProjectColor: "#CCFFFF"},
		ProjectType{ProjectTypeId: 500,  ProjectType: "Production", ProjectColor: "#FF99CC"},
	}

	var count int
	for _, rec := range defaults {
		db.Model(ProjectType{}).Where("project_type_id = ?", rec.ProjectTypeId).Count(&count)
		if count == 0 {
			db.Create(rec)
		}
	}
	return nil
}
