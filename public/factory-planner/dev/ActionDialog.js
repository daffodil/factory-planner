

//================================================================
Ext.define("FP.dev.ActionDialog",  {

extend: 'Ext.window.Window',

get_form: function(){

	if( !this._form){
		this._form = Ext.create("Ext.form.Panel", {
			url: this.url,
			waitMsgTarget: true,
			//ssreader: this.frmReader,
			defaults: {
				labelAlign: 'right'
			},
		  	bodyStyle: 'padding: 20px;',
			frame: true,
			items:[
				{fieldLabel: 'URL', xtype: 'textfield', name: 'url',  anchor: '90%', value: this.url},
				{fieldLabel: 'Reply', xtype: 'textarea', name: 'reply',
				    height: window.innerHeight - 150,
					allowBlank: false, msgTarget: 'side', anchor: '90%'
				},

			],
			buttons:[
			     {text: 'Cancel', iconCls: "icoBlack",
					listeners: {
						scope: this,
						click: function(){
								this.close();
							}
					}
			     },
				this.save_button()
			],
			buttonAlign: 'right'
		});
	}
	return this._form;
},

save_button: function(){
	if(!this._save_button){
		this._save_button = Ext.create("Ext.Button",{
			text: 'Create Views', iconCls: "icoSave",
			listeners: {
			    scope: this,
				click: this.send_request
			}
		});
	}
	return this._save_button;
},
run_show: function(){
	this.show();
	//this.send_request();
},
send_request: function(){
	this.show();
	Ext.Ajax.request({
		scope: this,
		url : AJAX_SERVER + this.url,
		params: this.params,
		method: 'GET',
		success: function(resp){
			var data = Ext.decode(resp.responseText);
			console.log(data)
            this.get_form().getForm().findField("reply").setValue(resp.responseText);
		}
	});
},


initComponent: function(){


	Ext.apply(this, {

		//title: "",
		//iconCls: this.ll_id > 0 ? 'icoLabLocationEdit' : 'icoLabLocationAdd',
		width: window.innerWidth - 50,


		layout: "fit",
		items:[
		    this.get_form()
        ]


	}); /* apply */
	this.callParent();
	//this.doLayout();
},  /* initComponent */

//this.tablesGrid.load_tables();
load:  function(){
	this.grid_tables().getStore().load();
},

do_request: function(endpoint, xparams){
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
}

});  // end function cinstructor