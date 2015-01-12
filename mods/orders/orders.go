package orders

import (
	"time"
	"github.com/jinzhu/gorm"
)



type Order struct {

	OrderId int `json:"order_id" gorm:"column:order_id; primary_key:yes"`
	OrderTypeId int `json:"order_type_id" sql:"not null;`
	AccountId int `json:"account_id" sql:"not null;`
	PartId int `json:"part_id" sql:"not null;`


	ClientOrderNo string `json:"client_order_no" sql:"type:varchar(100);not null;default:''" `
	OrderNotes string `json:"order_notes" sql:"type:varchar(100);not null;default:''" `

	OrderOrdered time.Time `json:"order_ordered" sql:"type:date" `
	OrderRequired time.Time `json:"order_required" sql:"type:date" `

}

func (me Order) TableName() string {
	return "orders"
}
func DB_IndexOrder(db gorm.DB) {

	cols := []string{
		"order_type_id", "account_id", "part_id", "order_required","client_order_no"}

	for _, c := range cols {
		db.Model(&Order{}).AddIndex("idx_" + c, c)
	}
}
