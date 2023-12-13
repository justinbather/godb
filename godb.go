package godb

import "fmt"

func Hello() {
	fmt.Println("hello")
}

var dbList []DB = []DB{}

type DB struct {
	Name string
}

func List() {

	for _, db := range dbList {
		fmt.Println(db.Name)
	}
}

func CreateDb(name string) *DB {
	newDb := &DB{}
	dbList = append(dbList, *newDb)
	return newDb
}
