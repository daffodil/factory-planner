
var WIDGET_WIDTH = window.innerWidth - 20;
var WIDGET_HEIGHT = window.innerHeight - 80;


var API_DATE_FORMAT = 'Y-m-d';
var API_TIME_FORMAT = 'H:i:s';
var API_DATE_TIME_FORMAT = 'Y-m-d H:i:s';


 Ext.Loader.setConfig({
        enabled : true,
        paths   : {
			Ext: "/public/ext-4.2.1.883",
			FP : '/public/js/factory-planner'
        }
});

Ext.QuickTips.init();
Ext.Ajax.disableCaching = false;


Ext.define('Account', {
	extend: 'Ext.data.Model',
	idProperty: 'account_id',
	fields: [
		{name: "account_id", type: 'int'},
		{name: "ticker", type: 'string'},
		{name: "company", type: 'string'},
		{name: "acc_ref", type: 'string'},
		{name: "is_supplier", type: 'bool'},
		{name: "is_client", type: 'bool'},
		{name: "acc_active", type: 'bool'}
	]
});
