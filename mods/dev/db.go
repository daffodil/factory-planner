
package dev

import (
	"os"
	//"path/filepath"
	//"io/ioutil"
	"fmt"
	"github.com/revel/revel"
	"database/sql"
)

// Information about the database, its env etc
// TODO later
type DB_Info struct {
	Success bool `json:"success"`
	//Table []navdb.DB_TableInfo `json:"table"`
	Db string // explain running env and alike
}

//===========================================
// Info about a table, columns etc
type DB_TablesPayload struct {
	Success bool `json:"success,omitempty"`
	Tables []DB_Table `json:"tables"`
}
func (me *DB_TablesPayload) AddTables(tables []DB_Table) {
	for _, v := range tables {
		me.Tables = append(me.Tables, v)
	}
}
func NewTablesPayload() DB_TablesPayload {
	t := DB_TablesPayload{Success: true}
	t.Tables = make([]DB_Table,0)
	return t
}


//===========================================
// Info about a table, columns etc
type DB_TablePayload struct {
	Success bool `json:"success"`
	Table DB_Table `json:"table"`
}
func NewTablePayload() DB_TablePayload {
	ob := DB_TablePayload{Success: true}
	return ob
}

type DB_TableColPayload struct {
	Success bool `json:"success"`
	Table DB_TableCol `json:"table"`
}
func NewTableColPayload() DB_TableColPayload {
	ob := DB_TableColPayload{Success: true}
	return ob
}




// returns views and tables
func DB_GetTablesAndViewsPayload(DB *sql.DB) (DB_TablesPayload,  error) {

	payload := NewTablesPayload()

	views, err := DB_GetViews(DB)
	if err != nil {

	}
	payload.AddTables(views)

	tables, err := DB_GetTables(DB)
	if err != nil {

	}
	payload.AddTables(tables)

	return payload, nil
}
// returns views
func DB_GetViewsPayload(DB *sql.DB) (DB_TablesPayload,  error) {

	payload := NewTablesPayload()

	views, err := DB_GetViews(DB)
	if err != nil {

	}
	payload.AddTables(views)

	return payload, nil
}

// returns tables
func DB_GetTablesPayload(DB *sql.DB) (DB_TablesPayload,  error) {

	payload := NewTablesPayload()

	tables, err := DB_GetTables(DB)
	if err != nil {

	}
	payload.AddTables(tables)

	return payload, nil
}



func DB_GetViews(DB *sql.DB) ([]DB_Table,  error) {

	lst :=  make( []DB_Table, 0)

	sql := "select table_namem table_type, engine, table_rows, auto_increment  "
	sql += " from INFORMATION_SCHEMA.views WHERE table_schema = ?"
	rows, err := DB.Query(sql)
	if err != nil {
		revel.ERROR.Println(err)
		return nil,  err
	}
	defer rows.Close()

	for rows.Next(){
		t := DB_Table{}
		err := rows.Scan( &t.Name )
		if err != nil {
			revel.ERROR.Println(err)
		} else {
			t.IsView = true
			lst = append(lst, t)
		}
	}
	return lst, nil
}

func GetDatabaseName() string {
	name, found := revel.Config.String("db.database")
	if found {
		return name
	}
	return ""
}
// Table details

type DB_Table struct {
	Name string `db:"table_name" json:"table_name"`
	IsView bool `json:"is_view"`
	//Definition string `db:"definition" json:"definition,omitempty"`
	//Name string `db:"table_name" json:"table_name"`
	//IsView bool `json:"is_view"`
	Engine *string `json:"engine"`
	//Columns []DB_Column  `json:"columns,omitempty"`
	//ColCount *int64 `json:"col_count"`
	RowCount *int64 `json:"row_count"`
	NextId *int64  `json:"next_id"`
}

func DB_GetTables(DB *sql.DB)([]DB_Table,  error) {


	lst :=  make( []DB_Table, 0)

	//sql := "SELECT table_name FROM  INFORMATION_SCHEMA.tables WHERE  TABLE_SCHEMA = 'trt-staff-test' and table_type = 'BASE TABLE' AND table_schema NOT IN ('pg_catalog', 'information_schema');"
	q := "select TABLE_NAME as table_name, TABLE_TYPE as table_type, ENGINE as engine, TABLE_ROWS as row_count, AUTO_INCREMENT as next_id  "
	q += " from INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = ?"
	rows, err := DB.Query(q, GetDatabaseName() )
	if err != nil {
		revel.ERROR.Println(err)
		return nil,  err
	}
	defer rows.Close()

	for rows.Next(){
		t := DB_Table{}
		var ttype sql.NullString

		err := rows.Scan( &t.Name, &ttype, &t.Engine, &t.RowCount, &t.NextId)

		if err != nil {
			revel.ERROR.Println(err)
		} else {
			if ttype.Valid && ttype.String == "VIEW" {
				t.IsView = true
			}
			lst = append(lst, t)
		}
	}
	return lst, nil
}


