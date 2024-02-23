package server

import (
	"encoding/json"
	"github.com/justinbather/godb/internal/types"
	"github.com/justinbather/godb/pkg/godb"
	"log"
	"net/http"
	"time"
)

const (
	toSeconds = 1000000000
)

func HandleRequest(db *godb.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RequestItem
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
