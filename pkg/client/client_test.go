package client_test

import (
	"github.com/justinbather/godb/pkg/client"
	"log"
	"testing"
)

func TestConnect(t *testing.T) {

	cfg := client.New("http://localhost:3000")

	err := client.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

}
