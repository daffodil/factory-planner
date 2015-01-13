/*global Ext: false, console: false, FP: false */

Ext.define("FP.calendar.WeeksViewGrid", {

extend: "Ext.grid.GridPanel",

get_store: function(){
	if(!this.xStore){
		this.xStore = Ext.create("Ext.data.Store", {
			deadmodel: "Account",
			proxy: {
				type: "ajax",
				url: '/ajax/accounts',
				method: "GET",
				reader: {
					type: "json",
					root: 'accounts'
				}
			},
			deadautoLoad: true,

			deadremoteSort: false,
			deadsortInfo: {
				field: "company",
				direction: 'ASC'
			}
		});
	}
	return this.xStore;
},

initComponent: function() {
	Ext.apply(this, {
        title: "Weeks",
        frame: false, plain: true, border: false,
        hideHeader: true,
        autoScroll: true,
        autoWidth: true,
        enableHdMenu: false,
        viewConfig: {
            emptyText: 'No weeks to view',
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

load: function(){
	//self.mask('Loading schedule');
	self.do_render = true;
    Ext.Ajax.request({
        scope: this,
        url: AJAX_SERVER + "/calendar/weeks",
        params: {

        },
        success: function(result){
            //console.log(result);

            var data = Ext.decode( result.responseText );

            var sto = this.get_store();

            var flds =  [];
            var headers = [];
            //console.log(data)
            var weeks = data.weeks;
            for(var i in weeks){
                var w = weeks[i];
                var yw = "wk_" + w.year + "_" + w.week
                flds.push({
                    name: yw,
                    type: "string"
                });

                var day_cols = [];
                var day_labels = ["Mn", "Tu", "We", "Th", "Fr", "Sa", "Su"];
                for(var d =0; d < 7; d++){
                    day_cols.push(
                        {text: day_labels[d], day: d, width: 20}
                    );
                }
                headers.push({
                    deaddataIndex: yw,
                    week: w,
                    header: "<b>" + w.week + "</b> <small> - " + w.year + "</small>",

                    menuDisabled: true,
                    columns: day_cols
                });
                console.log(w)
            }
            this.reconfigure(sto, headers);

        }

    });

}

}); // end class