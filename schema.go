package godb

import (
	"fmt"
)

type MemStore struct {
	Name   string
	Stores map[string]StoreSchema
}

var MemDb *MemStore

type StoreSchema interface{}

func (ms *MemStore) NewStore(name string, s interface) {

  store := &StoreSchema

}

func Init(name string) *MemStore {
	MemDb = &MemStore{Name: name}
	fmt.Println("New MemDb created")
	return MemDb
}
