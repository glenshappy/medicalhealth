package main

import (
	"fmt"
	"net/http"
)

type dollars float32
type database map[string]dollars

func (db database) price(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	item := values.Get("item")
	fmt.Printf("%#v", values)
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s", price)
}

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for i, v := range db {
		fmt.Fprintf(w, "%v---%v\n", i, v)
	}
}

func main() {
	db := database{"shoes": 32.2, "socks": 12.2}
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/list", db.list)
	http.ListenAndServe("localhost:8083", nil)
}
