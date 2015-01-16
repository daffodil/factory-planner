package orders

/* Orders and Legers

The orders belong in either the sales, purchase or quotes leger

The legers are defined as Globals as these need to be used in code
as well as the DB queries etc

 */

import (
	//"fmt"
	//"time"
	"github.com/jinzhu/gorm"

)


type Leger struct {
	LegerId int `json:"leger_id" gorm:"column:leger_id; primary_key:yes"`
	Leger string `json:"leger" sql:"type:int"`
}

var SALES_LEGER *Leger
var PURCHASE_LEGER *Leger
var QUOTE_LEGER *Leger

func init() {

	SALES_LEGER = new(Leger)
	SALES_LEGER.LegerId = 10
	SALES_LEGER.Leger = "Sales Leger"

	PURCHASE_LEGER = new(Leger)
	PURCHASE_LEGER.LegerId = 20
	PURCHASE_LEGER.Leger = "Purcahse Leger"

	QUOTE_LEGER = new(Leger)
	QUOTE_LEGER.LegerId = 30
	QUOTE_LEGER.Leger = "Quotes"

}


func (me Leger) TableName() string {
	return "legers"
}
func DB_IndexLegers(db gorm.DB) {

	cols := []string{"leger"}

	for _, c := range cols {
		db.Model(&Leger{}).AddIndex("idx_" + c, c)
	}
}


func DB_CreateDefaultLegers(db gorm.DB) error {


	legers := make([]*Leger, 0)

	var count int
	for _, rec := range legers {
		db.Model(Leger{}).Where("leger_id = ?", rec.LegerId).Count(&count)
		if count == 0 {
			db.Create(rec)
		}
	}
	return nil
}
