package godb

import (
	"testing"
)

var db *DB
var table *Table

func TestCreateDB(t *testing.T) {
	db = CreateDb("Hello")
	List()
}

func TestCreateTable(t *testing.T) {
	table = db.CreateTable("test table", []Column{{"col1", "int"}, {"col2", "char"}})
	table.ListColumns()
	db.ListTables()
}

func TestInsertRow(t *testing.T) {
	table.Insert([]Row{"hello", "1"})
}
