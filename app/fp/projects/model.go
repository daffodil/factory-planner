
package projects

import (
	//"fmt"
	"github.com/jinzhu/gorm"
)


// The end product model , eg X360, Steiger5, FooBar4
type Model struct {
	ModelId int ` json:"model_id" gorm:"column:model_id; primary_key:yes" `
	BrandId int ` json:"brand_id" `
	Model string   ` json:"model" `
}

func (me Model) TableName() string {
	return "models"
}

func DB_IndexModels(db gorm.DB) {

	cols := []string{"brand_id", "model"}

	for _, c := range cols {
		db.Model(&Model{}).AddIndex("idx_" + c, c)
	}
}


func InsertModel(db gorm.DB,  brand_id int, model string) (*ModelView, error) {
	b := new(Model)
	//b.AccountId = account_id
	b.BrandId = brand_id
	b.Model = model
	db.Create(b)
	return GetModelById(db, b.ModelId)
}


// Database view extends the Account with other stuff
type ModelView struct {
	Model
	Brand string  ` json:"brand" `
	Ticker string ` json:"ticker" `
	Company string ` json:"company" `
	AccRef string ` json:"acc_ref" `
}

var MODEL_VIEW string = "v_models"
var MODEL_VIEW_COLS string = `
model_id, model, brand_id, brand, account_id, company, ticker, acc_ref
`


// All brands
// TODO redirect brands and expired brands
func GetAllModels(db gorm.DB) ([]*ModelView, error) {
	var rows []*ModelView
	db.Table(MODEL_VIEW).Select(MODEL_VIEW_COLS).Scan(&rows)
	return rows, nil
}

// Brands for the account
func GetAccountModels(db gorm.DB, account_id int) ([]*ModelView, error) {
	var rows []*ModelView
	db.Table(MODEL_VIEW).Select(MODEL_VIEW_COLS).Where("account_id = ?", account_id).Scan(&rows)
	return rows, nil
}

func GetModelById(db gorm.DB, model_id int) (*ModelView, error) {
	var row ModelView
	db.Table(MODEL_VIEW).Select(MODEL_VIEW_COLS).Where("model_id = ?", model_id).Scan(&row)
	if row.ModelId == 0 {
		return nil, nil
	}
	return &row, nil
}

func GetModelByModel(db gorm.DB,  brand_id int, model string) (*ModelView, error) {
	var row ModelView
	db.Table(MODEL_VIEW).Select(MODEL_VIEW_COLS).Where("brand_id = ? and model = ?", brand_id, model).Scan(&row)
	if row.ModelId == 0 {
		return nil, nil
	}
	return &row, nil
}

func GetModelOrCreate(db gorm.DB, brand_id int, model string) (*ModelView, error) {
	ob, err := GetModelByModel(db, brand_id, model)
	if err != nil {
		return nil, err
	}
	if ob != nil {
		return ob, nil
	}
	return InsertModel(db, brand_id, model)
}
