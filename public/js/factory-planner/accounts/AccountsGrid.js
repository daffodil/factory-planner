/*global Ext: false, console: false, FGx: false */

Ext.define("FP.accounts.AccountsGrid", {

extend: "Ext.grid.GridPanel",

initComponent: function() {
	Ext.apply(this, {
            frame: false, plain: true, border: false,
			hideHeader: true,
			autoScroll: true,
			autoWidth: true,
			enableHdMenu: false,
			viewConfig: {
				emptyText: 'No accounts in view',
				deferEmptyText: false,
				forceFit: true
			},

			stripeRows: true,
			//store: this.get_airports_store(),
			loadMask: true,
			tbar: [
				//this.action_new_tab()
			],
			columns: [
				{header: 'Account', dataIndex:'company',
					sortable: true, flex: 1, menuDisabled: true,
					renderer: function(v, meta, rec){
						return v //rec.get("ident") + ": " + rec.get("name");
					}
				}
			]

    });
    this.callParent();
}, // initComponent


});