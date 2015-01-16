package orders

import (
	//"fmt"
	"time"
	"github.com/jinzhu/gorm"

	//"github.com/daffodil/factory-planner/mods/accounts"
)

type FPDate struct {
	time.Time
}

func (t *FPDate) MarshalJSON() ([]byte, error) {
	//ts := time.Time(*FPDate).Unix()
	stamp := "foo" //fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02"))

	return []byte(stamp), nil
}


type Order struct {

	OrderId int ` json:"order_id" gorm:"column:order_id; primary_key:yes"`
	LegerId int ` json:"leger_id" `
	OrderTypeId *int `json:"order_type_id" sql:"type:int"`
	AccountId int `json:"account_id" sql:"type:int"`
	PartId *int `json:"part_id" sql:"type:int;"`


	PurchaseOrder *string `json:"purchase_order" sql:"type:varchar(100)" `
	ClientExtraRef *string `json:"client_extra_ref" sql:"type:varchar(100)" `
	OrderNotes *string `json:"order_notes" sql:"type:varchar(100);not null;default:''" `

	OrderOrdered *time.Time `json:"order_ordered" sql:"type:date" `
	OrderRequired *time.Time `json:"order_required" sql:"type:date" `
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




type OrderView struct {
	Order
	OrderType string ` json:"order_type" `
	OrderColor string ` json:"order_color" `
	Company string ` json:"company"  `
	Ticker string ` json:"ticker"  `
}
var order_view_cols string = `
order_id, order_type_id, order_type, order_color,
client_order_no, account_id, company, ticker,
order_ordered, order_required
`

func GetOrders(db gorm.DB) ([]OrderView, error) {

	var orders []OrderView
	db.Table("v_orders").Select(order_view_cols).Scan(&orders)

	return orders, nil
}

func GetAccountOrders(db gorm.DB, account_id int) ([]OrderView, error) {

	var orders []OrderView
	db.Table("v_orders").Select(order_view_cols).Where("account_id=?", account_id).Scan(&orders)

	return orders, nil
}
