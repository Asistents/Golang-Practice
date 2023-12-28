// Server 1 - min "echo"- server

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path =%q\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler) // all request called handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
