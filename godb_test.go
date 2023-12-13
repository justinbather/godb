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
	table = db.CreateTable("test table")
	db.ListTables()
}

func TestCreateCol(t *testing.T) {

	table.AppendColumn("Test Column", "int")
	table.ListColumns()
}
