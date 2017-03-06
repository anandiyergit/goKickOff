package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/julienschmidt/httprouter"
// )

// func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	fmt.Fprint(w, "Welcome!\n")
// }

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name1"))
// }

// func main() {
// 	router := httprouter.New()
// 	router.GET("/", Index)
// 	router.GET("/hello/:name1", Hello)
// 	log.Fatal(http.ListenAndServe(":8080", router))
// }
