

package calendar

import (
	//"fmt"
	"time"

)


func Now() time.Time {
	return time.Now().UTC()
}

func ToString(d time.Time) string {
	return d.Format("2006-01-02")
}
