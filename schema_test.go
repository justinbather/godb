package godb

import (
	"fmt"
	"testing"
)

var store *Store

func TestCreateStore(t *testing.T) {

	store = NewStorer()
}

func TestSet(t *testing.T) {
	data := []byte{'i'}
	store.Set("test", data)
}

func TestGet(t *testing.T) {
	data, err := store.Get("test")

	if err != nil {
		t.Fatalf("err %s", err)
	}

	fmt.Println(data)
}

func TestPrint(t *testing.T) {
	store.ListContents()
}
