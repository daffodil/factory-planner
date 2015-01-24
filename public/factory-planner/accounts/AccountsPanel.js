/*global Ext: false, console: false, FP: false */

Ext.define("FP.accounts.AccountsPanel", {

extend: "Ext.Panel",

get_accounts_grid: function(){
	if(!this.xAccountsGrid){
		this.xAccountsGrid = Ext.create("FP.accounts.AccountsGrid", {
		    region: "center", flex: 1, height: WIDGET_HEIGHT
        });
		this.xAccountsGrid.on("account", this.load_account, this )
		    //console.log("got_account", account)
		//}, this)

    }
    return this.xAccountsGrid;
},

get_account_panel: function(){
	if(!this.xAccountPanel){
		this.xAccountPanel = Ext.create("FP.accounts.AccountPanel", {
		    region: "east", flex: 1, height: WIDGET_HEIGHT
		});
    }
    return this.xAccountPanel;
},


initComponent: function() {
	Ext.apply(this, {
		iconCls: "icoAccounts",
		title: "Accounts Portal",
		layout: "hbox",
		frame: false, plain: true, border: false,
		width: "100%",

		items: [
            this.get_accounts_grid(),
            this.get_account_panel()
		],
		/*
		tbar: [
			{xtype: 'buttongroup',
				title: 'Ticker',
				columns: 2,
				items: [
					{iconCls: "icoClr",	scope: this, tooltip: "Clear text box",
						handler: function(){
							var widget = this.down("textfield[name=search_ticker]");
							widget.setValue("");
							widget.focus();
						}
					},
					{xtype: "textfield",  name: "search_ticker",
						width: this.txt_width,
						enableKeyEvents: true,
						listeners: {
							scope: this,
							keyup: function(txtFld, e){
								txtFld.setValue( txtFld.getValue().trim() );
								var s = txtFld.getValue();
								if(s.length > 1){
									this.get_store().load({params: {
										ticker: s
									}});
								}
							}
						}
					}
				]
			},
			{xtype: 'buttongroup',
				title: 'Search',
				columns: 2,
				items: [
					{iconCls: "icoClr",	scope: this, tooltip: "Clear text box",
						handler: function(){
							var widget = this.down("textfield[name=search_account]");
							widget.setValue("");
							widget.focus();
						}
					},
					{xtype: "textfield",  name: "search_account",
						width: this.txt_width,
						enableKeyEvents: true,
						listeners: {
							scope: this,
							keyup: function(txtFld, e){
								if(txtFld.getValue().length > 3){
									var s = txtFld.getValue().trim();
									if(s.length > 3){
										this.get_store().load({params: {search: s}});
									}
								}
							}
						}
					}
				]
			}
		]
		*/

	});
	this.callParent();
}, // initComponent

load_account: function(acc) {
    this.get_account_panel().load_account(acc)
}
});