package server

import (
	"net/http"

	"github.com/justinbather/godb/pkg/godb"
)

func HandleRequest(db *godb.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Set the header to json
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			GetHandler(w, r, db)

		case http.MethodPost:
			SetHandler(w, r, db)

		case http.MethodDelete:
			DeleteHandler(w, r, db)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
