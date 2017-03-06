package main

import (
	"fmt"
	"runtime"
	"time"
)

func Runner(baton chan int) {
	var newRunner int
	runner := <-baton
	fmt.Printf("The Runner %d Running with Baton!\n", runner)
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d to the line\n", newRunner)
		go Runner(baton)
	}
	if runner != 4 {
		fmt.Println("For every Runner to complete we will wait for 1 second!")
	}
	time.Sleep(1 * time.Second)
	if runner == 4 {
		fmt.Printf("Runner %d finished the race!\n", runner)
		return
	}
	fmt.Printf("Runner %d Exchanged with Runner %d\n", newRunner, runner)
	baton <- newRunner
}
func main() {
	baton := make(chan int)
	go Runner(baton)
	fmt.Println("Printing!")
	baton <- 1
	fmt.Println("The main goroutine will wait for 5 seconds!")
	fmt.Println("No. of goroutines active are:", runtime.NumGoroutine())
	time.Sleep(5 * time.Second)
	fmt.Println("Final No. of goroutines active are:", runtime.NumGoroutine())
}
