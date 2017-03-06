package main

// // This program showcases leaks in go routines
// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"runtime"
// 	"strconv"
// )

// // function to add an array of numbers.
// func sum(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	// writes the sum to the go routines.
// 	c <- sum // send sum to c
// }

// // HTTP handler for /sum
// func sumConcurrent(w http.ResponseWriter, r *http.Request) {
// 	s := []int{7, 2, 8, -9, 4, 0}

// 	c1 := make(chan int)
// 	c2 := make(chan int)
// 	// spin up a goroutine.
// 	go sum(s[:len(s)/2], c1)
// 	// spin up a goroutine.
// 	go sum(s[len(s)/2:], c2)
// 	// not reading from c2.
// 	// go routine writing to c2 will be blocked.
// 	x, y := <-c1, <-c2
// 	// write the response.
// 	fmt.Fprintf(w, strconv.Itoa(x+y))
// }

// // get the count of number of go routines in the system.
// func countGoRoutines() int {
// 	return runtime.NumGoroutine()
// }

// func getGoroutinesCountHandler(w http.ResponseWriter, r *http.Request) {
// 	// Get the count of number of go routines running.
// 	count := countGoRoutines()
// 	w.Write([]byte(strconv.Itoa(count)))
// }

// func main() {
// 	http.HandleFunc("/sum", sumConcurrent) // set router
// 	http.HandleFunc("/_count", getGoroutinesCountHandler)
// 	err := http.ListenAndServe(":8001", nil) // set listen port
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }
