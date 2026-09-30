![godb](./media/brand.png)

# godb

An in-memory key/value store in Go with per-key TTL and optional sliding expiry windows.
Usable as a library, or as a standalone HTTP server with a matching client package.

Written to understand how cache expiry actually works rather than to compete with Redis.

## Install

```sh
go get github.com/justinbather/godb@latest
```

## Library

```go
package main

import (
	"log"
	"time"

	"github.com/justinbather/godb/pkg/godb"
)

func main() {
	db := godb.New()

	// key, value, ttl, sliding
	db.Set("foo", "bar", 5*time.Minute, true)

	value, err := db.Get("foo")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("value: %s", value)

	db.Delete("foo")
}
```

| Method | Behaviour |
| --- | --- |
| `Set(key, value, ttl, sliding)` | Stores a value and schedules its removal after `ttl`. |
| `Get(key)` | Returns the value, or an error if absent. Extends the window if the key is sliding. |
| `Delete(key)` | Removes a key immediately. |
| `GetKeys()` | Returns all currently stored keys. |

## Server

`godb` also runs as an HTTP service, with `pkg/client` as a typed client for it.

```sh
make                 # build
./bin/server         # or: docker build -t godb . && docker run -p 8080:8080 godb
```

Routes are dispatched by method on a single path: `GET` to read, `POST` to write, `DELETE` to remove.

## Design notes

- **Storage** is a `map[string]*item` behind a `sync.RWMutex`. Values are `interface{}`, so the store
  is type-agnostic and callers assert on read.
- **Expiry is per-key and push-based.** `Set` spawns a goroutine that sleeps for the TTL and then
  deletes the key. There's no background sweeper and no lazy expiry check on read, which keeps the
  read path free of clock comparisons at the cost of one goroutine per live key.
- **Sliding windows** move a key's expiry forward each time it's read, so frequently-accessed keys
  stay resident and idle ones fall out.

## Known limitations

Honest list, because the expiry model above has real consequences:

- **`Get` takes the write lock, not the read lock.** Reads are serialised despite the `RWMutex`,
  and sliding mode needs a write anyway. Splitting the sliding and non-sliding paths would let
  ordinary reads run concurrently.
- **Sliding expiry updates the timestamp but not the timer.** The removal goroutine was scheduled
  at `Set` time and still fires then, so a sliding key can be evicted while it's actively in use.
  Fixing this properly means a cancellable timer per key, or replacing push-based expiry with a
  sweeper plus a lazy check on read.
- **Overwriting a key leaves the previous timer running.** The old goroutine still fires and deletes
  whatever is at that key by then, which can evict the newer value early.
- **One goroutine per live key** puts a ceiling on how many keys the store can hold comfortably.
- **`Data` is exported**, so callers can reach the map directly and bypass the mutex.
- Everything is in memory: no persistence, no eviction policy under memory pressure, no clustering.

## Roadmap

- [x] TTL per key
- [x] Sliding TTL windows
- [ ] Cancellable expiry timers (fixes the two eviction bugs above)
- [ ] Lazy expiry check on read plus a single background sweeper
- [ ] Key search

![logo](./media/logo.png)
