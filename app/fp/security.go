
package fp

import (
	"github.com/jinzhu/gorm"
)

type Security struct {
	SecurityId int ` json:"security_id"   gorm:"column:security_id; primary_key:yes" `
	Security string ` json:"security"  sql:"type:varchar(100)" `
	SecurityKey string ` json:"security_key"  sql:"type:varchar(100)" `
}
func (me Security) TableName() string {
	return "security"
}
func DB_IndexSecurity(db gorm.DB) {

	cols := []string{
		"security", "security_key"}

	for _, c := range cols {
		db.Model(&Security{}).AddIndex("idx_" + c, c)
	}
}
