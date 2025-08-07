package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	items map[string]dollars
	sync.RWMutex
}

func (db *database) list(w http.ResponseWriter, _ *http.Request) {
	db.RLock()
	defer db.RUnlock()
	for item, price := range db.items {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	db.RLock()
	defer db.RUnlock()
	price, ok := db.items[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db *database) add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")

	if item == "" || priceStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad query string")
		return
	}

	price, err := strconv.ParseFloat(priceStr, 32)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "wrong price value: %s", priceStr)
		return
	}

	if price <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "price value must be greater than 0: %s", priceStr)
		return
	}

	db.Lock()
	defer db.Unlock()
	_, exist := db.items[item]
	if exist {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item already added: %q\n", item)
		return
	}

	db.items[item] = dollars(price)
	fmt.Fprintf(w, "item:%s with price: %s successfully added\n", item, dollars(price))
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")

	if item == "" || priceStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad query string")
		return
	}

	price, err := strconv.ParseFloat(priceStr, 32)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "wrong price value: %s", priceStr)
		return
	}

	if price <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "price value must be greater than 0: %s", priceStr)
		return
	}

	db.Lock()
	defer db.Unlock()
	_, exist := db.items[item]
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	db.items[item] = dollars(price)
	fmt.Fprintf(w, "item:%s with price: %s successfully updated\n", item, dollars(price))
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	db.Lock()
	defer db.Unlock()
	price, ok := db.items[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	delete(db.items, item)
	fmt.Fprintf(w, "item:%s with price: %s successfully deleted\n", item, dollars(price))
}

func main() {
	db := database{items: map[string]dollars{"shoes": 50, "socks": 5}}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/add", db.add)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
