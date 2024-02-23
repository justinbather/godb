package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/justinbather/godb/pkg/godb"
)

// when a user runs this binary we should launch a server with a port expoesd
// so the user can send requests to this and we manage the godb instance

const (
	toSeconds     = 1000000000
	serverTimeout = 10 * time.Second
)

type requestItem struct {
	Value   interface{} `json:"value"`
	Key     string      `json:"key"`
	TTL     int         `json:"ttl"`
	Sliding bool        `json:"sliding"`
}

func handleRequest(db *godb.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req requestItem
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Set the header to json
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			key := r.URL.Query().Get("key")
			data, getErr := db.Get(key)
			if getErr != nil {
				http.Error(w, getErr.Error(), http.StatusNotFound)
				log.Fatal(getErr)
				return
			}
			err = json.NewEncoder(w).Encode(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				log.Fatal(err)
				return
			}
		case http.MethodPost:
			// convert ttl to seconds
			db.Set(req.Key, req.Value, time.Duration(req.TTL)*toSeconds, req.Sliding)
			w.WriteHeader(http.StatusCreated)
		case http.MethodDelete:
			db.Delete(req.Key)
			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func main() {
	db := godb.New()
	r := mux.NewRouter()
	r.HandleFunc("/", handleRequest(db)).Methods("GET", "POST", "DELETE")
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
