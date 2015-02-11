package controllers

import (
	"strings"
	"time"

	"github.com/revel/revel"

	"github.com/daffodil/factory-planner/app"
	//"github.com/daffodil/factory-planner/app/fp"
	"github.com/daffodil/factory-planner/app/fp/accounts"
)




func MakePayload() map[string]interface{} {
	payload := make(map[string]interface{})
	payload["success"] = true
	return payload
}

func GetS(params *revel.Params, key string) string {
	s := params.Get(key)
	s = strings.TrimSpace(s)
	return s
}




func checkUser (c *revel.Controller) revel.Result  {

	if c.Session["user"] == "" {
		p := MakePayload()
		p["error"] = "Not Logged In"
		return c.RenderJson(p)
	}

	return nil
}

func setContext(c *revel.Controller) revel.Result  {

	var e error
	c.RenderArgs["root"], e = accounts.GetRootAccount(app.Db)
	if e != nil {

	}
	c.RenderArgs["now"] = time.Now()
	return nil
}



/*
type Payload map[string]interface{}
func (me Payload)SetError(err interface{}) {

	switch v := err.(type) {

	case error:
		me["error"] = err.Error()

	case string:
		me["error"] = err

	default:
		me["error"] = err

	}
}
*/
