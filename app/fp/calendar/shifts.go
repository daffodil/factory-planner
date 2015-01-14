package calendar


import (
	//"time"
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


type ShiftsPayload struct {
	Success bool `json:"success"`
	Shifts map[string]*Shift  `json:"shifts"`
}

// returns the default shift.. hard coded for now
// TODO make this dynamic
func GetShifts() (*ShiftsPayload, error) {

	payload := new(ShiftsPayload)
	payload.Shifts = ThreeShifts
	return payload, nil

}

