// Server 2 - min "echo" - server with request counter
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

// Handler, return path request URL
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.Path)
}

//Counter, return coun of request
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()                           // mu.Lock() and mu.Unlock() are neded for variable
	fmt.Fprintf(w, "Count %d\n", count) // against race condition
	mu.Unlock()
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
