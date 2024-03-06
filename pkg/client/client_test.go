package client_test

import (
	"github.com/justinbather/godb/pkg/client"
	"log"
	"testing"
)

func TestConnect(t *testing.T) {

	cfg := client.New("http://localhost:8080")
	err := client.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

}

func TestDb(t *testing.T) {
	cfg := client.New("http://localhost:8080/")

	err := client.Set(cfg, "hello", "world", 100, false)
	if err != nil {
		t.Fatalf("Error setting value")
	}

	val, err := client.Get(cfg, "hello")
	if err != nil {
		t.Fatalf("Error getting value")
	}

	log.Println("Success")
	log.Println("value: ", val)
}
