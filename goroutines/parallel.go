package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	fmt.Println("Strating a parallel goroutine...")
// 	go funcA(&wg)
// 	go funcB(&wg)
// 	wg.Wait()
// 	fmt.Println("Terminating...")
// }

// func funcA(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for index := 0; index < 20; index++ {
// 		time.Sleep(2 * time.Microsecond)
// 		fmt.Printf("%d ", index)
// 	}
// }

// func funcB(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for index := 1000; index < 1020; index++ {
// 		time.Sleep(1 * time.Microsecond)
// 		fmt.Printf("%d ", index)
// 	}
// }
