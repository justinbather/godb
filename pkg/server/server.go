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

func HandleRequest(db *godb.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Set the header to json
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
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

		case http.MethodPost:
			var req requestItem

			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			// convert ttl to seconds
			db.Set(req.Key, req.Value, time.Duration(req.TTL)*toSeconds, req.Sliding)
			w.WriteHeader(http.StatusCreated)

		case http.MethodDelete:
			key := r.URL.Query().Get("key")
			db.Delete(key)
			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
