package controllers

import (
	//"github.com/revel/revel"
)


type SNav struct {
	Path string
	Label string
	Title string
	Icon string
}

func StaffNav() []SNav {
	nav := make([]SNav, 0)
	nav = append(nav, SNav{"/", "Home", "Home", "home"})
	//nav = append(nav, SNav{"/staff/", "Staff Home", "Staff Home", "home"})
	nav = append(nav, SNav{"/staff/accounts", "Accounts", "Accounts Index", "info"})
	//nav = append(nav, SNav{"/staff/parts", "Parts", "Parts Index", "grid"})
	//nav = append(nav, SNav{"/staff/orders", "Orders", "Orders Index", "grid"})
	nav = append(nav, SNav{"/staff/planner", "Planner", "Orders Index", "grid"})
	nav = append(nav, SNav{"/staff/dev", "Dev", "Developer CP", "grid"})
	return nav
}
