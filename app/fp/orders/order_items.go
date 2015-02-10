package orders

import (
	//"fmt"
	//"time"
	"github.com/jinzhu/gorm"

	//"github.com/daffodil/factory-planner/mods/accounts"
)


type OrderItem struct {
	OrderItemId int `json:"order_item_id" gorm:"column:order_item_id; primary_key:yes"`
	OrderId int `json:"order_id" sql:"type:int"`
	ProjectId int `json:"project_id" sql:"type:int"`
	OrderItemQty int `json:"qty_ordered" sql:"type:int"`
	OrderItemDescription int `json:"order_item_description" sql:"type:varchar(255)"`
	OrderItemImport int `json:"order_item_import" sql:"type:varchar(255)"`
}


func (me OrderItem) TableName() string {
	return "order_items"
}
func DB_IndexOrderItems(db gorm.DB) {

	cols := []string{
		"order_item_id", "order_id", "project"}

	for _, c := range cols {
		db.Model(&OrderItem{}).AddIndex("idx_" + c, c)
	}
}
