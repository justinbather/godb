package godb

import (
	//"fmt"
	"sync"
	"testing"
	"time"
)

var testStore *Store = New()

func TestSet(t *testing.T) {

	key := "foo"
	val := "bar"

	testStore.Set(key, val, 5*time.Second, false)

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

func TestTTL(t *testing.T) {
	// Set value with 5 second expiration
	testStore.Set("ttl", "baz", 5*time.Second, false)
	//Sleep for 6 seconds, we should be able to get the value after the sleep

	time.Sleep(6 * time.Second)

	_, err := testStore.Get("ttl")

	// This should return an error if test passes, we expect no value to be found
	if err == nil {
		t.Fatal("Error running TTL test, expected no value but found value after ttl duration")
	}
}

func TestSlidingTTL(t *testing.T) {
	testStore.Set("slide", "baz", 4*time.Second, true)

	time.Sleep(3 * time.Second)

	_, err := testStore.Get("slide")
	if err != nil {
		t.Fatal("Error in TestSlidingTTL: Value not found after initial sleep")
	}

	time.Sleep(5 * time.Second)

	_, err = testStore.Get("slide")
	if err == nil {
		t.Fatal("Error in TestSlidingTTL: Value found but expected error to be != nil")
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
		testStore.Set("foo", i, 5*time.Second, false)
	}
}

func read(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		testStore.Get("foo")
	}
}
