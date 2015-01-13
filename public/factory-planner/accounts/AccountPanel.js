/*global Ext: false, console: false, FP: false */

Ext.define("FP.accounts.AccountPanel", {

extend: "Ext.Panel",


get_account_tree: function(){
	if(!this.xAccountPanel){
		this.xAccountPanel = Ext.create("FP.accounts.AccountTree", {});
    }
    return this.xAccountPanel;
},

initComponent: function() {
	Ext.apply(this, {
		iconCls: "icoAccounts",
		title: "Account",
		layout: "vbox",
		frame: false, plain: true, border: false,
		items: [
            //this.get_account_tree()
           // this.get_account_panel()
		],
		tbar: [
            /*
			 */
		]

	});
	this.callParent();
}, // initComponent


});