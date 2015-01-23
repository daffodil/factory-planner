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

	payload["brands"], err = projects.GetAccountBrands(app.Db, account_id)
	if err != nil {
		payload["error"] = err.Error()
		return c.RenderJson(payload)
	}
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
	brandOb, eb := projects.GetBrandByBrand(app.Db, account_id, brand)
	if eb != nil {
		payload["error"] = eb.Error()
		return c.RenderJson(payload)
	}
	fmt.Println("brand=", brandOb)
	// create object if nil
	if 1 == 0 && brandOb == nil {
		fmt.Println("NO BRAND", brandOb)
		brandOb, err = projects.InsertBrand(app.Db, account_id, brand)
		if err != nil {
			payload["error"] = eb.Error()
			return c.RenderJson(payload)
		}
		fmt.Println("BRAND CREATED", brandOb)

	}

	//modelOb, e = projects.GetBrandByBrand(app.Db, model)
	//if e != nil {
	//	payload["error"] = e.Error()
	//}
	return c.RenderJson(payload)
}


// POST vars = brand and model insert if not exitst, and returned as Model record
func (c Projects) BrandModelImportJson(account_id int) revel.Result {

	var e error
	pay := MakePayload()

	brand := GetS(c.Params, "brand")
	model := GetS(c.Params, "model")
	if brand == "" {
		pay["error"] = "No `brand`"
		return c.RenderJson(pay)
	}
	if model == "" {
		pay["error"] = "No `model`"
		return c.RenderJson(pay)
	}

	brandOb, eerrb := projects.GetBrandOrCreate(app.Db, account_id, brand)
	if eerrb != nil {
		pay["error"] = eerrb.Error()
		return c.RenderJson(pay)
	}
	//pay["brand"] = brandOb

	pay["model"], e = projects.GetModelOrCreate(app.Db, brandOb.BrandId, model)
	if e != nil {
		pay["error"] = e.Error()
		return c.RenderJson(pay)
	}


	return c.RenderJson(pay)
}

