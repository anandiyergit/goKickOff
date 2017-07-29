package main

// import (
// 	"fmt"
// 	"strconv"
// 	"time"
// )

// // 1. Make the cake.
// // 2. Send the cake.
// var i int

// func main() {
// 	cake := make(chan string)
// 	for index := 0; index < 5; index++ {
// 		go makeCake(cake)
// 		go sendCake(cake)
// 	}
// 	time.Sleep(1 * time.Second)
// }

// func makeCake(cake chan string) {
// 	i = i + 1
// 	cakeName := "Cake " + strconv.Itoa(i)
// 	fmt.Println(cakeName, " is made and ready to be sent!")
// 	cake <- cakeName
// }

// func sendCake(cake chan string) {
// 	msg := <-cake
// 	fmt.Println(msg, " is sent for delivery!")
// }
