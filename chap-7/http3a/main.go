package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32
type database map[string]dollars

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		_, _ = fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	fmt.Fprintf(w, "%s\n", price)

}

func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}

	mux := http.NewServeMux()
	// The HandlerFunc type is an adapter to allow the use of
	// ordinary functions as HTTP handlers. If f is a function
	// with the appropriate signature, HandlerFunc(f) is a
	// Handler that calls f.
	// type HandlerFunc func(ResponseWriter, *Request)

	// mux.Handle("/price", http.HandlerFunc(db.price))
	// mux.Handle("/list", http.HandlerFunc(db.list))

	// mux.HandleFunc() calls mux.Handle(pattern, HandlerFunc(handler))
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}