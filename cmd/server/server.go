package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/justinbather/godb/pkg/godb"
	"github.com/justinbather/godb/pkg/server"
)

// when a user runs this binary we should launch a server with a port expoesd
// so the user can send requests to this and we manage the godb instance

const (
	serverTimeout = 10 * time.Second
)

func main() {
	db := godb.New()
	r := mux.NewRouter()
	r.HandleFunc("/", server.HandleRequest(db)).Methods("GET", "POST", "DELETE")
	log.Print("Server started on port 8080")

	server := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  serverTimeout,
		WriteTimeout: serverTimeout,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
