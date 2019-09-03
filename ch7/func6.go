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

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for i, v := range db {
			fmt.Fprintf(w, "%v---%v\n", i, v)
		}
	case "/price":
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
	case "/calc":
		sum := 0
		for i := 0; i < 5000000; i++ {
			sum += i
		}
		//fmt.Fprintf(w,"结果输出：%d-----%T",sum,w)
		http.Error(w, "no such file!", http.StatusInternalServerError)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page:%s", r.URL.Path)
	}
}

func main() {
	db := database{"shoes": 32.2, "socks": 12.2}
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}
