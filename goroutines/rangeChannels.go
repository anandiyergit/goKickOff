package main

import (
	"fmt"
	"strconv"
	"time"
)

var i int

func main() {
	cake := make(chan string)
	chocolate := make(chan string)
	for index := 0; index < 5; index++ {
		go recvAndPackCake(cake)
		go recvAndPackCake(chocolate)
		go makeAndSendCake(cake)
		go makeAndSendCake(chocolate)

	}
	time.Sleep(2 * time.Second)
}
func makeAndSendCake(cake chan string) {
	i = i + 1
	cakeName := "Cake " + strconv.Itoa(i)
	fmt.Println("Made and Sent Cake ", strconv.Itoa(i))
	cake <- cakeName
}

func recvAndPackCake(cake chan string) {
	msg := <-cake
	fmt.Println("Recieved " + msg + " and is sent out to packing")
}
