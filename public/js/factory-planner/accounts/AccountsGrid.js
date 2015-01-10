/*global Ext: false, console: false, FP: false */

Ext.define("FP.accounts.AccountsGrid", {

extend: "Ext.grid.GridPanel",
get_store: function(){
	if(!this.xStore){
		this.xStore = Ext.create("Ext.data.JsonStore", {
			model: "Account",
			proxy: {
				type: "ajax",
				url: '/ajax/accounts',
				method: "GET",
				reader: {
					type: "json",
					root: 'accounts'
				}
			},
			autoLoad: true,

			remoteSort: false,
			sortInfo: {
				field: "company",
				direction: 'ASC'
			}
		});
	}
	return this.xStore;
},

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
			store: this.get_store(),
			loadMask: true,
			tbar: [
				{text: "New"},{text: "View"},
			],
			columns: [
				{header: 'Ticker', dataIndex:'ticker',
					sortable: true, width: 100, menuDisabled: true,
					renderer: function(v, meta, rec){
						return v;
					}
				},
				{header: 'Account', dataIndex:'company',
                    sortable: true, flex: 3, menuDisabled: true,
                    renderer: function(v, meta, rec){
                        return "<b>" + v + "</b>";
                    }
                },
                {header: 'On Hold', dataIndex:'on_hold'},
                {header: 'Client', dataIndex:'is_client'},
                {header: 'Supplier', dataIndex:'is_supplier'},
                {header: 'Acc Ref', dataIndex:'acc_ref'},
                {header: "Active", dataIndex: 'acc_active'}
			]

    });
    this.callParent();
}, // initComponent


});