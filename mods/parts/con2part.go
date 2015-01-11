
package parts

import (


)



type Contact2Part struct {

	CPId int 	` json:"cp_id"  " gorm:"column:part_id; primary_key:yes" `
	PartId int ` json:"part_id" `
	ContactId int ` json:"contact_id" `

}


func (me Contact2Part) TableName() string {
	return "contacts_2_parts"
}
