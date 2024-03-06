package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/justinbather/godb/pkg/godb"
)

const (
	// Conversion from micro-seconds to seconds.
	toSeconds = 1000000000
)

func GetHandler(w http.ResponseWriter, r *http.Request, db *godb.Store) {
	key := r.URL.Query().Get("key")

	log.Println("key: ", key)

	data, getErr := db.Get(key)
	if getErr != nil {
		http.Error(w, getErr.Error(), http.StatusNoContent)
	}

	log.Println("Fetched value: ", data)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SetHandler(w http.ResponseWriter, r *http.Request, db *godb.Store) {
	var req requestItem

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// convert ttl to seconds
	db.Set(req.Key, req.Value, time.Duration(req.TTL)*toSeconds, req.Sliding)
	w.WriteHeader(http.StatusCreated)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request, db *godb.Store) {
	key := r.URL.Query().Get("key")
	db.Delete(key)
	w.WriteHeader(http.StatusNoContent)
}