// returns table information, columns
func DB_GetTablePayload(DB *sql.DB, table_name string)(DB_TableColPayload, error) {

	var err error
	tinfo := DB_TableColPayload{}

	tinfo.Table, err = DB_GetTable(DB, table_name)
	if err != nil {
		return tinfo, err
	}
	//tinfo.Table.Columns, err = DB_GetColumns(table_name)

	return tinfo, err
}
type DB_TableCol struct {
	Name string `db:"table_name" json:"table_name"`
	IsView bool `json:"is_view"`
	//Definition string `db:"definition" json:"definition,omitempty"`
	//Name string `db:"table_name" json:"table_name"`
	//IsView bool `json:"is_view"`
	//Engine *string `json:"engine"`
	Columns []DB_Column  `json:"columns,omitempty"`
	//ColCount *int64 `json:"col_count"`
	//RowCount *int64 `json:"row_count"`
	//NextId *int64  `json:"next_id"`
}
func DB_GetTable(DB *sql.DB, table_name string)(DB_TableCol,  error) {

	tbl  := DB_TableCol{Name: table_name, IsView: false}

	var err error
	tbl.Columns, err = DB_GetColumns(DB, table_name)
	if err != nil {
		return tbl, err
	}
	//stbl.Definition, err = DB_GetCreateTableDefinition(DB, table_name)
	return tbl, nil
}



// Column Details
type DB_Column struct {
	Name string `json:"name"`
	Ordinal int `json:"ordinal"`
	Type string `json:"type"`
	//MaxLen *int `json:"max_len"`
	Nullable bool `json:"nullable"`
	Default *string `json:"default"`
	PK bool `json:"primary_key"`
}


func DB_GetColumns(DB *sql.DB, table_name string)([]DB_Column,  error) {
	fmt.Println("DB_GetColumns", table_name)
	lst := make([]DB_Column, 0)
	//sql := "select column_name, ordinal_position, column_default, is_nullable, data_type, character_maximum_length from INFORMATION_SCHEMA.COLUMNS where table_name = $1;"
	q  := "select column_name, ordinal_position, column_type, is_nullable, column_default, extra "
	q += " from INFORMATION_SCHEMA.COLUMNS "
	q += " where table_schema = ? and table_name = ? "
	q += " order by  ordinal_position "
	rows, err := DB.Query(q, GetDatabaseName(), table_name)
	if err != nil {
		revel.ERROR.Println(err)
		return nil,  err
	}
	defer rows.Close()

	var null_str string
	var col_default *string
	var extra *string
	//var default_str sqlsql := "select column_name, ordinal_position, data_type, is_nullable from INFORMATION_SCHEMA.COLUMNS where table_name = $1;".NullString
	for rows.Next(){
		col := DB_Column{}

		//err := rows.Scan( &t.Name, &t.Ordinal, &t.Default, &t.Nullable, &t.Type, &t.MaxLen )
		err := rows.Scan( &col.Name, &col.Ordinal,  &col.Type, &null_str, &col_default, &extra)
		if err != nil {
			revel.ERROR.Println(err)
		} else {
			if null_str == "YES" {
				col.Nullable = true
			}
			if extra != nil && *extra == "auto_increment" {
				col.PK = true
			}
			lst = append(lst, col)
		}
	}
	fmt.Println(lst)
	return lst, nil
}



//=================================================================
func DB_GetViewPayload(DB *sql.DB, table_name string) (DB_TableColPayload , error) {
	t := NewTableColPayload()
	var err error



	t.Table = DB_TableCol{Name: table_name, IsView: true}
	//t.Table.Definition, err = DB_GetViewDefinition(DB, table_name)
	t.Table.Columns, err = DB_GetColumns(DB, table_name)
	//t.Script, err = GetSqlScript(table)

	return t, err
}

func DB_GetCreateTableDefinition(DB *sql.DB, table_name string)(string,  error) {

	sql := "SELECT generate_create_table_statement($1);"

	var s string
	err := DB.QueryRow(sql, table_name).Scan(&s)
	if err != nil {
		return "", err
	}
	return s, nil
}



func DB_GetViewDefinition(DB *sql.DB, table string)(string,  error) {

	sql := "select view_definition from INFORMATION_SCHEMA.views WHERE table_name = $1"

	var s string
	err := DB.QueryRow(sql, table).Scan(&s)
	if err != nil {
		return "", err
	}
	return s, nil
}


func RootPath() string {
	path := os.Getenv("GOPATH")
	path += "/src/github.com/daffodil/factory-planner"
	return path
}
/*
func GetSqlScript(table string)(string,  error) {

	file_path := RootPath() + "/sql/" + table + ".sql"
	fmt.Println(file_path)
	contents, err := ioutil.ReadFile(file_path)
	return string(contents), err
}
*/


/*
func DB_UpdateView(table string) error {

	var err error
	sqldrop := "drop view if exists $1"
	_, errd := DB.Exec(sqldrop, table)
	if  errd != nil {
		return err
	}

	sql, err := GetSqlScript(table)

	if  err != nil {
		return err
	}
	fmt.Print(sql)
	_, err = DB.Exec(sql)

	return err
}

*/

