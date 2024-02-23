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

var db *godb.Store = godb.New()

const toSeconds = 1000000000

type requestItem struct {
	Value   interface{} `json:"value"`
	Key     string      `json:"key"`
	TTL     int         `json:"ttl"`
	Sliding bool        `json:"sliding"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var req requestItem
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// Set the header to json
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		key := r.URL.Query().Get("key")
		data, err := db.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Fatal(err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	case "POST":
		// convert ttl to seconds
		db.Set(req.Key, req.Value, time.Duration(req.TTL)*toSeconds, req.Sliding)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	case "DELETE":
		db.Delete(req.Key)
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleRequest).Methods(http.MethodGet, http.MethodPost, http.MethodDelete)
	log.Print("Server started on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
