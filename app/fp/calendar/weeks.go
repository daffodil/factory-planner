

package calendar

import (
	//"fmt"
	"time"
	"errors"
)

const(
	LAST_WEEK = "last_week"
	WEEK_BEFORE = "week_before"

	THIS_WEEK = "this_week"

	NEXT_WEEK = "next_week"
	WEEK_AFTER = "week_after"
)
var errInvalidView error = errors.New("Invalid View")


// Get the date for first day of the week
// - thanks to
// http://stackoverflow.com/questions/18624177/go-unix-timestamp-for-first-day-of-the-week-from-iso-year-week
func FirstDayOfISOWeek(year int, week int) time.Time {
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
	DateFirst string `json:"date_first,omitempty"`
	DateLast string `json:"date_last,omitempty"`
	WeekLast *Week `json:"week_last,omitempty"`
	WeekNext *Week `json:"week_next,omitempty"`
}

func (me *Week) Setup(inc_weeks bool) {

	first_date := FirstDayOfISOWeek(me.Year, me.Week)
	me.DateFirst = ToString(first_date)

	last_date := first_date.AddDate(0, 0, 6)
	me.DateLast = ToString(last_date)

	if inc_weeks {
		lastwk := first_date.AddDate(0, 0, -7)
		me.WeekLast = new(Week)
		me.WeekLast.Year, me.WeekLast.Week = lastwk.ISOWeek()
		me.WeekLast.Setup(false)

		nextwk := first_date.AddDate(0, 0, 7)
		me.WeekNext = new(Week)
		me.WeekNext.Year, me.WeekNext.Week = nextwk.ISOWeek()
		me.WeekNext.Setup(false)
	}
}

func NewWeek(year, week int) Week {
	w := Week{}
	w.Year = year
	w.Week = week
	return w
}

func WeekFromView(view string) (*Week, error) {
	d := Now()
	switch view {

	case WEEK_BEFORE:
		d = d.AddDate(0, 0, -14)

	case LAST_WEEK:
		d = d.AddDate(0, 0, -7)

	case NEXT_WEEK:
		d = d.AddDate(0, 0, 7)

	case WEEK_AFTER:
		d = d.AddDate(0, 0, 14)

	case THIS_WEEK:

	default:
		return nil, errInvalidView
	}

	w := WeekFromTime(d, true)
	return w, nil
}


func WeekFromTime(d time.Time, inc_prev_next_weeks bool) *Week {

	w := new(Week)
	w.Year, w.Week = d.ISOWeek()
	w.Setup(inc_prev_next_weeks)
	return w
}
func WeekFromYearWeek(year, week int) *Week {

	d := FirstDayOfISOWeek(year, week)
	w := WeekFromTime(d, false)
	//w.Year, w.Week = d.ISOWeek()
	//w.Setup(inc_prev_next_weeks)
	return w
}

// Returns a list of week, previious next eg 1, 4 = include one before, and two after
func WeeksView(year, week, prev, ahead int) ([]*Week, error) {

	weeks := make([]*Week, 0)


	now_date := Now()

	for i := prev; i < ahead; i++ {
		d := now_date.AddDate(0, 0, i * 7)
		weeks = append(weeks, WeekFromTime(d, false))
	}

	return weeks, nil

}



// Returns a list of week, previious next eg 1, 4 = include one before, and two after
/*
func DEADGetWeek(year, week int) ([]*Week, error) {

	weeks := make([]*Week, 0)

	now_date := Now()

	for i := prev; i < ahead; i++ {
		d := now_date.AddDate(0, 0, i * 7)
		weeks = append(weeks, WeekFromTime(d, false))
	}

	return weeks, nil

}
*/
