package controllers

import (
	"fmt"
	//"strings"
	"github.com/revel/revel"

	"github.com/daffodil/factory-planner/app"
	//"github.com/daffodil/factory-planner/app/fp/orders"
	"github.com/daffodil/factory-planner/app/fp/projects"
)


type Projects struct {
	*revel.Controller
}

func (c Projects) AccountBrandsJson(account_id int) revel.Result {
	var err error
	payload := MakePayload()
	payload["account_id"] = account_id

	payload["brands"], err = projects.GetBrandsForAccount(app.Db, account_id)

	return c.RenderJson(payload)
}

func (c Projects) ModelPostJson(account_id int, model_id string) revel.Result {

	var err error
	payload := MakePayload()
	payload["account_id"] = account_id
	payload["model_id"] = model_id


	brand := GetS(c.Params, "brand")
	fmt.Println("brand=", brand)
	if brand == "" {
		payload["error"] = "No brand"
		return c.RenderJson(payload)
	}
	// find object
	brandOb, eb := projects.GetBrandByBrand(app.Db, brand)
	if eb != nil {
		payload["error"] = eb.Error()
		return c.RenderJson(payload)
	}
	// create object if nil
	if brandOb == nil {
		fmt.Println("NO BRAND", brandOb)
		brandOb, err = projects.InsertBrand(app.Db, account_id, brand)
		if err != nil {
			payload["error"] = eb.Error()
			return c.RenderJson(payload)
		}

	}

	//modelOb, e = projects.GetBrandByBrand(app.Db, model)
	//if e != nil {
	//	payload["error"] = e.Error()
	//}
	return c.RenderJson(payload)
}

/*
func (c Projects) ModelPostJson() revel.Result {

	var e error
	payload := MakePayload()

	payload["brand"], e = projects.GetBrandByBrand(app.Db, "foo")
	if e != nil {
		payload["error"] = e.Error()
	}

	return c.RenderJson(payload)
}
*/
