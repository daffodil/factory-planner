
package cal

import (
	"time"
	"github.com/revel/revel"

)


// Shifts start in two realms..
// Factory shift - eg car manufacturing



type Shift struct {
	Start hour
	End hour
}

var shifts = make(map[string]Shift)
shifts["morning"] = Shift(6, 14)
shifts["evening"] = Shift(14, 20)
shifts["night"] = Shift(20, 6)

