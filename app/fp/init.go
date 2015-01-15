
package fp


import (
	"os"
	//"strings"
	//"net/url"
	//"github.com/jinzhu/gorm"

	//"github.com/daffodil/factory-planner/app/fp/accounts"
)




func RootPath() string {
	path := os.Getenv("GOPATH")
	path += "/src/github.com/daffodil/factory-planner"
	return path
}
