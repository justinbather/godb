package client

import (
	"github.com/gorilla/mux/"
	"github.com/justinbather/godb/pkg/godb"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Get the key from the request
		key := r.URL.Query().Get("key")
		//

	}
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/").Methods("POST", "GET", "PUT", "DELETE")
	http.ListenAndServe(":8080", r)
}
