
package projects

import (
	"github.com/jinzhu/gorm"
)


// The end product Brand , eg Volvo, Class, Jaguar
type Brand struct {
	BrandId int ` json:"brand_id" gorm:"column:brand_id; primary_key:yes" `
	AccountId int ` json:"account_id" `
	Brand string   ` json:"brand" `
}

func (me Brand) TableName() string {
	return "brands"
}

func DB_IndexBrands(db gorm.DB) {

	cols := []string{"account_id", "brand"}

	for _, c := range cols {
		db.Model(&Brand{}).AddIndex("idx_" + c, c)
	}
}

