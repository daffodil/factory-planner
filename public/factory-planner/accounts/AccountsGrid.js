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
    var ww = 40;
	Ext.apply(this, {
        frame: false, plain: true, border: false,
        hideHeader: true,
        autoScroll: false,
        autoWidth: true,
        enableHdMenu: false,
        layout: "fit",
        viewConfig: {
            emptyText: 'No accounts in view',
            deferEmptyText: false,
            forceFit: true
        },
        stripeRows: true,

        store: this.get_store(),
        loadMask: true,

        tbar: [
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
        ],
        columns: [
            {header: 'Ticker', dataIndex:'ticker',
                sortable: true, width: 80, menuDisabled: true,
                renderer: function(v, meta, rec){
                    return v;
                }
            },
            {header: 'Orders', dataIndex:'orders_due',
                width: ww, menuDisabled: true,  align: "center",
                renderer: function(v, meta, rec){
                    if(v > 0){
                        meta.style = "font-weight: bold;"
                        return v
                    }
                    return "-"
                }
            },
            {header: 'Account', dataIndex:'company',
                sortable: true, flex: 3, menuDisabled: true,
                renderer: function(v, meta, rec){
                    var s = "";
                    if( rec.get("acc_active") ){
                        s += "font-weight: bold;"
                        if( rec.get("root") ){
                           s += "color: red;";
                        } else {
                           s += "color: #000088;";
                        }
                    }

                    if (s != ""){
                        meta.style = s;
                    }
                    console.log(s);
                    return v;
                }
            },
            {header: 'Hold', dataIndex:'on_hold', width: ww, menuDisabled: true, renderer: this.render_yn, align: "center"},
            {header: 'Client', dataIndex:'is_client', width: ww, menuDisabled: true, renderer: this.render_yn, align: "center"},
            {header: 'Supplier', dataIndex:'is_supplier', width: ww, menuDisabled: true, renderer: this.render_yn, align: "center"},
            {header: 'Acc Ref', dataIndex:'acc_ref'},
            {header: "Active", dataIndex: 'acc_active', width: ww, menuDisabled: true, ssrenderer: this.render_yn, align: "center"}
        ],
        bbar: Ext.create('Ext.PagingToolbar', {
                store: this.get_store(),
                displayInfo: true,
                displayMsg: 'Displaying accounts {0} - {1} of {2}',
                emptyMsg: "No accounts to display",
                items:[
                    '-',
                    {
                        text: 'Show Preview',
                        //pressed: pluginExpanded,
                        enableToggle: true
                    }
                ]
        }),
        listeners: {
            scope: this,
            //itemclick: this.on_grid_click,
            //itemdblclick: this.on_grid_dblclick,
            selectionchange: this.on_grid_selection_change
        }

    });
    this.callParent();
}, // initComponent

on_grid_selection_change: function(view, selections, opts){
    console.log(selections);
    if (selections.length > 0 ){
        this.fireEvent("account", selections[0].getData())
    } else {
        this.fireEvent("account", null)
    }
},

render_yn: function(v){
    if (v){
        return "Y";
    }
    return "-";
}

});