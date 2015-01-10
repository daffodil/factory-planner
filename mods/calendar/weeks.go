

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



type Week struct {
	Year int `json:"year"`
	Week int `json:"week"`
}

func (me Week) Start() string {
	return "return start date"
}
