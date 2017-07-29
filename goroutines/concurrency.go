package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(7)

	fmt.Println("Starting Go Routines")
	go func() {
		defer wg.Done()

		for number1 := 1001; number1 < 1027; number1++ {
			fmt.Printf("%d ", number1)
		}
	}()

	go func() {
		defer wg.Done()

		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
