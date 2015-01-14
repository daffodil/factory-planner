

//================================================================
Ext.define("FP.dev.DbBrowser",  {

extend: "Ext.Panel",


//======================================================
// Tables Grid
grid_tables: function(){
	if(!this.gridTables){
		this.gridTables = Ext.create("Ext.grid.Panel", {
			region: 'center',
			title: "Tables &amp; Views",
			width: 240,
			frame: false,
			store:  new Ext.data.JsonStore({
				fields: [	
					{name: "table_name", type:"string"},
					{name: "is_view", type:"bool"},
					{name: "row_count", type:"int", useNull: true},
					{name: "next_id", type:"int", useNull: true},
					{name: "engine", type:"string", useNull: true}
				],
				idProperty: "table_name",
				
				proxy: {
					type: "ajax",
					url:  AJAX_SERVER + "/dev/db/tables",
					method: 'GET',
					reader: {
						type: "json",
						root: "tables"
					}
				},
				autoLoad: true
			}),
			viewConfig:{
				forceFit: true
			},
			columns: [ 
				{header: "Type", dataIndex: "is_view", flex:1, menuDisabled: true,
					renderer: function(val, meta, record, idx){
						sty =  "font-family: monospace;"
						if( val ){
							sty += "color: #009900;";
						}else{
							sty += "color: #000099;";
						}
						meta.style = sty;
						return val == true ? "View" : "Table";
					}
				},
				{header: "Table", dataIndex: "table_name", flex:4, menuDisabled: true,
					renderer: function(val, meta, record, idx){
						sty =  "font-weight: bold; font-family: monospace;";
						if( record.get("is_view") ){
							sty += "color: #009900;";
						}else{
							sty += "color: #000099;";
						}
						meta.style = sty;
						
						return val;
					}
				},
                {header: "Engine", dataIndex: "engine", flex: 2,
                    renderer: function(v, meta, record, idx) {
                        if (v) {
                            return v
                        }
                        return ""
                    }
                },
                {header: "Next ID", dataIndex: "next_id", align: "right", width: 50,
                    renderer: function(v, meta, record, idx) {
                        if (v === 0 || v > 0) {
                            return v
                        }
                        return ""
                    }
                },
                {header: "Rows", dataIndex: "row_count", align: "right", width: 50,
                    renderer: function(v, meta, record, idx) {

                        if (v === 0 || v > 0) {
                            return v
                        }
                        return ""
                    }
                }
			],
			listeners:{
				scope: this,
				selectionchange: function(grid, selection, e){
					
					if(selection.length == 0){
						return;
					}
					var rec = selection[0];
					var table_name = rec.get("table_name");
					var what = rec.get("is_view") == true ? "view" : "table";
					
					var url = AJAX_SERVER + "/dev/db/" + what + "/" + table_name;
					Ext.Ajax.request({
						url: url,
						method: "GET",
						scope: this,
						success: function(response, opts) {
							var data = Ext.decode(response.responseText);
							console.log(data.table);
							var sto = this.grid_columns().getStore()
							sto.loadData(data.table.columns);
							//console.log("sucksless", data);
							this.down("[name=definition_text]").setText(data.table.definition);
						},
						failure: function(response, opts) {
							console.log('server-side failure with status code ' + response.status);
						}
						
					});
				}
			}
		});
	}
	return this.gridTables;
},


//=================================================================
//== Columns
grid_columns: function(){
	if(!this.gridColumns){
		this.gridColumns = Ext.create("Ext.grid.Panel", {
			//region: 'east', 
			title: "Columns",
			flex: 1,
			store: Ext.create("Ext.data.Store", {
				fields: [	
					{name: "name", type:"string"},
					{name: "type", type:"string"},
					//{name: "max_char", type:"string"},
					{name: "nullable", type:"boolean"}
				],
				idProperty: "name",
				proxy: {
					type: "ajax",
					url: AJAX_SERVER + "/dev/db/table/_TABLE_NAME_/columns",
					method: 'GET',
					reader: {
						type: "json",
						root: "columns"
					}
				},
				
				autoLoad: false
				
			}),
			viewConfig:{
				forceFit: true,
				emptyText: "< Select a table",
				deferEmptyText: false
			},
			columns: [ 
				{header: "Column", dataIndex: "name", flex: 1, menuDisabled: true,
					renderer: function(val, meta, record, idx){
						meta.style = "font-weight: bold;"
						return val;
					}
				},
				
				{header: "Type", dataIndex: "type", flex: 1, menuDisabled: true},
				{header: "Nullable", dataIndex: "nullable", width: 50,
					renderer: function(v){
						return v ? "Yes" : "No";
					}
				}
			]
		});
	}
	return this.gridColumns;
},


on_select_view: function(butt, checked){
	//return
	if(checked){
		this.curr_database = butt.text;
		
		
		this.grid_columns().getStore().removeAll();
		
		this.grid_tables().getStore().getProxy().url = "/ajax/dev/db/tables";
		this.grid_tables().getStore().load();
	}
	butt.setIconCls( checked ? "icoOn" : "icoOff");
},


initComponent: function() {
	
	Ext.apply(this, {
		layout: 'border',
		fgxType: "DbBrowser",
		title: "DB Schema",
		iconCls: "icoDatabase",
		activeTab: 1,
		plain: true,
		frame: false,
		border: 0, bodyBorder: "",
		tbar: [
		    {text: "Create Views", handler: this.on_create_views, scope: this},
		    {text: "Update Searches", handler: this.on_update_searches, scope: this},
		],
		items: [
			this.grid_tables(),
			Ext.create("Ext.tab.Panel", {
				region: "east", flex: 2,
				items: [
					this.grid_columns(),
					{xtype: "panel", title: "Definition",
						items: [
							{xtype: "text", flex: 1, name: "definition_text", style: "font-family: monospace;", padding: 10}
						   
						]
					}
					
				]
			 })
		]
	
	});
	this.callParent();
	
},

//this.tablesGrid.load_tables();
load:  function(){
	this.grid_tables().getStore().load();
},

do_request: function(title, url, params){
    var d = Ext.create("FP.dev.ActionDialog", {title: title, url: url, params: params});
    d.run_show();
    return
    Ext.Ajax.request({
        scope: this,
        url: AJAX_SERVER + endpoint,
        params: xparams,
        success: function(result){
            var data = Ext.decode( result.responseText );
            if(data){
                Ext.Msg.show({
                    title: "Create Views",
                    msg: "Server replied",
                    value: result.responseText,
                    multiline: true,
                    icon: Ext.MessageBox.INFO,
                    width: window.innerWidth - 10,
                    height: window.innerHeight - 10,

                });
            }
            return data

        }
    });
},
on_create_views: function(){
    var data = this.do_request("Create SQL Views", "/dev/db/views/create", {});


},

on_update_searches: function(){
    var data = this.do_request("Update Searches", "/dev/db/update_searches", {});
}

});