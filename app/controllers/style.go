package controllers

import (
	"fmt"
	"github.com/revel/revel"
)


var Icons map[string]string

const (
	FAMFAM_SERVER_URL = "/icons/famfam"
)

func init(){

	Icons = make(map[string]string)

	// Accounts
	Icons["Account"] = "accept.png"

	Icons["Db"] = "database.png"
	Icons["DbAction"] = "database_go.png"

}

type Style struct {
	*revel.Controller
}

func (c Style) CssIcons() revel.Result {

	s := ""
	for class, file_name := range Icons {
		s += fmt.Sprintf(".ico%s{background-image: url(%s/%s) !important; background-repeat: no-repeat;}\n", class, FAMFAM_SERVER_URL, file_name)
	}
	//c.Response.ContentType = "test/css"
	return c.RenderText(s)
}

/*
s = ''
for k in sorted(icons.keys()):
s += ".%s{background-image: url('%s/%s') !important; background-repeat: no-repeat;}\n" % (k, ICON_SERVER_URL, icons[k])
s += "\n\n" # incase

return s
*/
