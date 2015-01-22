
package projects

import (
	"github.com/jinzhu/gorm"
)


// The end product Brand , eg Volvo, Class, Jaguar
type Brand struct {
	BrandId int ` json:"brand_id" gorm:"column:brand_id; primary_key:yes" `
	AccountId int ` json:"account_id" `
	Brand string   ` json:"brand" `
}

func (me Brand) TableName() string {
	return "brands"
}

func DB_IndexBrands(db gorm.DB) {

	cols := []string{"account_id", "brand"}

	for _, c := range cols {
		db.Model(&Brand{}).AddIndex("idx_" + c, c)
	}
}

func InsertBrand(db gorm.DB, account_id int, brand string) (*BrandView, error) {
	b := new(Brand)
	b.AccountId = account_id
	b.Brand = brand
	db.Create(b)
	return GetBrandById(db, b.BrandId)
}


// Database view extends the Account with other stuff
type BrandView struct {
	Brand
	Ticker int ` json:"ticker" `
	Company int ` json:"company" `
	AccRef int ` json:"acc_ref" `
}

var BRAND_VIEW string = "v_brands"
var BRAND_VIEW_COLS string = `
brand_id, account_id, company, ticker, acc_ref, brand
`


// All brands
// TODO redirect brands and expired brands
func GetAllBrands(db gorm.DB) ([]*BrandView, error) {
	var rows []*BrandView
	db.Table(BRAND_VIEW).Select(BRAND_VIEW_COLS).Scan(&rows)
	return rows, nil
}

// Brands for the account
func GetAccountBrands(db gorm.DB, account_id int) ([]*BrandView, error) {
	var rows []*BrandView
	db.Table(BRAND_VIEW).Select(BRAND_VIEW_COLS).Where("account_id = ?", account_id).Scan(&rows)
	return rows, nil
}

func GetBrandById(db gorm.DB, brand_id int) (*BrandView, error) {
	var row *BrandView = new(BrandView)
	db.Table(BRAND_VIEW).Select(BRAND_VIEW_COLS).Where("brand_id = ?", brand_id).Scan(row)
	return row, nil
}

func GetBrandByBrand(db gorm.DB, brand string) (*BrandView, error) {
	var row *BrandView
	db.Table(BRAND_VIEW).Select(BRAND_VIEW_COLS).Where("brand = ?", brand).Scan(&row)
	return row, nil
}

