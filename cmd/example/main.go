package main

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/justinbather/godb/pkg/godb"
)

const (
	ttl = 5 * time.Minute
)

func main() {
	// Instantiate the db
	db := godb.New()

	/*
	 * Set a value in your database
	 * @params: key string, value any (interface{}), ttl time.Duration, sliding bool
	 */
	db.Set("foo", "bar", ttl, true)
	// GoDB will take the given TTL duration and have a worker thread remove the entry automatically
	// once the TTL value has elapsed. If sliding is set to true for this value, the window will be
	// extended by the TTL value whenever the KV pair is accessed

	/*
	 * Get a value from your database. If sliding is true, expiration time will be moved forward
	 * @params: key string
	 * @returns: value any, ok bool
	 */
	value, err := db.Get("foo")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("value: %s", value)

	/*
	 * Remove a key/value pair from the database
	 * @Params: key string
	 */
	db.Delete("foo")
}
