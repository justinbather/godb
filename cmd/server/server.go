package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/justinbather/godb/pkg/server"

	"github.com/justinbather/godb/pkg/godb"
)

func main() {
	db := godb.New()
	r := mux.NewRouter()
	r.HandleFunc("/", server.HandleRequest(db)).Methods("GET", "POST", "DELETE")

	// NOTE: need to move this into a cli.go file? need to create a config struct

	config := server.ConfigFromFlags()

	server := http.Server{
		Addr:         fmt.Sprintf(":%s", config.Port),
		Handler:      r,
		ReadTimeout:  time.Duration(config.ReadTimeout * int(time.Second)),
		WriteTimeout: time.Duration(config.WriteTimeout * int(time.Second)),
	}

	log.Printf("Server started on port %s", config.Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
