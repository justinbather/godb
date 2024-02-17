package godb_test

import (
	"sync"
	"testing"
	"time"

	"github.com/justinbather/godb/pkg/godb"

	log "github.com/sirupsen/logrus"
)

func write(wg *sync.WaitGroup, store *godb.Store) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		store.Set("foo", i, 5*time.Second, false)
	}
}

func read(wg *sync.WaitGroup, store *godb.Store) error {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		if _, err := store.Get("foo"); err != nil {
			return err
		}
	}
	return nil
}

func TestSmoketests(t *testing.T) {
	testStore := godb.New()

	t.Run("set", func(t *testing.T) {
		key := "foo"
		val := "bar"

		testStore.Set(key, val, 5*time.Second, false)

		setVal, ok := testStore.Data[key]
		if !ok {
			t.Fatalf("Value did not set, expected %s", val)
		}

		if setVal.Value != val {
			t.Fatalf("Value set incorrectly, expected %s but got %s", val, setVal.Value)
		}
	})

	t.Run("get", func(t *testing.T) {
		val, err := testStore.Get("foo")
		if err != nil {
			t.Fatalf("Error occured in TestGet(): %s", err)
		}

		if val != "bar" {
			t.Fatalf("Error: expected bar but got %s", val)
		}
	})

	t.Run("delete", func(t *testing.T) {
		testStore.Delete("foo")

		_, ok := testStore.Data["foo"]

		if ok {
			t.Fatalf("Error deleting value, key:foo still has a value")
		}
	})

	t.Run("ttl", func(t *testing.T) {
		// Set value with 5 second expiration
		testStore.Set("ttl", "baz", 5*time.Second, false)
		// Sleep for 6 seconds, we should be able to get the value after the sleep

		time.Sleep(6 * time.Second)

		_, err := testStore.Get("ttl")

		// This should return an error if test passes, we expect no value to be found
		if err == nil {
			t.Fatal("Error running TTL test, expected no value but found value after ttl duration")
		}
	})

	t.Run("sliding_ttl", func(t *testing.T) {
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
	})
}

// Tests the mutex locks in Store.
func TestConcurrency(_ *testing.T) {
	testStore := godb.New()
	var wg sync.WaitGroup
	wg.Add(2)
	go write(&wg, testStore)
	go func() {
		if err := read(&wg, testStore); err != nil {
			log.Println(err)
		}
	}()

	// Wait to exit function once both routines have completed
	wg.Wait()
}
