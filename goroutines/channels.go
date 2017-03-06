package main

// // Perfect ping program to understand how the channels work!
// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// func SendPing(c chan string) {
// 	c <- "Hello, sending the ping!"
// }

// func AnotherSendPing(c1 chan string) {
// 	c1 <- "Hello, I am also sending the ping!"
// }

// func RecvPing(c chan string) {
// 	msg := <-c
// 	fmt.Println("Hi, I have recieved your message.")
// 	fmt.Println("Your message is:", msg)
// 	//	time.Sleep(time.Second * 1)
// }
// func AnotherRecvPing(c1 chan string) {
// 	msg := <-c1
// 	fmt.Println("Hi, I have also recieved your message.")
// 	fmt.Println("Your message is this:", msg)
// 	//	time.Sleep(time.Second * 1)
// }
// func main() {
// 	c := make(chan string)
// 	go AnotherRecvPing(c)
// 	go RecvPing(c)
// 	go AnotherSendPing(c)
// 	go SendPing(c)

// 	time.Sleep(1 * time.Second)
// 	fmt.Println("finished!")
// 	fmt.Println("No. of goroutines active are:", runtime.NumGoroutine())
// }
