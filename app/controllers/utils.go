package controllers

import (
	"strings"
	"github.com/revel/revel"
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
