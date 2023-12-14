package godb

import "fmt"

// Keep all created DB's in memory
var dbList []DB = []DB{}

type DB struct {
	Name   string
	Tables []Table
}

type Table struct {
	Name    string
	Columns []Column
	Rows    [][]Row
}

type Column struct {
	Name string
	Type string
}

type Row struct {
	Values []string
}

// Lists all created databases in memory
func List() {
	for _, db := range dbList {
		fmt.Println(db.Name)
	}
}

// Creates an empty database with a name
func CreateDb(name string) *DB {
	newDb := &DB{}
	dbList = append(dbList, *newDb)
	return newDb
}

func (db *DB) CreateTable(name string, cols []Column) *Table {

	table := &Table{Name: name, Columns: cols}
	db.Tables = append(db.Tables, *table)
	return table
}

func (db *DB) ListTables() {
	for idx, table := range db.Tables {
		fmt.Println(idx, table.Name)
	}
}

func (t *Table) AppendColumn(name string, colType string) error {
	col := &Column{Name: name, Type: colType}
	t.Columns = append(t.Columns, *col)

	return nil
}

func (t *Table) ListColumns() {
	for idx, col := range t.Columns {
		fmt.Println(idx, col.Name, col.Type)
	}
}

func (t *Table) Insert(values []Row) int {
	t.Rows = append(t.Rows, values)
	return len(values)
}
