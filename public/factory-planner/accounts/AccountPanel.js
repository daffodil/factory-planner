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

load_account: function(acc){
    Ext.Ajax.request({
        scope: this,
        url: AJAX_SERVER + "/account/" + acc.account_id,
        params: { },
        success: function(result){
            //console.log(result);

            var data = Ext.decode( result.responseText );
            console.log(data)
            return

            var sto = this.get_store();

            var flds =  [];
            var headers = [];
            //console.log(data)
            var weeks = data.weeks;

            var week_width = Math.round( (window.innerWidth - 30) / weeks.length );
            var day_width = week_width / 7;
            //console.log("weeks=", weeks.length, week_width, day_width);

            for(var i in weeks){
                var w = weeks[i];
                var yw = "wk_" + w.year + "_" + w.week
                flds.push({
                    name: yw,
                    type: "string"
                });

                headers.push({
                    dataIndex: yw,
                    week: w,
                    header: "<b>" + w.week + "</b><br><small>" + w.year + "</small>",
                    // DEAflex: 1,
                    width: week_width,
                    menuDisabled: true,
                });
                console.log(w)
            }
            this.reconfigure(sto, headers);



        }

    });

}

});