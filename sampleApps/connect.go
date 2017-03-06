package main

// import (
// 	"fmt"
// 	"time"
// )

// var isConnected bool = false

// func main() {
// 	fmt.Printf("Connection open: %v\n", isConnected)
// 	//doSomething()
// 	defer sleepLess()
// 	defer sleepMore()

// 	fmt.Printf("Connection open: %v\n", isConnected)
// }

// func doSomething() {
// 	connect()
// 	fmt.Println("Deferring disconnect!")
// 	defer disconnect()
// 	fmt.Printf("Connection open: %v\n", isConnected)
// 	fmt.Println("Doing something!")
// }

// func connect() {
// 	isConnected = true
// 	fmt.Println("Connected to database!")
// }

// func disconnect() {
// 	isConnected = false
// 	fmt.Println("Disconnected!")
// }

// func sleepLess() {
// 	fmt.Println("Going to sleep for 100")
// 	time.Sleep(1000 * time.Millisecond)
// 	fmt.Println("Done sleeping for 100")
// }

// func sleepMore() {
// 	fmt.Println("Going to sleep for 500")
// 	time.Sleep(5000 * time.Millisecond)
// 	fmt.Println("Done sleeping for 500")
// }
