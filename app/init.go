package app

import (

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"

	"github.com/revel/revel"

	"github.com/daffodil/factory-planner/app/fp/accounts"
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}
	revel.OnAppStart( SetupTemplates )
	//InitDB()
	// register startup functions with OnAppStart
	// ( order dependent )
	revel.OnAppStart(InitDB)
	revel.OnAppStart(InitFP)
}

var Db gorm.DB

func InitDB(){
	var db_user, db_password, db_database, db_options string
	var found bool
	db_user, found = revel.Config.String("db.user")
	if !found {
		revel.ERROR.Printf("no db.user")
	}
	db_password, found = revel.Config.String("db.password")
	if !found {
		revel.ERROR.Printf("no db.password")
	}
	db_database, found = revel.Config.String("db.database")
	if !found {
		revel.ERROR.Printf("no db.database")
	}
	db_options, found = revel.Config.String("db.options")
	if !found {
		revel.ERROR.Printf("no db.options")
	}
	connect_str := db_user + ":" + db_password + "@/" + db_database + "?" + db_options
	revel.INFO.Printf("attempting connect with", connect_str)
	var err error
	Db, err = gorm.Open("mysql", connect_str)
	if err != nil {
		//todo throw tantrum
	}
	Db.LogMode(true)
	Db.DB().Ping()
}


func InitFP() {

	accounts.InitRoot(Db)

}















// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
