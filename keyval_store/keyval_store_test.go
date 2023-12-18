package keyvalstore

import (
	//"fmt"
	"sync"
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

	if setVal.value != val {
		t.Fatalf("Value set incorrectly, expected %s but got %s", val, setVal.value)
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

// Tests the mutex locks in Store
func TestConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go write(&wg)
	go read(&wg)

	//Wait to exit function once both routines have completed
	wg.Wait()
}

func write(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		testStore.Set("foo", i)
	}
}

func read(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		testStore.Get("foo")
	}
}
