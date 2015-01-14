/*global Ext: false, console: false, FP: false */

Ext.define("FP.calendar.PlannerPanel", {

extend: "Ext.tab.Panel",

get_weeks_view: function(){
	if(!this.xWeeksView){
		this.xWeeksView = Ext.create("FP.calendar.WeeksViewGrid", {});
    }
    return this.xWeeksView;
},
get_weeks_days_view: function(){
	if(!this.xWeeksDaysView){
		this.xWeeksDaysView = Ext.create("FP.calendar.WeeksDaysViewGrid", {});
    }
    return this.xWeeksDaysView;
},

initComponent: function() {
	Ext.apply(this, {

		layout: "fit",
		frame: false, plain: true, border: false,
		width: "100%",
		height: WIDGET_HEIGHT,
		items: [
            this.get_weeks_view(),
            this.get_weeks_days_view()
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