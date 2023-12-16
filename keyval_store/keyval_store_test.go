package keyvalstore

import (
	"testing"
)

var testStore *Store = New()

func TestSet(t *testing.T) {

	key := "foo"
	val := "bar"

	testStore.Set(key, val)

	setVal, ok := testStore.Data[key]
	if !ok {
		t.Fatalf("Value did not set, expected %s", val)
	}

	if setVal != val {
		t.Fatalf("Value set incorrectly, expected %s but got %s", val, setVal)
	}
}

func TestGet(t *testing.T) {

	val, err := testStore.Get("foo")
	if err != nil {
		t.Fatalf("Error occured in TestGet(): %s", err)
	}

	if val != "bar" {
		t.Fatalf("Error: expected bar but got %s", val)
	}
}

func TestDelete(t *testing.T) {
	testStore.Delete("foo")

	_, ok := testStore.Data["foo"]

	if ok {
		t.Fatalf("Error deleting value, key:foo still has a value")
	}

}
