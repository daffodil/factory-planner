
package fp

import (
	"fmt"
	"strings"
	"net/url"
)


const (

)

// Hold the search parameters from query
// ?search=foo bar - split to parts
// ?active = all - return active, inactive or all, also 0, 1, 2 respectively
type SearchVars struct {
	Search []string
	FilterActive string
	//Enabled bool
}

func (me SearchVars) GetSQL(fld, active string) string {
	P := "%%"
	sql := "  1 = 1 "
	if len(me.Search) > 0 {
		for _, snippet := range me.Search {
			s := fmt.Sprintf(" and `%s` like ", fld)
			s +=  "'" + P + snippet + P + "' "
			sql += s
		}
	}
	return sql
}


func SplitString(s string) []string {
	ret := make([]string, 0)
	parts := strings.Split(s, " ")
	for _, p := range parts {
		stripped := strings.TrimSpace(p)
		if stripped != "" {
			ret = append(ret, stripped)
		}
	}
	return ret
}

// Load SearchVars from query
func GetSearchVars(query url.Values) SearchVars {

	sv := SearchVars{}
	//sv.Search = nil
	search := strings.TrimSpace( query.Get("search") )
	if search != "" {
		sv.Search = SplitString(search)
	}
	fmt.Println(sv.Search)
	return sv
}
