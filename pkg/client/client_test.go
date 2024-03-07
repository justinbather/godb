package client_test

import (
	"log"
	"strings"
	"testing"

	"github.com/justinbather/godb/pkg/client"
)

func TestConnect(t *testing.T) {
	cfg := client.New("http://localhost:8080", 10)
	err := client.Connect(cfg)
	if err != nil {
		t.Fatal("failed to connect to server")
	}
}

func TestDb(t *testing.T) {
	cfg := client.New("http://localhost:8080/", 10)

	err := client.Set(cfg, "hello", "world", 10000, false)
	if err != nil {
		t.Fatalf("Error setting value")
	}

	val, err := client.Get(cfg, "hello")
	if err != nil {
		t.Fatalf("Error getting value")
	}

	if strings.TrimSpace(val) != `"world"` {
		log.Printf("value: %s, hex: %x", val, []byte(val))
		log.Printf("expected: %s, hex: %x", "world", []byte("world"))
		log.Printf("len of val: %d", len(val))
		log.Printf("len of world: %d", len("world"))
		log.Print(val)
		t.Fatalf("Did not get the correct response. Expected world, got %s", val)
	}

	log.Println("Success")
	log.Println("value: ", val)

	err = client.Delete(cfg, "hello")
	if err != nil {
		t.Fatalf("Failed deleting key/value pair. Error: %s", err)
	}
}
