
var AJAX_SERVER = "/ajax";

//var WIDGET_WIDTH = window.innerWidth - 20;
var WIDGET_HEIGHT = window.innerHeight - 120;


var API_DATE_FORMAT = 'Y-m-d';
var API_TIME_FORMAT = 'H:i:s';
var API_DATE_TIME_FORMAT = 'Y-m-d H:i:s';


 Ext.Loader.setConfig({
        enabled : true,
        paths   : {
			Ext: "/public/ext-4.2.1-gpl",
			FP : '/public/factory-planner'
        }
});

Ext.QuickTips.init();
Ext.Ajax.disableCaching = false;


Ext.define('Account', {
	extend: 'Ext.data.Model',
	idProperty: 'account_id',
	fields: [
		{name: "account_id", type: 'int'},
		{name: "acc_active", type: 'bool'},
		{name: "root", type: 'bool'},
		{name: "ticker", type: 'string'},
		{name: "company", type: 'string'},
		{name: "acc_ref", type: 'string'},
		{name: "is_supplier", type: 'bool'},
		{name: "on_hold", type: 'bool'},
		{name: "is_client", type: 'bool'},
		{name: "orders_due", type: 'int'}
	]
});
Ext.define('Order', {
	extend: 'Ext.data.Model',
	idProperty: 'order_id',
	fields: [
		{name: "order_id", type: 'int'},
		{name: "purchase_order", type: 'string'},
		{name: "client_extra_ref", type: 'string'},
		{name: "ticker", type: 'string'},
		{name: "company", type: 'string'},
		{name: "work_orders", type: 'int'},
		{name: "order_type_id", type: 'int'},
		{name: "order_type", type: 'string'},
		{name: "order_color", type: 'string'},
		{name: "order_ordered", type: 'date'},
		{name: "order_required", type: 'date'}
	]
});