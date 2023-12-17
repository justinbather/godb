# GoDB - Key/Value In-Memory Database Implementation

## Usage

Library can be required by using:
```sh
go get github.com/justinbather/godb@latest
```

```go
package main

import (
  "github.com/justinbather/godb/keyval_store"
  "log"
  "fmt"
)

func main() {
  // Instantiate the db
  db := keyval_store.New()

  /*
   * Set a value in your database
   * @params: key string, value any (interface{})
   */
  db.Set("foo", "bar")

  /* 
   * Get a value from your database
   * @params: key string
   * @returns: value any, ok bool 
   */
  value, ok := db.Get("foo")

  if !ok {
    log.Fatal("Error: value not found at key 'foo'")
  }

  fmt.Printf("value: %s", value)

  /*
   * Remove a key/value pair from the database
   * @Params: key string
  */
  db.Delete("foo")
}
```

## Todo

- [ ] Add TTL to key/value pairs
- [ ] Add Sliding TTL 


