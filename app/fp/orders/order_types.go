package orders

import (
	"github.com/jinzhu/gorm"
)


type OrderType struct {
	OrderTypeId int ` json:"order_type_id" gorm:"column:order_type_id; primary_key:yes"`
	OrderType string ` json:"order_type"  sql:"type:varchar(100);not null;default:''" `
	OrderColor string ` json:"order_color sql:"type:varchar(20)" `

}

func (me OrderType) TableName() string {
	return "order_types"
}

func DB_IndexOrderType(db gorm.DB) {

	cols := []string{
		"order_type"}

	for _, c := range cols {
		db.Model(&OrderType{}).AddIndex("idx_" + c, c)
	}
}
