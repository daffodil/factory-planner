package orders

/*
NOW Project Types

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
func DB_CreateDefaultOrderTypes(db gorm.DB) error {

	defaults := []OrderType {
		OrderType{OrderTypeId: 100,  OrderType: "Not Specified", OrderColor: "#aaaaaa"},
		OrderType{OrderTypeId: 200,  OrderType: "Concept", OrderColor: "#FFFF99"},
		OrderType{OrderTypeId: 300,  OrderType: "Prototype", OrderColor: "#CCFFCC"},
		OrderType{OrderTypeId: 400,  OrderType: "Pre Volume", OrderColor: "#CCFFFF"},
		OrderType{OrderTypeId: 500,  OrderType: "Production", OrderColor: "#FF99CC"},
	}

	var count int
	for _, rec := range defaults {
		db.Model(OrderType{}).Where("order_type_id = ?", rec.OrderTypeId).Count(&count)
		if count == 0 {
			db.Create(rec)
		}
	}
	return nil
}
*/
