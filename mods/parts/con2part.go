
package parts

import (


)



type Contact2Part struct {

	C2PId int 	` json:"c2p_id"  " gorm:"column:part_id; primary_key:yes" `
	PartId int ` json:"part_id" `
	ContactId int ` json:"contact_id" `

}


func (me Contact2Part) TableName() string {
	return "contact_2_part"
}
