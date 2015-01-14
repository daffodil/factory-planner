
Ext.define('G2.accounts.AccountTree', {
	
extend: 'Ext.tree.Panel',
requires: [
	//'G2.lab_locations.LabLocationDialog'
],
loading: false,

initComponent: function(){
	
	Ext.apply(this, {
		
		useArrows: false,
		autoScroll: true,
		animate: false,
		enableDD: false,
		containerScroll: true,
		
		
		layout: "fit",
		
		rootVisible: false,
		//store: this.get_store(),
		listeners: { 
			scope: this,
			//selectionchange: this.on_selection_change
		},
		
	    columns: [
	        {	xtype: 'treecolumn',
	            text: 'Node',
	            flex: 2,
	            sortable: true,
	            dataIndex: 'lab_location', 
				renderer: this.r_node,
				menuDisabled: true
	        },{
	            text: 'Action',
	            flex: 1,
	            dataIndex: 'samples_count',
	            sortable: true,
				menuDisabled: true,
				align: "center",
				renderer: function(v){
					if(v > 0){
						return v;
					}
					return "";
				}
	        }
	   ]
		
		

	}); /* apply */
	this.callParent();
},  /* initComponent */

r_node: function(v){
	return "<b>" + v + "</b>";
},
r_locked: function(v){
	return v ? "Yes" : "-";
},
set_actions_enabled: function(ena){
	this.act_add().setDisabled(!ena);
	this.act_edit().setDisabled(!ena);
},

on_selection_change: function(model, selected, eOpts){
	console.log("on_selection");
	var has_sel = selected.length > 0
	this.set_actions_enabled(has_sel);
	if(has_sel){
		this.fireEvent("LAB_LOCATION", selected[0].getData());
	}else{
		this.fireEvent("LAB_LOCATION", false);
	}
},
act_add_top: function(){
	if(!this.xActionAddTop){
		this.xActionAddTop = Ext.create("Ext.Button",{
			text: "Add Root", mode: "add_root", iconCls: 'icoFolder',  
			listeners: {click: this.on_edit, scope: this}
		});
	}
	return this.xActionAddTop;
},
act_add: function(){
	if(!this.xActionAdd){
		this.xActionAdd = Ext.create("Ext.Button",{
			text: "Add Node", mode: "add_node", iconCls: 'icoLabLocationAdd',  disabled: true,
			listeners: {click: this.on_edit, scope: this}
		});
	}
	return this.xActionAdd;
},
act_edit: function(){
	if(!this.xActionEdit){
		this.xActionEdit = Ext.create("Ext.Button",{
			text: "Edit", mode: "edit", iconCls: 'icoLabLocationEdit', disabled: true,
			listeners: {
				scope: this,
				click: this.on_edit
			}
		});
	}
	return this.xActionEdit;
},
   
on_edit: function(butt){
	
	var rec = 0;
	var ll_id = 0;
	var p_id = 0;
	
	if (butt.mode == "add_root"){
		ll_id = 0;
		p_id = "root";
	
	}else if (butt.mode == "add_node"){
		ll_id = 0;
		rec = this.getSelectionModel().getSelection()[0]
		p_id = rec.get("ll_id");
	}else{
		//p_id = rec.get("parent_id");
		rec = this.getSelectionModel().getSelection()[0]
		ll_id = rec.get("ll_id");
		 //p_id = rec.get("parent_id");
	}
	
	var d = Ext.create("G2.lab_locations.LabLocationDialog", {
		ll_id: ll_id, 
		parent_id: p_id,
		listeners: {
			scope: this,
			do_update: function(data){
				console.log(data.parent_id);
				var pnode = this.get_store().getNodeById( data.parent_id );
				this.get_store().load({node: pnode});
			}
		}
	});
	d.run_show();
},


get_toolbar: function(){
	//if( this.g2_mode == "stores_browse" ) {
	//	return null;
	//}
	var arr = [];
	if( this.g2_mode == "admin"){
		arr.push( this.act_add_top() );
	}
	arr.push( this.act_add() );
	arr.push( this.act_edit() );
	
	arr.push( "->" );
	
	arr.push(
		{iconCls: 'icoRefresh2',
			listeners: { 
				scope: this,
				click: function(){
					this.set_actions_enabled(false);
					this.get_store().getRootNode().removeAll();
					this.fetch_stores();
					//this.get_store().load();
				}
			}
		}
	);
	return arr;
},



get_store: function(){
	if(!this.xTreeStore){
		
		var rootT = this.g2_mode == "stores"
					? {}
					: {nodeType: 'async',
						lab_location: 'Root',
						ll_id: 0, draggable: false, expanded: true,	id: 0
					}
		
		this.xTreeStore = Ext.create('Ext.data.TreeStore', {
			model: 'LabLocation',
			autoLoad: false,
			proxy: {
				type: 'ajax',
				url: '/rpc/lab_locations/tree_nodes',
				reader: {
					type: 'json',
					idProperty: 'll_id',
					root: 'lab_locations'
				}
			},
			root: rootT
		});
	}
	return this.xTreeStore;
},

fetch_stores: function(){
	Ext.Ajax.request({
		url : "/rpc/lab_locations/stores",
		method: 'GET',
		scope: this, 
		success: function(resp){	
			var data = Ext.decode(resp.responseText);
			this.load_lab_locations(data.lab_locations);
		}, 
		failure: function(){
			
			G2.msg("OOPS", "Some error");
		}
	})
},

load_lab_locations:function(tree){
	var pNode = this.get_store().getRootNode();
	
	for(var i= 0; i < tree.length; i++){
		this.append_node(pNode, tree[i] );
	}
	pNode.expand();
},

append_node: function(pNode, row){
	var nNode = pNode.appendChild(row);
	if(row.children){
		for(var i= 0; i < row.children.length; i++){
			var r = row.children[i];
			r.leaf = true;
			this.append_node(nNode, r );
		}	
		
	}
}

});
		