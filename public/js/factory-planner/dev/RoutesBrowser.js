
function wrap_color(s, color){
	return "<span style='color: " + color + ";'>" + s + "</span>"
}
//================================================================
Ext.define("FP.dev.RoutesBrowser",  {

extend: "Ext.Panel",

//======================================================
grid_routes: function(){
	if(!this.gridRoutes){
		this.gridRoutes = Ext.create("Ext.grid.Panel", {
			region: 'center',
			title: "Urls",
			//width: 200,
			store:  new Ext.data.JsonStore({
				fields: [	
					{name: "url", type:"string"},
					{name: "controller", type:"string"},
					{name: "action", type:"string"}
				],
				idProperty: "url",
				
				proxy: {
					type: "ajax",
					url: AJAX_SERVER + "/dev/routes",
					method: 'GET',
					reader: {
						type: "json",
						root: "routes"
					}
				},
				autoLoad: true
			}),
			viewConfig:{
				forceFit: true
			},
			columns: [ 
				{header: "Controller", dataIndex: "controller", flex:1, menuDisabled: true},
				{header: "Action", dataIndex: "action", flex:2, menuDisabled: true,
					renderer: function(val, meta, record, idx){
						meta.style = "font-weight: bold; font-family: monospace;"
						parts = val.split(".")
						
						return "<b><span style='color: #009900;font-family: monospace;'>" + parts[0] + "</span>.<span style='color: #000099;font-family: monospace;'>" + parts[1] + "</span></b>"
					}
				},
				{header: "URL", dataIndex: "url", flex:4, menuDisabled: true,
					renderer: function(val, meta, record, idx){
						meta.style = "font-weight: bold; font-family: monospace;"
						var arr = [];
						var pp;
						var parts = val.split("/");
						//console.log(parts)
						if (parts[1] == "json") {
							arr.push( wrap_color(parts[1], "#47874B") );
							
						} else if (parts[1] == "xml") {
							arr.push( wrap_color(parts[1], "#474887") );
						}
						for(var i=2; i < parts.length; i++){
							pp = parts[i];
							if( pp[0] == ":"){
								arr.push( wrap_color(pp, "#CC2E1D") );
							}else{
								arr.push( wrap_color(pp, "#63166B") );	
							}
						}
						//console.log(arr)
						var ass = "/" + arr.join("/")
						return ass;
					}
				},
				
			],
			listeners:{
				scope: this,
				selectionchange: function(grid, selection, e){
					
				}
			}
		});
	}
	return this.gridRoutes;
},


initComponent: function() {
	
	Ext.apply(this, {
		layout: 'border',
		fgxType: "RoutesBrowser",
		title: "Url and Routes",
		iconCls: "icoRoutes",
		activeTab: 0,
		items: [
			this.grid_routes(),
		]
	
	});
	this.callParent();
	
},

load:  function(){
	this.grid_routes().getStore().load();
}


});  //end class