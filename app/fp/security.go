
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


func DB_CreateDefaultSecurityLevels(db gorm.DB) error {

	security_defaults := []Security {
		Security{SecurityId: 20,  SecurityKey: "public", Security: "Public"},
		Security{SecurityId: 40,  SecurityKey: "client", Security: "Client"},
		Security{SecurityId: 60,  SecurityKey: "staff", Security: "Staff"},
		Security{SecurityId: 80,  SecurityKey: "supervisor", Security: "Supervisor"},
		Security{SecurityId: 100, SecurityKey: "manager", Security: "Manager"},
		Security{SecurityId: 120, SecurityKey: "administration", Security: "Administration"},
		Security{SecurityId: 200, SecurityKey: "director", Security: "Director"},
	}

	var count int
	for _, sec := range security_defaults {
		db.Model(Security{}).Where("security_id = ?", sec.SecurityId).Count(&count)
		if count == 0 {
			db.Create(sec)
		}
	}
	return nil
}
