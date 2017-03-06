package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// func f(n int) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(n, ":", i)
// 	}
// 	fmt.Println("No. of goroutines active are:", runtime.NumGoroutine())
// }

// func main() {
// 	for index := 0; index < 50; index++ {
// 		go f(index)
// 	}
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Finished!")
// 	fmt.Println("No. of goroutines active are:", runtime.NumGoroutine())
// 	var input string
// 	fmt.Scanln(&input)
// }
