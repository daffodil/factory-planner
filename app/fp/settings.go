
package fp


import (
	"github.com/jinzhu/gorm"
)

type Setting struct {

	SettingId int ` json:"setting_id" gorm:"column:setting_id; primary_key:yes"`
	Key string ` json:"key" sql:"type:varchar(100);not null" `
	Value string ` json:"value" sql:"type:varchar(255);not null"`
	Description string ` json:"description" sql:"type:varchar(100)"`


}
func (me Setting) TableName() string {
	return "settings"
}
func DB_IndexSetting(db gorm.DB) {

	cols := []string{
		"key", "value"}

	for _, c := range cols {
		db.Model(&Setting{}).AddIndex("idx_" + c, c)
	}
}
