
package fp

import (
	"github.com/jinzhu/gorm"
)

type SysLog struct {

	SysLogId int ` json:"sys_log_id" gorm:"column:sys_log_id; primary_key:yes"`
	UserId int ` json:"user_id" `

	AccountId int ` json:"account_id" `
	AddressId int ` json:"address_id" `
	ContactId int ` json:"contact_id" `
	OrderId int ` json:"order_id" `
	WorkOrderId int ` json:"work_order_id" `

	Action string ` json:"action" sql:"type:varchar(255);not null"`
	Log string ` json:"log" sql:"type:varchar(255); not null"`


}
func (me SysLog) TableName() string {
	return "sys_log"
}
func DB_IndexSysLog(db gorm.DB) {

	cols := []string{
		"account_id", "address_id", "contact_id",
		"order_id"	, "work_order_id", "action", "log"}

	for _, c := range cols {
		db.Model(&SysLog{}).AddIndex("idx_" + c, c)
	}
}
