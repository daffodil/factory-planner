

package calendar

import (
	"time"
)

const(


	LAST_WEEK = "last_week"
	WEEK_BEFORE = "week_before"

	THIS_WEEK = "this_week"

	NEXT_WEEK = "next_week"
	WEEK_AFTER = "week_after"

)


func firstDayOfISOWeek(year int, week int) time.Time {
	UTC, errloc := time.LoadLocation("UTC")
	if errloc != nil {

	}
	date := time.Date(year, 0, 0, 0, 0, 0, 0, UTC)
	isoYear, isoWeek := date.ISOWeek()
	for date.Weekday() != time.Monday { // iterate back to Monday
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoYear < year { // iterate forward to the first day of the first week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoWeek < week { // iterate forward to the first day of the given week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	return date
}


type Week struct {
	Year int `json:"year"`
	Week int `json:"week"`
	DateFirst string `json:"date_first"`
	DateLast string `json:"date_last"`
}

func (me Week)Setup() {
	d := firstDayOfISOWeek(me.Year, me.Week)
	me.DateFirst = ToString(d)
}

func NewWeek(year, week int) Week {
	w := Week{}
	w.Year = year
	w.Week = week
	return w
}

func WeekFromView(view string) Week {
	d := Now()
	switch view {
	case LAST_WEEK:
		//

	case THIS_WEEK:

	}

	w := WeekFromTime(d)
	return w
}
func WeekFromTime(d time.Time) Week {

	w := Week{}
	w.Year, w.Week = d.ISOWeek()
	w.Setup()
	return w
}

func (me Week) DEADStart() Week {
	w := NewWeek(2015, 20)
	return w
}
