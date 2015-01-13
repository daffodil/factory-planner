/*global Ext: false, console: false, FP: false */

Ext.define("FP.calendar.PlannerPanel", {

extend: "Ext.tab.Panel",

get_weeks_view: function(){
	if(!this.WeeksView){
		this.WeeksView = Ext.create("FP.calendar.WeeksViewGrid", {});
    }
    return this.WeeksView;
},

initComponent: function() {
	Ext.apply(this, {

		layout: "fit",
		frame: false, plain: true, border: false,
		width: "100%",
		height: WIDGET_HEIGHT,
		items: [
            this.get_weeks_view()
            //Ext.create("FP.dev.RoutesBrowser", {})

		]
	});
	this.callParent();

	this.load();
}, // initComponent

load: function(){
    console.log("load")
    this.get_weeks_view().load();
}

});