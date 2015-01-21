
package projects

// The end product model , eg X360, Steiger5, FooBar4

type Model struct {
	ModelId int ` json:"model_id" gorm:"column:model_id; primary_key:yes" `
	BrandId int ` json:"brand_id" `
	Model string   ` json:"model" `
}

func (me Model) TableName() string {
	return "brands"
}

func DB_IndexModels(db gorm.DB) {

	cols := []string{"brand_id", "model"}

	for _, c := range cols {
		db.Model(&Model{}).AddIndex("idx_" + c, c)
	}
}

