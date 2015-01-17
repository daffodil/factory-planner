
package files


import (

	//"fmt"
	"time"
	"github.com/jinzhu/gorm"

	"github.com/daffodil/factory-planner/app/fp"
)


type File struct {

	FileId int `json:"file_id" gorm:"primary_key:yes"`
	AccountId int `json:"file_id" sql:"type:int(2);not null;" `
	//FileTypeId int `json:"file_type_id"`

	FileName string `json:"file_name"`
	FileDescription string `json:"file_description"`
	MimeType string `json:"mime_type" `

	Revision int ` json:"revision" sql:"type:int(2);not null;default:0" `
	FileDate time.Time  ` json:"file_date"  `

	FileSHA string `json:"file_hash" sql:"type:varchar(255);not null;default:''" `
	FileUID string `json:"file_uid" sql:"type:varchar(255);not null;default:''" `
	FileSearch string `json:"file_search" sql:"type:varchar(255);not null;default:''" `

}
func (me File) TableName() string {
	return "files"
}
// Adds indexes to files table
func DB_IndexFiles(db gorm.DB) {

	cols := []string{
		"file_name", "file_description", "account_id", "mime_type", "file_date",
		"revision"	, "file_hash", "file_uid", "file_search"}

	for _, c := range cols {
		db.Model(&File{}).AddIndex("idx_" + c, c)
	}
}


// Database view extends the Account with other stuff
type FileView struct {
	File
	AccountId int ` json:"account_id" `
	Company string ` json:"company" `
	Ticker string ` json:"ticker" `

}

// Columns for select, messy I know if anyones got brighter ideas
var file_view_cols string = `
account_id, company, ticker, acc_ref, root, acc_active,
on_hold, is_client, is_supplier, orders_due
`

// returns search and view results
func GetFilesIndex(db gorm.DB, search_vars fp.SearchVars) ([]FileView, error) {

	var rows []FileView
	//fmt.Println("getttttts=", search_vars)

	//where := search_vars.GetSQL("company", "acc_active")
	//fmt.Println("where=", where)
	//db.Table("v_accounts").Select(account_view_cols).Where(where).Scan(&rows)
	db.Table("v_files").Select(file_view_cols).Scan(&rows)

	return rows, nil
}

// Return file by ID
func GetFile(db gorm.DB, file_id int)(*FileView, error) {

	var row *FileView = new(FileView)
	db.Table("v_files").Select(file_view_cols).Where("file_id = ?", file_id).Scan(row)

	return row, nil
}
