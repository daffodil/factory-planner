package orders

type WorkOrder struct {

	WorkOrderId int `json:"work_order_id" gorm:"column:work_order_id; primary_key:yes"`
	OrderId int `json:"order_id" sql:"type:int(11);not null;" `


	WorkOrderNo string `json:"work_order_no" sql:"type:varchar(100);not null;default:''" `
	WorkOrderNotes string `json:"work_order_notes" sql:"type:varchar(255);not null;default:''" `

	WorkOrderRequired string `json:"work_order_required" sql:"type:varchar(10);not null;default:''" `
}

func (me WorkOrder) TableName() string {
	return "works_orders"
}
