/*global Ext: false, console: false, FGx: false */

Ext.define("FP.accounts.AccountsPanel", {

extend: "Ext.Panel",



initComponent: function() {
	Ext.apply(this, {
		iconCls: "icoAccounts",
		title: "Accounts",
		layout: "border",
		frame: false, plain: true, border: false,
		width: "100%",
		height: WIDGET_HEIGHT,
		items: [
            Ext.create("FP.accounts.AccountsGrid", {region: "center"})
            //this.get_runways_tree()
		],
		tbar: [
			{xtype: 'buttongroup',
				title: 'Find Ident',
				columns: 2,
				items: [
					{iconCls: "icoClr",	scope: this, tooltip: "Clear text box",
						handler: function(){
							var widget = this.down("textfield[name=search_apt_ident]");
							widget.setValue("");
							widget.focus();
						}
					},
					{xtype: "textfield",  name: "search_apt_ident",
						width: this.txt_width,
						enableKeyEvents: true,
						listeners: {
							scope: this,
							keyup: function(txtFld, e){
								txtFld.setValue( txtFld.getValue().trim() );
								var s = txtFld.getValue();
								if(s.length > 1){
									this.get_airports_store().load({params: {
										ident: s,
										//apt_type: this.get_apt_types(),
										//apt_size: this.get_apt_sizes()
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
							var widget = this.down("textfield[name=search_apt_text]");
							widget.setValue("");
							widget.focus();
						}
					},
					{xtype: "textfield",  name: "search_apt_text",
						width: this.txt_width,
						enableKeyEvents: true,
						listeners: {
							scope: this,
							keyup: function(txtFld, e){
								if(txtFld.getValue().length > 3){
									var s = txtFld.getValue().trim();
									if(s.length > 3){
										this.get_airports_store().load({params: {search: s}});
									}
								}
							}
						}
					}
				]
			}
		]

	});
	this.callParent();
}, // initComponent


});