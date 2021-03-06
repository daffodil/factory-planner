package calendar


import (
	"time"
	//"github.com/revel/revel"

)

// Shifts start in two realms..
// Factory shift - eg car manufacturing

type Shift struct {

	// name of shift, eg afternoon
	Name string ` json:"name" `

	// start hour of shift
	Start int ` json:"start" `

	// length of shift in hours
	Length int ` json:"length" `

	// end hour of shift (this might roll over ??
	End int ` json:"end" `

	Week *int  ` json:"week" `
	Year *int ` json:"year" `

	UTCStart *time.Time  ` json:"time_start" `
	UTCEnd *time.Time ` json:"time_end" `
}

func NewShift(name string, hour_start int, length_hours int) *Shift{
	s := new(Shift)
	s.Name = "morning"
	s.Start = 6
	s.Length = 8
	s.End = hour_start + length_hours
	return s
}

var ThreeShifts map[string]*Shift

func init(){
	ThreeShifts = make(map[string]*Shift)
	ThreeShifts["morning"] = NewShift("morning", 6, 8)
	ThreeShifts["evening"] = NewShift("evening", 14, 8)
	ThreeShifts["night"] = NewShift("night", 20, 8)
}

// returns the default shift.. hard coded for now
// TODO make this dynamic
func GetShifts() (map[string]*Shift, error) {
	return ThreeShifts, nil
}

// returns a list of shifts for given week
func GetWeekShifts(year, week int) ([]*Shift, error) {

	//weekOb := WeekFromYearWeek(year, week)
	//weekOb.Setup(false)

	//start_date := weekOb.DateFirst
	shifts := make([]*Shift, 0)

	for idx := 0; idx < 7; idx++ {
		for sh := range ThreeShifts {
			//var ds =
			xShift := new(Shift)
			//{Name: ThreeShifts[sh].Name, Start: ThreeShifts[sh].Start}
			xShift.Name = sh
			xShift.Start = ThreeShifts[sh].Start
			xShift.Length = ThreeShifts[sh].Length
			xShift.End = ThreeShifts[sh].End
			shifts = append(shifts, xShift)
		}

	}
	//payload.Shifts = ThreeShifts
	return shifts, nil

}




