package app

import (
	"github.com/revel/revel"
)



func SetupTemplates() {

	// returns the iso week number {{week .xdate}}
	revel.TemplateFuncs["week"] = func(a, b interface{}) string {
		return "55"
	}

	// returns the iso year to four digits {{year .xdate}}
	revel.TemplateFuncs["year"] = func(a, b interface{}) string {
		return "55"
	}




}
