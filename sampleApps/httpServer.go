package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func listProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Man")
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	// add a product
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["productID"]
	log.Printf("fetching product with ID %q", id)
	fmt.Fprintln(w, "Hello Man, I am l", id)
}

func main() {
	r := mux.NewRouter()
	// match only GET requests on /product/
	r.HandleFunc("/product/", listProducts).Methods("GET")

	// match only POST requests on /product/
	r.HandleFunc("/product/", addProduct).Methods("POST")

	// match GET regardless of productID
	r.HandleFunc("/product/{productID}", getProduct)

	// handle all requests with the Gorilla router.
	http.Handle("/", r)
	if err := http.ListenAndServe("127.0.0.1:7777", nil); err != nil {
		log.Fatal(err)
	}
}
