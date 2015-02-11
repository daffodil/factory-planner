package app

import (
	"github.com/revel/revel"
)



func SetupTemplates() {

	// returns the iso week number {{week .xdate}}
	revel.TemplateFuncs["week"] = func(a interface{}) string {
		return "55"
	}

	// returns the iso year to four digits {{year .xdate}}
	revel.TemplateFuncs["year"] = func(a interface{}) string {
		return "55"
	}

	// returns the js for ext_widgtet init
	// {{render_ext "FP.calendar.PlannerPanel"}}
	revel.TemplateFuncs["render_ext"] = func(widget string) string {
		s := "<div id='ext_widget'></div>\n"
		s += "<script>\n"
		s += "var Widget;\n"
		s += "Ext.onReady(function(){\n"
		s += "    Ext.Msg.wait('Loading .....');\n"
		s += "    Widget = Ext.create('" + widget + "', {renderTo: 'ext_widget'});\n"
		s += "    Ext.Msg.hide();\n"
		s += "});</script>"
		return s
	}


}
