/*global Ext: false, console: false, FGx: false */

Ext.define("FP.dev.DevPanel", {

extend: "Ext.tab.Panel",



initComponent: function() {
	Ext.apply(this, {

		layout: "fit",
		frame: false, plain: true, border: false,
		width: "100%",
		height: WIDGET_HEIGHT,
		items: [
            Ext.create("FP.dev.RoutesBrowser", {}),
            Ext.create("FP.dev.DbBrowser", {})
		]
	});
	this.callParent();
}, // initComponent


});