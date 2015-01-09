
package cal

import (
	"time"
	"github.com/revel/revel"

)

// We Purposely not create a NOW()
// but for template design it is local
// for system it is utc..

// We need ALERTS of when bst kicks in and this will
// make shift shorter and longer by one hours
// need to fire off events for sending wake up messages to crew
type Ticker struct {
	// current time
	UTC time

}


// send out ws message with tick change
func (me Ticker) TickMinute(){


}
// send out ws message with tick change of 15 mins
func (me Ticker) Tick15min(){


}
// send out ws message with tick change of 10 mins
func (me Ticker) Tick10min(){


}
// send out ws message with tick change of 5 mins
func (me Ticker) Tick5min(){


}

// we use calendar/ in code, `cal` is avoided for clash with  calibration/
type Calendar struct {


}


var MasterClock = *Ticker

func StartMasterClock() {

	// add ticker and start clock
	MasterClock = new(Ticker)
	MasterClock.Start()

	// load holidays and days not work
	Calendar := new(Calendar)
	Calendar.SetShist()

	// load now and find current state from log

	// load caches for lookups

	// load cron hourly
	// view jobs

}


func Now() {
	return time.Now.UTC()
}

func Week() int {
	return 3
}

func init(){
	// need to init things

}


func SetupTemplates() {

	// returns the iso week number {{week .xdate}}
	revel.TemplateFuncs["week"] = func(a, b interface{}) string {
		return "55"
	}

	// returns the iso year to four digits {{year .xdate}}
	revel.TemplateFuncs["year"] = func(a, b interface{}) string {
		return "55"
	}

}
