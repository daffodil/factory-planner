
package files


import (

	//"fmt"
	"time"
	"github.com/jinzhu/gorm"

	"github.com/daffodil/factory-planner/app/fp"
)


type File struct {

	FileId int `json:"file_id" gorm:"primary_key:yes"`
	//FileTypeId int `json:"file_type_id"`

	// person who uploaded this file
	ContactId int `json:"contact_id" sql:"type:int(2);not null;" `

	FileName string `json:"file_name"`
	FileDescription string `json:"file_description"`
	MimeType string `json:"mime_type" `

	Revision int ` json:"revision" sql:"type:int(2);not null;default:0" `
	FileDate time.Time  ` json:"file_date"  `

	FileChecksum string `json:"file_hash" gorm:"column:file_checksum;" sql:"type:varchar(255);not null;default:''" `
	FileUID string `json:"file_uid" gorm:"column:file_uid;" sql:"type:varchar(255);not null;default:''" `
	//FileSearch string `json:"file_search" sql:"type:varchar(255);not null;default:''" `

}
func (me File) TableName() string {
	return "files"
}
// Adds indexes to files table
func DB_IndexFiles(db gorm.DB) {

	cols := []string{
		"file_name", "file_description", "account_id", "mime_type", "file_date",
		"revision"	, "file_hash", "file_uid"}

	for _, c := range cols {
		db.Model(&File{}).AddIndex("idx_" + c, c)
	}
}


// Database view extends the Account with other stuff
type FileView struct {
	File
	Company string ` json:"company" `
	Ticker string ` json:"ticker" `
	Contact string ` json:"contact" `
}
var FILES_VIEW string = "v_files"
var FILES_VIEW_COLS string = `
file_id, file_name,
file_date, file_uid, file_checksum,
mime_type, revision,
account_id, company, ticker,
contact_id, contact
`

// returns search and view results
func GetFilesIndex(db gorm.DB, search_vars fp.SearchVars) ([]*FileView, error) {

	var rows []*FileView
	//fmt.Println("getttttts=", search_vars)

	//where := search_vars.GetSQL("company", "acc_active")
	//fmt.Println("where=", where)
	//db.Table("v_accounts").Select(account_view_cols).Where(where).Scan(&rows)
	db.Table(FILES_VIEW).Select(FILES_VIEW_COLS).Scan(&rows)

	return rows, nil
}

// Return file by ID
func GetFile(db gorm.DB, file_id int)(*FileView, error) {

	var row *FileView = new(FileView)
	db.Table(FILES_VIEW).Select(FILES_VIEW_COLS).Where("file_id = ?", file_id).Scan(row)

	return row, nil
}
