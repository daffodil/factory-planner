
package fp


import (
	"os"
	//"strings"
	//"net/url"
	"github.com/jinzhu/gorm"

	//"github.com/daffodil/factory-planner/app/fp/accounts"
)




func RootPath() string {
	path := os.Getenv("GOPATH")
	path += "/src/github.com/daffodil/factory-planner"
	return path
}


func GetSysInfo(db gorm.DB) map[string]interface{} {

	m := make(map[string]interface{})
	m["server_nick"] = "staff-play"
	m["server_host"] = "staff-play.domain"

	m["desktop_version"] = "115622"

	return m
}
