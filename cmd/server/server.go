package main

import (
	"flag"
	"fmt"
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

	// NOTE: need to move this into a cli.go file? need to create a config struct
	portFlag := flag.String("port", "8080", "server port")

	flag.Parse()

	server := http.Server{
		Addr:         fmt.Sprintf(":%s", *portFlag),
		Handler:      r,
		ReadTimeout:  serverTimeout,
		WriteTimeout: serverTimeout,
	}

	log.Printf(fmt.Sprintf("Server started on port %s", *portFlag))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
