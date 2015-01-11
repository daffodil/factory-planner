package orders

type Order struct {

	OrderId int `json:"order_id" gorm:"column:order_id; primary_key:yes"`
	OrderTypeId int `json:"order_type_id" sql:"not null;`
	AccountId int `json:"account_id" sql:"not null;`
	PartId int `json:"part_id" sql:"not null;`


	ClientOrderNo string `json:"client_order" sql:"type:varchar(100);not null;default:''" `
	OrderNotes string `json:"order_notes" sql:"type:varchar(100);not null;default:''" `

	OrderRequired string `json:"order_required" sql:"type:varchar(10);not null;default:''" `

}

func (me Order) TableName() string {
	return "orders"
}
